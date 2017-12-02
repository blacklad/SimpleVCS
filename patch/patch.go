package patch

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

//Read reads a patch
func Read(filename string) (Patch, error) {
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
