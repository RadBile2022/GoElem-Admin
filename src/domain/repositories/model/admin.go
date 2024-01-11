package model

import (
	"elementary-admin/constants"
	"encoding/json"
	"log"
)

type Admin struct {
	IdAdmin       int64  `json:"id_admin"`
	UsernameAdmin string `json:"username_admin"`
	PasswordAdmin string `json:"password_admin"`
	Position      string `json:"position"`
}

// AdminNullable TODO for login
type AdminNullable struct {
	IdAdmin       constants.NullInt64  `json:"id_admin"`
	UsernameAdmin constants.NullString `json:"username_admin"`
	PasswordAdmin constants.NullString `json:"password_admin"`
	Position      constants.NullString `json:"position"`
}

func (di *AdminNullable) ToAdmin() (admin Admin, error error) {
	buf, err := json.Marshal(di)
	if err != nil {
		log.Println(err)
		return
	}

	admin, err = admin.FromJsonString(string(buf))
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (di *Admin) FromJsonString(jsonStr string) (Admin, error) {
	admin := *di
	if err := json.Unmarshal([]byte(jsonStr), &admin); err != nil {
		return admin, err
	}
	return admin, nil
}
