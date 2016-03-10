require 'asciidoctor'
require 'asciidoctor/extensions'
require 'pathname'

CACHE_DIR = Pathname.new('.cache')

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

class GoDocMacro < Asciidoctor::Extensions::BlockMacroProcessor
  use_dsl

  named :godoc

  def process(parent, target, attrs)
    m = %r<^((?:[\w.]+/)?\w+)\.([\w.]+)$>.match(target)
    pkg, entry = m[1], m[2]
    decl = %x(go doc #{target}).gsub("\t", '    ').lines.map(&:chomp).take_while { |line| not line.empty? }
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
end
