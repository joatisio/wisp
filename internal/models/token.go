package models

import "encoding/json"

type Token struct {
	ID ID `gorm:"type:uuid;default:uuid_generate_v4()" json:"-"`

	UserId User `json:"user_id"`

	Access  string `json:"access"`
	Refresh string `json:"refresh"`

	Blocked int `gorm:"type:tinyint(1);default:0"`
}

func JsonToToken(j string) (t *Token, err error) {
	err = json.Unmarshal([]byte(j), t)
	return
}

func (t *Token) Json() (data []byte, err error) {
	data, err = json.Marshal(t)
	return
}
