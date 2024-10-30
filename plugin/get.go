package plugin

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/robertkrimen/otto"
)

func GetAllPluginsInfos() []PluginInfo {
	var i []PluginInfo
	rootPath, _ := os.Getwd()
	filepath.Walk(filepath.Join(rootPath, "data", "plugin"), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".js" {
			vm := otto.New()
			codeText, _ := os.ReadFile(path)
			vm.Run(string(codeText))
			vm.Run(`var info = PluginInfoGet()`)
			var _cnt PluginInfo
			var v otto.Value
			_cnt.EName = strings.ReplaceAll(filepath.Base(path), filepath.Ext(path), "")
			v, _ = vm.Run(`info.name`)
			_cnt.Name, _ = v.ToString()
			v, _ = vm.Run(`info.author`)
			_cnt.Author, _ = v.ToString()
			v, _ = vm.Run(`info.description`)
			_cnt.Description, _ = v.ToString()
			i = append(i, _cnt)
		}
		return nil
	})
	return i
}
