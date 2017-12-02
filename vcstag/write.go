package vcstag

import "os"

//Write writes an array to tags.txt.
func Write(tags []Tag) error {
	tagsFile, err := os.Create(tagsFile)
	if err != nil {
		return err
	}
	for _, tag := range tags {
		_, err = tagsFile.WriteString(tag.Name + " " + tag.Commit.Hash + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
