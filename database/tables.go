package database

type KeyWordTable struct {
	Key     string
	Text    string
	Precise string
}

func (KeyWordTable) TableName() string {
	return "keyword"
}
