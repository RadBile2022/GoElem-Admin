package usecases

import (
	"context"
	"crypto/md5"
	"elementary-admin/src/domain/repositories/model"
	"elementary-admin/src/domain/repositories/mysql"
	"elementary-admin/src/http/middleware"
	"elementary-admin/src/http/requests"
	"elementary-admin/src/http/responses"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func InitAdminUseCase(adminRepo mysql.AdminRepoImpl) AdminUseCaseImpl {
	return &adminUseCase{
		adminRepo: adminRepo,
		tagError:  "ERROR LOGIC -:- ADMIN USE CASE",
	}
}

type AdminUseCaseImpl interface {
	// TODO --- write more ---

	Login(ctx context.Context, username string, password string) (int, *responses.AdminRes)
	Delete(ctx context.Context, IdAdmin int64) (int, *responses.AdminRes)
	Update(ctx context.Context, req requests.AdminReq, IdAdmin int64) (int, *responses.AdminRes)
	Create(ctx context.Context, req requests.AdminReq) (int, *responses.AdminRes)
	GetAll(ctx context.Context) (int, *responses.AdminRes)
}

func (di *adminUseCase) Login(ctx context.Context, username string, password string) (code int, res *responses.AdminRes) {
	//fmt.Println(hashString)
	adminRepo, code, err := di.adminRepo.Login(username)
	res = new(responses.AdminRes)
	code = fiber.StatusOK

	if err != nil {
		res.Msg = "An Error Occurred, please contact Customer Service"
		code = fiber.StatusInternalServerError
		return
	}

	if len(adminRepo) == 0 {
		res.Msg = "Data Admin is not Found"
		code = fiber.StatusNotFound
		return
	}

	hash := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(hash[:])
	if !strings.EqualFold(hashString, adminRepo[0].PasswordAdmin) {
		res.Msg = "Username and Password is False"
		code = fiber.StatusNotFound
		return
	}

	res.Success = code == fiber.StatusOK
	tokenString, err := middleware.GenerateJwt(adminRepo[0].UsernameAdmin, adminRepo[0].IdAdmin, adminRepo[0].Position)
	res.Msg = "Login is Success"
	res.Token = tokenString
	res.Position = adminRepo[0].Position
	res.User = adminRepo[0].UsernameAdmin
	return
}

func (di *adminUseCase) Delete(ctx context.Context, IdAdmin int64) (code int, res *responses.AdminRes) {
	adminRepo, err := di.adminRepo.Delete(IdAdmin)
	res = new(responses.AdminRes)
	code = fiber.StatusOK
	if err != nil {
		res.Msg = "An Error Occurred, please Contact Customer Service"
		code = fiber.StatusInternalServerError
		return
	}

	res.Msg = "Delete Admin has been Successfully"
	res.Success = adminRepo == fiber.StatusOK
	return
}

func (di *adminUseCase) Update(ctx context.Context, req requests.AdminReq, IdAdmin int64) (code int, res *responses.AdminRes) {
	hash := md5.Sum([]byte(req.PasswordAdmin))
	hashString := hex.EncodeToString(hash[:])

	admin := model.Admin{
		UsernameAdmin: req.UsernameAdmin,
		PasswordAdmin: hashString,
		Position:      req.Position,
	}

	res = new(responses.AdminRes)
	code = fiber.StatusOK
	adminRepo, err := di.adminRepo.Update(admin, IdAdmin)
	if err != nil {
		res.Msg = "An Error Occurred, Please contact Customer Service"
		code = fiber.StatusInternalServerError
		return
	}

	res.Msg = "Update Admin has been Successfully"
	res.Success = adminRepo == fiber.StatusOK
	return
}

func (di *adminUseCase) Create(ctx context.Context, req requests.AdminReq) (code int, res *responses.AdminRes) {
	// REQUEST:
	// Convert the password to hash. Hash to a  hexadecimal string
	hash := md5.Sum([]byte(req.PasswordAdmin))
	hashString := hex.EncodeToString(hash[:])

	admin := model.Admin{
		UsernameAdmin: req.UsernameAdmin,
		PasswordAdmin: hashString,
		Position:      req.Position,
	}

	adminRepo, err := di.adminRepo.Create(admin)

	// RESPONSE :
	res = new(responses.AdminRes)
	code = fiber.StatusOK

	if err != nil {
		res.Msg = "An Error Occurred, Please contact Customer Service"
		code = fiber.StatusInternalServerError
		return
	}

	res.Msg = "Save Admin has been Successfully"
	res.Success = adminRepo == fiber.StatusOK
	return
}

func (di *adminUseCase) GetAll(ctx context.Context) (code int, res *responses.AdminRes) {
	res = new(responses.AdminRes)
	code = fiber.StatusOK
	adminRepo, code, err := di.adminRepo.GetAll()
	if err != nil {
		res.Msg = "An Error Occurred, Please contact Customer Service"
		code = fiber.StatusInternalServerError
		return
	}

	if len(adminRepo) == 0 {
		res.Msg = "Data Admin is not Found"
		code = fiber.StatusNotFound
		return
	}

	adminPart := responses.AdminPart{}
	adminResponse := adminPart.AdminArrMysql(adminRepo)

	res.Admin = adminResponse
	res.Msg = "Admins have been Found"
	res.Success = code == fiber.StatusOK
	return
}

type adminUseCase struct {
	adminRepo mysql.AdminRepoImpl
	tagError  string
}
