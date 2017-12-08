package util

//Head is the head object.
type Head struct {
	Branch   string
	Detached bool
}

//GetHead returns the head.
func GetHead() Head {
	config := &Config{}
	DB.Where(&Config{Name: "head"}).First(config)
	headObj := Head{Detached: false}
	if config.Value == "DETACHED" {
		headObj.Detached = true
	} else {
		headObj.Branch = config.Value
	}
	return headObj
}
