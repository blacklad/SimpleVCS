package lib

import (
	"io/ioutil"
	"os"
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
		return nil, nil
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

//InitModules initializes the modules
func InitModules() error {
	modules, err := GetModules()
	if err != nil {
		return err
	}
	for _, module := range modules {
		err = os.MkdirAll(module.Name, 700)
		if err != nil {
			return err
		}
		err = os.Chdir(module.Name)
		if err != nil {
			return err
		}
		Init(module.Name, true, false)
		err = Pull(module.URL, os.Getenv("SVCS_MODULE_"+module.Name+"_USERNAME"), os.Getenv("SVCS_MODULE_"+module.Name+"_PASSWORD"))
		if err != nil {
			return err
		}
		err = Checkout(module.Hash, false)
		if err != nil {
			return err
		}
		err = os.Chdir("..")
		if err != nil {
			return err
		}
	}
	return nil
}
