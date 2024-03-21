package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
)

func main() {
	// Clone the repository
	r, err := git.PlainOpen("/Users/huanhuansun/github/yaml-test")
	if err != nil {
		panic(err)
	}
	// Get the HEAD reference
	ref, err := r.Head()
	if err != nil {
		log.Fatalf("Failed to get HEAD: %s", err)
	}

	// Get the commit object from the HEAD reference
	headCommit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatalf("Failed to get commit object for HEAD: %s", err)
	}

	fmt.Printf("HEAD commit: %s\n", headCommit)

	// Now to get the previous commit, we can use the parents
	// Note: This assumes the current HEAD is not the initial commit
	iter := headCommit.Parents()
	defer iter.Close()

	// Get the first parent (previous commit)
	prevCommit, err := iter.Next()
	if err != nil {
		log.Fatalf("Failed to get previous commit: %s", err)
	}
	// Get the diff between the previous and current commit
	diff, err := prevCommit.Patch(headCommit)
	if err != nil {
		panic(err)
	}
	fmt.Println(prevCommit.Hash.String())
	fmt.Println(headCommit.Hash.String())
	for _, file := range diff.FilePatches() {
		from, to := file.Files()
		if from != nil {
			fmt.Println(from.Path())
		} else {
			fmt.Println("added")
		}
		if to != nil {
			fmt.Println(to.Path())
		} else {
			fmt.Println("deleted")
		}
		for _, chunk := range file.Chunks() {
			fmt.Printf("\n%v\n%s", chunk.Type(), chunk.Content())
			fmt.Println()
		}
	}
}
