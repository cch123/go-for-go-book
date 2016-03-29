require 'asciidoctor'
require 'asciidoctor/extensions'
require 'pathname'

CACHE_DIR = Pathname.new('.cache')

# ref: http://asciidoctor.org/docs/user-manual/#block-macro-processor-example

# Expands to an example code under ./examples or its output
# TODO: place a link to Go playground
#
#   goexample::parsefile[]
#   goexample::parsefile[output]
#
# Runs examples/parseexpr/parseexpr.go
class GoExampleMacro < Asciidoctor::Extensions::BlockMacroProcessor
  use_dsl

  named :goexample

  def process(parent, target, attrs)
    file = Pathname.new("examples/#{target}/#{target}.go")
    style = attrs.delete(1)

    if style === 'output'
      cache_file = Pathname.new(CACHE_DIR+'go-run'+file)
      cache_mtime = cache_file.mtime rescue Time.new(0)
      content = if cache_mtime < file.mtime
        c = %x(go run #{file} 2>&1)
        cache_file.parent.mkpath
        cache_file.write(c)
        c
      else
        cache_file.read
      end

      create_listing_block(
        parent,
        content,
        attrs
      )
    else
      block = create_listing_block(
        parent,
        IO.read(file).gsub("\t", '    '),
        attrs.merge({
          'style'    => 'source',
          'language' => 'go',
          'title'    => "#{target}.go",
        })
      )
      block.title = "#{target}.go"
      block.assign_caption
      block
    end
  end
end

# Expands to `go doc` output
#
#   godoc::go/ast.Print[]
class GoDocMacro < Asciidoctor::Extensions::BlockMacroProcessor
  use_dsl

  named :godoc

  def process(parent, target, attrs)
    m = %r<^((?:[\w.]+/)?\w+)\.([\w.]+)$>.match(target)
    opts = attrs.delete(1)

    pkg, entry = m[1], m[2]
    decl = %x(go doc #{opts} #{target}).gsub(/^ {4}.*/m, '').gsub("\t", '    ').lines.map(&:chomp)
    create_listing_block(
      parent,
      decl,
      attrs.merge({
        'style'    => 'source',
        'language' => 'go',
        'title'    => "godoc: http://godoc.org/pkg/#{pkg}##{entry}[#{target}]",
      })
    )
  end
end

Asciidoctor::Extensions.register do
  block_macro GoExampleMacro
  block_macro GoDocMacro

  if @document.basebackend?('html') && ENV['PRODUCTION']
    postprocessor do
      process do |doc, output|
        output.sub('</html>', <<-GA_HTML)
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
  ga('create', 'UA-34276254-7', 'auto');
  ga('send', 'pageview');
</script>
</html>
        GA_HTML
      end
    end
  end
end
