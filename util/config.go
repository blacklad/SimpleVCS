package util

//GetConfig gets the config.
func GetConfig(key string) string {
	config := &Config{}
	DB.Where(&Config{Name: key}).First(config)
	return config.Value
}

//Initialized checks whether the repo is initialized
func Initialized() bool {
	config := &Config{}
	DB.Where(&Config{Name: "name"}).First(config)
	if config.Value == "" {
		return false
	}
	return true
}
