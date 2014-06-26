package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)

	// Parse command line arguments.
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		log.Fatal("path required")
	}

	// Iterate over the tree and process files importing boltdb/raw.
	if err := filepath.Walk(root, walk); err != nil {
		log.Fatal(err)
	}
}

// Walk recursively iterates over all files in a directory and processes any
// file that imports "github.com/boltdb/raw".
func walk(path string, info os.FileInfo, err error) error {
	if info == nil {
		return fmt.Errorf("file not found: %s", err)
	} else if info.IsDir() {
		return nil
	} else if filepath.Ext(path) != ".go" {
		return nil
	}

	// TODO: Parse Go file for imports only and check if imports contain boltdb/raw.

	// Process each file.
	if err := process(path); err != nil {
		return err
	}

	return nil
}

func process(path string) error {
	// TODO: Parse Go file.
	// TODO: Remove code between begin/end comments.
	// TODO: Codegen exported struct.
	// TODO: Codegen encoding.
	// TODO: Codegen decoding.
	// TODO: Pretty print file to buffer.
	// TODO: Rewrite file.
}
