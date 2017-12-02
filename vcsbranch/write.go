package vcsbranch

import "os"

//Write writes the array to branches.txt.
func Write(branches []Branch) error {
	branchesFile, err := os.Create(branchesFile)
	if err != nil {
		return err
	}
	for _, branch := range branches {
		_, err = branchesFile.WriteString(branch.Name + " " + branch.Commit.Hash + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
