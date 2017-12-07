package types

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/MSathieu/Gotils"
)

//Patch is the patch object
type Patch struct {
	FromHash string
	Changes  []string
}

//Save saves the patch
func (patch Patch) Save(filename string) error {
	patchContent := "parent " + patch.FromHash + "\n"
	for _, change := range patch.Changes {
		patchContent = patchContent + change + "\n"
	}
	patchContent = gotils.GZip(patchContent)
	patchFile, err := os.Create(filename + ".patch")
	if err != nil {
		return err
	}
	defer patchFile.Close()
	_, err = patchFile.WriteString(patchContent)
	return err
}

//ReadPatch reads a patch
func ReadPatch(filename string) (Patch, error) {
	contentBytes, err := ioutil.ReadFile(filename + ".patch")
	if err != nil {
		return Patch{}, err
	}
	content := gotils.UnGZip(string(contentBytes))
	patch := Patch{Changes: []string{}}
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		if split[0] == "parent" {
			patch.FromHash = split[1]
			continue
		}
		patch.Changes = append(patch.Changes, line)
	}
	return patch, nil
}

//GeneratePatch generates a patch
func GeneratePatch(fromSha string, toSha string, filename string) error {
	fromCommit, err := GetCommit(fromSha)
	if err != nil {
		return err
	}
	toCommit, err := GetCommit(toSha)
	if err != nil {
		return err
	}
	changes := GenerateChange(fromCommit.Tree.Files, toCommit.Tree.Files)
	patchObj := Patch{FromHash: fromSha, Changes: []string{}}
	for _, change := range changes {
		changedFile, err := GetFile(change.Hash)
		if err != nil {
			return err
		}
		patchObj.Changes = append(patchObj.Changes, change.Type+" "+change.Name+" "+gotils.Encode(changedFile.Content))
	}
	err = patchObj.Save(filename)
	return err
}

//ApplyPatch applies a path.
func ApplyPatch(filename string) error {
	patchObj, err := ReadPatch(filename)
	if err != nil {
		return err
	}
	fromCommit, err := GetCommit(patchObj.FromHash)
	if err != nil {
		return err
	}
	var changes []Change
	for _, change := range patchObj.Changes {
		split := strings.Split(change, " ")
		hash := ""
		if split[0] != "deleted" {
			content, err := gotils.Decode(split[2])
			if err != nil {
				return err
			}
			hash = gotils.GetChecksum(content)
			fileObj := File{Content: content, Hash: hash}
			err = fileObj.Save()
			if err != nil {
				return err
			}
		}
		changes = append(changes, Change{Type: split[0], Name: split[1], Hash: hash})
	}
	files := ApplyChange(fromCommit.GetFiles(), changes)
	commitHash, err := CreateCommit("Applied patch "+filename, files)
	if err != nil {
		return err
	}
	return CreateBranch("patch-"+filename, commitHash)
}
