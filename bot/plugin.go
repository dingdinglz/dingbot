package bot

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/robertkrimen/otto"
)

func RunPlugin(fRun func(*otto.Otto)) {
	rootPath, _ := os.Getwd()
	filepath.Walk(filepath.Join(rootPath, "data", "plugin"), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".js" {
			vm := otto.New()
			AddVMFuncs(vm)
			codeText, _ := os.ReadFile(path)
			vm.Run(string(codeText))
			fRun(vm)
		}
		return nil
	})
}
