package database

func PluginDataExist(pluginName string, Name string) bool {
	var i int64
	DB.Model(&PluginData{}).Where("plugin = ?", pluginName).Where("name = ?", Name).Count(&i)
	return i != 0
}

func PluginDataSet(pluginName string, Name string, Value string) {
	var i PluginData
	if PluginDataExist(pluginName, Name) {
		DB.Delete(&PluginData{}, "plugin = ? and name = ?", pluginName, Name)
	}
	i.Plugin = pluginName
	i.Name = Name
	i.Value = Value
	DB.Create(&i)
}

func PluginDataGet(pluginName string, Name string) string {
	if !PluginDataExist(pluginName, Name) {
		return ""
	}
	var i PluginData
	DB.Model(&PluginData{}).Where("plugin = ?", pluginName).Where("name = ?", Name).First(&i)
	return i.Value
}
