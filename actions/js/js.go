package js

import (
	"strings"

	"github.com/Mrpye/go-workflow/workflow"
	"github.com/Mrpye/golib/lib"
)

func Action_RunJS(w *workflow.Workflow) error {

	code, err := w.GetConfigTokenString("js", w.Model, false)
	if err != nil {
		return err
	}

	if code == "" {
		//*********************
		//Get the config values
		//*********************
		js_file, err := w.GetConfigTokenString("js_file", w.Model, true)
		if err != nil {
			return err
		}

		//*****************************
		//See if we have multiple files
		//*****************************
		files := strings.Split(js_file, ";")

		for _, o := range files {
			file_part := o
			file_data, err := lib.ReadFileToString(file_part)
			if err != nil {
				return err
			}
			code = code + file_data + "\n"
		}
	}
	//************
	//Run our code
	//************
	vm := w.CreateJSEngine()
	vm.Set("model", w.Model)
	_, err = vm.RunString(code)
	if err != nil {
		return err
	}
	return nil
}
