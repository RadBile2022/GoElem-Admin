package api

import (
	"elementary-admin/constants"
	"elementary-admin/src/http/middleware"
	"elementary-admin/src/http/requests"
	"elementary-admin/src/http/usecases"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func InitAdminController(auc usecases.AdminUseCaseImpl) AdminController {
	return &adminController{
		adminUseCase: auc,
		tagError:     "ERROR CONTROLLER -:- ADMIN CONTROLLER",
	}
}

type AdminController interface {
	Groups(r fiber.Router)
}

func (di *adminController) Groups(r fiber.Router) {
	groupRoute := r.Group("/admin")
	groupRoute.Use(middleware.AuthenticationJwt())
	//groupRoute.Get("/", middleware.OnlyOwner(di.Create))
	groupRoute.Delete("/delete/:id", middleware.OnlyOwner(di.Delete))
	groupRoute.Put("/update/:id", middleware.OnlyOwner(di.Update))
	groupRoute.Post("/create", middleware.OnlyOwner(di.Create))
	groupRoute.Get("/", middleware.OnlyOwner(di.GetAll))

	authRoute := r.Group("/auth")
	authRoute.Post("/login", di.Login)
}

func (di *adminController) Login(fc *fiber.Ctx) error {
	var valid, msgValidation, req = requests.ValidateRequestLogin(fc)
	if !valid {
		return fc.Status(constants.CodeErrRequestNotValid).JSON(constants.JsonRequestNotValid(msgValidation))
	} else {
		ctx := fc.Context()
		code, reqAdmin := di.adminUseCase.Login(ctx, req.UsernameAdmin, req.PasswordAdmin)
		return fc.Status(code).JSON(constants.JsonRes(code, reqAdmin))
	}
}
func (di *adminController) Delete(fc *fiber.Ctx) error {
	idAdmin := fc.Params("id")
	ctx := fc.Context()
	idAdminConv, _ := strconv.ParseInt(idAdmin, 10, 64)
	code, reqAdmin := di.adminUseCase.Delete(ctx, idAdminConv)
	return fc.Status(code).JSON(constants.JsonRes(code, reqAdmin))
}
func (di *adminController) Update(fc *fiber.Ctx) error {
	var valid, msgValidation, req = requests.ValidateRequestAdmin(fc)
	idAdmin := fc.Params("id")
	if !valid {
		return fc.Status(constants.CodeErrRequestNotValid).JSON(constants.JsonRequestNotValid(msgValidation))
	} else {
		if idAdmin == "" {
			return fc.Status(constants.CodeErrRequestNotValid).JSON(constants.JsonRequestNotValid(""))
		}
		ctx := fc.Context()
		adminIdConv, _ := strconv.ParseInt(idAdmin, 10, 64)
		code, reqAdmin := di.adminUseCase.Update(ctx, *req, adminIdConv)

		return fc.Status(code).JSON(constants.JsonRes(code, reqAdmin))
	}
}
func (di *adminController) Create(fc *fiber.Ctx) error {
	var valid, msgValidation, req = requests.ValidateRequestAdmin(fc)

	if !valid {
		return fc.Status(constants.CodeErrRequestNotValid).JSON(constants.JsonRequestNotValid(msgValidation))
	} else {
		ctx := fc.Context()
		code, reqAdmin := di.adminUseCase.Create(ctx, *req)
		return fc.Status(code).JSON(constants.JsonRes(code, reqAdmin))
	}
}
func (di *adminController) GetAll(fc *fiber.Ctx) error {
	ctx := fc.Context()
	code, admin := di.adminUseCase.GetAll(ctx)
	return fc.Status(code).JSON(constants.JsonRes(code, admin))

}

type adminController struct {
	adminUseCase usecases.AdminUseCaseImpl
	tagError     string
}
