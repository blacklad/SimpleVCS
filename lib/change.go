package lib

import (
	"strings"

	"github.com/MSathieu/SimpleVCS/vcstree"
)

//Change is the change object.
type Change struct {
	Type string
	Name string
	Hash string
}

//GenerateChange returns the changes between two file arrays.
func GenerateChange(fromFiles []vcstree.TreeFile, toFiles []vcstree.TreeFile) []Change {
	var changes []Change
	for _, fromFile := range fromFiles {
		change := Change{Type: "deleted", Name: fromFile.Name}
		changed := true
		for _, toFile := range toFiles {
			if fromFile.Name == toFile.Name {
				if fromFile.File.Hash == toFile.File.Hash {
					changed = false
				} else {
					change.Type = "changed"
					change.Hash = toFile.File.Hash
				}
			}
		}
		if changed {
			changes = append(changes, change)
		}
	}
	for _, toFile := range toFiles {
		change := Change{Type: "created", Name: toFile.Name, Hash: toFile.File.Hash}
		changed := true
		for _, fromFile := range fromFiles {
			if toFile.Name == fromFile.Name {
				changed = false
			}
		}
		if changed {
			changes = append(changes, change)
		}
	}
	return changes
}

//ApplyChange applies a changeset to a fileset.
func ApplyChange(files []string, changes []Change) []string {
	for _, change := range changes {
		switch change.Type {
		case "created":
			files = append(files, change.Name+" "+change.Hash)
		case "changed":
			for i := range files {
				fileMapping := strings.Split(files[i], " ")
				if fileMapping[0] == change.Name {
					files[i] = change.Name + " " + change.Hash
				}
			}
		case "deleted":
			for i := range files {
				fileMapping := strings.Split(files[i], " ")
				if fileMapping[0] == change.Name {
					files = append(files[:i], files[i+1:]...)
					break
				}
			}
		}
	}
	return files
}
