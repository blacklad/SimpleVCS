package lib

import (
	"io/ioutil"
	"strings"

	"github.com/MSathieu/Gotils"
)

//Module is the module object
type Module struct {
	URL  string
	Name string
	Hash string
}

//GetModules gets an array of modules
func GetModules() ([]Module, error) {
	content, err := ioutil.ReadFile(".svcsmodules")
	if err != nil {
		return nil, err
	}
	modulesFile := gotils.PreProcess(string(content))
	var modules []Module
	for _, line := range strings.Split(modulesFile, "\n") {
		if line == "" {
			continue
		}
		split := strings.Fields(line)
		modules = append(modules, Module{Name: split[0], URL: split[1], Hash: split[2]})
	}
	return modules, nil
}
