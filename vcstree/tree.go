package vcstree

import (
	"github.com/MSathieu/SimpleVCS/vcsfile"
)

//Tree is the tree object.
type Tree struct {
	Hash  string
	Files []File
}

//File is the object that has a file and a name.
type File struct {
	File vcsfile.File
	Name string
}
