package lib

// CheckForFastForward checkis if fastforward merge is possible.
func CheckForFastForward(fromBranch string, toBranch string) (bool, error) {
	fromCommit, _, err := ConvertToCommit(fromBranch)
	if err != nil {
		return false, err
	}
	toCommit, _, err := ConvertToCommit(toBranch)
	if err != nil {
		return false, err
	}
	if toCommit.Hash == "" || fromCommit.Hash == "" {
		return false, nil
	}
	for currentCommit := fromCommit; true; {
		if currentCommit.Hash == toCommit.Hash {
			return true, nil
		}
		if err != nil {
			return false, err
		}
		parentCommit, err := GetCommit(currentCommit.Parent)
		if err != nil {
			return false, err
		}
		currentCommit = parentCommit
	}
	return false, nil
}

//PerformFastForward performs fastforward merge, before calling this you should call CheckForFastforward.
func PerformFastForward(fromBranch string, toBranch string) error {
	fromSha, _, err := ConvertToCommit(fromBranch)
	if err != nil {
		return err
	}
	err = UpdateBranch(toBranch, fromSha.Hash)
	return err
}
