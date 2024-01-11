package responses

import "elementary-admin/src/domain/repositories/model"

type AdminRes struct {
	Admin    []*AdminPart `json:"admin"`
	Token    string       `json:"token"`
	Position string       `json:"position"`
	User     string       `json:"user"`
	Msg      string       `json:"msg"`
	Success  bool         `json:"success"`
}

type AdminPart struct {
	IdAdmin       int64  `json:"id_admin"`
	UsernameAdmin string `json:"username_admin"`
	Position      string `json:"position"`
}

// AdminArrMysql TODO
func (AdminPart) AdminArrMysql(admins []*model.Admin) (adminPart []*AdminPart) {
	for _, admin := range admins {
		admPart := new(AdminPart)
		admPart = admPart.AdminMysql(*admin)
		adminPart = append(adminPart, admPart)
	}
	return
}

func (AdminPart) AdminMysql(admin model.Admin) (adminPart *AdminPart) {
	adminPart = new(AdminPart)
	adminPart.IdAdmin = admin.IdAdmin
	adminPart.UsernameAdmin = admin.UsernameAdmin
	adminPart.Position = admin.Position
	return
}
