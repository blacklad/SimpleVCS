package lib

import (
	"os"

	"github.com/MSathieu/Gotils"
)

//Patch is the patch object
type Patch struct {
	FromHash string
	Changes  []string
}

func (patch Patch) Save(filename string) error {
	patchContent := "parent " + patch.FromHash + "\n"
	for _, change := range patch.Changes {
		patchContent = patchContent + change + "\n"
	}
	patchContent = gotils.GZip(patchContent)
	file, err := os.Create(filename + ".patch")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(patchContent)
	return err
}
