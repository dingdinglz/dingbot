package bot

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/database"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/robertkrimen/otto"
)

func AddVMFuncs(Vm *otto.Otto) {
	Vm.Set("SendGroupMessage", func(call otto.FunctionCall) otto.Value {
		DingQQBot.SendGroupMessage(tool.StringToUint32(call.Argument(0).String()), []message.IMessageElement{message.NewText(call.Argument(1).String())})
		return otto.Value{}
	})
	Vm.Set("SendPrivateMessage", func(call otto.FunctionCall) otto.Value {
		DingQQBot.SendPrivateMessage(tool.StringToUint32(call.Argument(0).String()), []message.IMessageElement{message.NewText(call.Argument(1).String())})
		return otto.Value{}
	})
	Vm.Set("DingbotPrint", func(call otto.FunctionCall) otto.Value {
		fmt.Println("[plugin] " + call.Argument(0).String())
		return otto.Value{}
	})
	Vm.Set("DataSet", func(call otto.FunctionCall) otto.Value {
		database.PluginDataSet(call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String())
		return otto.Value{}
	})
	Vm.Set("DataGet", func(call otto.FunctionCall) otto.Value {
		v, _ := Vm.ToValue(database.PluginDataGet(call.Argument(0).String(), call.Argument(1).String()))
		return v
	})
}

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
