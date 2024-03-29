package checkout

import (
	"os"

	"github.com/MSathieu/SimpleVCS/cmd/initialize"
	"github.com/MSathieu/SimpleVCS/cmd/modules"
	"github.com/MSathieu/SimpleVCS/cmd/pull"
)

func checkoutModules() error {
	modules, err := modules.Get()
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
		initialize.Initialize(module.Name)
		err = pull.Pull(module.URL, os.Getenv("SVCS_MODULE_"+module.Name+"_USERNAME"), os.Getenv("SVCS_MODULE_"+module.Name+"_PASSWORD"))
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
