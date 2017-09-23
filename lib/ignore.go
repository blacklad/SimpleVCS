package lib

//CheckIgnored checks if the file/directory must be ignored.
func CheckIgnored(file string) (bool, error) {
	switch file {
	case ".svcs", ".git", ".svn", ".hg":
		return true, nil
	}
	return false, nil
}
