package lib

// CheckForFastForward checkis if fastforward merge is possible.
func CheckForFastForward(fromBranch string, toBranch string) (bool, error) {
	fromSha, _, err := ConvertToCommit(fromBranch)
	if err != nil {
		return false, err
	}
	toSha, _, err := ConvertToCommit(toBranch)
	if err != nil {
		return false, err
	}
	if toSha == "" || fromSha == "" {
		return false, nil
	}
	for currentSha := fromSha; true; {
		if currentSha == toSha {
			return true, nil
		}
		currentSha, err = GetParent(currentSha)
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

//PerformFastForward performs fastforward merge, before calling this you should call CheckForFastforward.
func PerformFastForward(fromBranch string, toBranch string) error {
	fromSha, _, err := ConvertToCommit(fromBranch)
	if err != nil {
		return err
	}
	err = UpdateBranch(toBranch, fromSha)
	return err
}
