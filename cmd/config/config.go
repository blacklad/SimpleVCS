package config

import (
	"fmt"

	"github.com/MSathieu/SimpleVCS/util"
)

//GetConfig prints the value of a config
func GetConfig(key string) error {
	config := util.GetConfig(key)
	fmt.Println(config)
	return nil
}

//SetConfig sets a config key.
func SetConfig(key string, value string) error {
	util.DB.Where(&util.Config{Name: key}).First(&util.Config{}).Update("value", value)
	return nil
}
