package requests

import (
	"github.com/gofiber/fiber/v2"
)

// AdminReq TODO
type AdminReq struct {
	IdAdmin       int64  `json:"id_admin"`
	UsernameAdmin string `json:"username_admin"`
	PasswordAdmin string `json:"password_admin"`
	Position      string `json:"position"`
}

// ValidateRequestAdmin TODO
func ValidateRequestAdmin(fc *fiber.Ctx) (valid bool, errorMsg map[string]interface{}, req *AdminReq) {
	valid = true
	errorMsg = map[string]interface{}{}
	req = new(AdminReq)

	if err := fc.BodyParser(&req); err != nil {
		valid = false
		errorMsg["message"] = "Request is not Valid"
		return
	}

	valid = req.UsernameAdmin != ""
	if !valid {
		errorMsg["username_admin"] = "Username must be filled"
	}

	valid = req.PasswordAdmin != ""
	if !valid {
		errorMsg["password_admin"] = "Password must be filled"
	}

	valid = req.Position != ""
	if !valid {
		errorMsg["position"] = "Position must be filled"
	}
	return
}

// ValidateRequestLogin TODO for Login
func ValidateRequestLogin(fc *fiber.Ctx) (valid bool, errorMsg map[string]interface{}, req *AdminReq) {
	valid = true
	errorMsg = map[string]interface{}{}
	req = new(AdminReq)

	if err := fc.BodyParser(&req); err != nil {
		valid = false
		errorMsg["message"] = "Request is not Valid"
		valid = false
		return
	}

	valid = req.UsernameAdmin != ""
	if !valid {
		errorMsg["username_admin"] = "Username must be Filled"
	}

	valid = req.PasswordAdmin != ""
	if !valid {
		errorMsg["password_admin"] = "Password must be Filled"
	}
	return
}
