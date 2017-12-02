package types

//Tree is the tree object.
type Tree struct {
	Hash  string
	Files []TreeFile
}

//TreeFile is the object that has a file and a name.
type TreeFile struct {
	File File
	Name string
}
