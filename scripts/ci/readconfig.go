package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Versions map[string]string
	Sources  map[string]string
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(filepath.Base(os.Args[0]))

	var conf Config
	err := json.NewDecoder(os.Stdin).Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "version":
		fmt.Print(conf.Versions[os.Args[2]])

	case "sources":
		for pkg, rev := range conf.Sources {
			fmt.Printf("%s\t%s\n", pkg, rev)
		}

	default:
		log.Fatalf("unknown key %q", os.Args[1])
	}
}
