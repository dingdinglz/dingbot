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
