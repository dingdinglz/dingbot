package database

type KeyWordTable struct {
	Key     string
	Text    string
	Precise string
}

func (KeyWordTable) TableName() string {
	return "keyword"
}

type OpenTable struct {
	Uin      uint32
	Typename string
}

func (OpenTable) TableName() string {
	return "open"
}

type StopTable struct {
	Text string
}

func (StopTable) TableName() string {
	return "stop"
}

type MessageTable struct {
	Typename string
	Message  string
	Seq      uint32
}

func (MessageTable) TableName() string {
	return "message"
}

type PluginData struct {
	Plugin string
	Name   string
	Value  string
}

func (PluginData) TableName() string {
	return "plugin"
}

type StarTable struct {
	Repository string
	Sender     string
}

func (StarTable) TableName() string {
	return "star"
}

type GithubHookTable struct {
	ID    uint
	Name  string
	Group string
}

func (GithubHookTable) TableName() string {
	return "github"
}
