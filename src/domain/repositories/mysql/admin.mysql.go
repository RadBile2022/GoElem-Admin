package mysql

import (
	"context"
	"database/sql"
	"elementary-admin/config"
	"elementary-admin/constants"
	"elementary-admin/src/domain/repositories/model"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

func InitAdminRepository(cfg config.Config) AdminRepoImpl {
	return &adminRepo{
		cfg:      cfg,
		sqlDb:    cfg.Mysql().GetConnection(),
		tagError: "ERROR DB - ADMIN REPOSITORY -",
	}
}

type AdminRepoImpl interface {
	// TODO --- write more func parent and override ---

	Login(username string) ([]*model.Admin, int, error)
	Delete(IdAdmin int64) (int, error)
	Update(admin model.Admin, idAdmin int64) (int, error)
	Create(admin model.Admin) (int, error)
	GetAll() ([]*model.Admin, int, error)
}

func (di *adminRepo) Login(username string) (admins []*model.Admin, code int, error error) {
	db := di.sqlDb
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stmt := "SELECT " + strings.Join(colAdmin, ",") + " FROM admin WHERE username_admin =?"
	rows, err := db.QueryContext(ctx, stmt, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			admins = nil
			code = fiber.StatusNotFound
			error = err
			return
		}

		admins = nil
		code = fiber.StatusBadRequest
		error = err
		return
	}

	for rows.Next() {
		admin := model.Admin{}
		adminNull := new(model.AdminNullable)
		err = rows.Scan(
			&adminNull.IdAdmin,
			&adminNull.UsernameAdmin,
			&adminNull.PasswordAdmin,
			&adminNull.Position)

		if err != nil {
			code = fiber.StatusInternalServerError
			fmt.Println("err -:-", err.Error())
			return
		}

		admin, err = adminNull.ToAdmin()
		admins = append(admins, &admin)
	}

	code = fiber.StatusOK
	error = nil
	return
}
func (di *adminRepo) Delete(IdAdmin int64) (code int, error error) {
	db := di.sqlDb
	stmt := "DELETE FROM admin WHERE id_admin = ?"

	_, err := db.Exec(stmt, IdAdmin)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	return fiber.StatusOK, nil
}

func (di *adminRepo) Update(admin model.Admin, idAdmin int64) (code int, error error) {
	db := di.sqlDb
	stmt := "UPDATE admin SET username_admin = ?, password_admin =?, position =? WHERE id_admin =?"
	_, err := db.Exec(stmt, admin.UsernameAdmin, admin.PasswordAdmin, admin.Position, idAdmin)

	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (di *adminRepo) Create(admin model.Admin) (code int, error error) {
	db := di.sqlDb
	//stmt := "INSERT INTO admin (username_admin, password_admin, position) VALUES (?,?,?)"
	stmt := "INSERT INTO admin (" + strings.Join(colAdminNoId, ",") + ") VALUES (?,?,?)"
	_, err := db.Exec(stmt, admin.UsernameAdmin, admin.PasswordAdmin, admin.Position)

	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

var colAdminNoId = []string{
	"username_admin",
	"password_admin",
	"position",
}
var colAdmin = []string{
	"id_admin",
	"username_admin",
	"password_admin",
	"position",
}

func (di *adminRepo) GetAll() (admins []*model.Admin, code int, error error) {
	db := di.sqlDb
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stmt := "SELECT * FROM admin ORDER BY username_admin ASC"
	rows, err := db.QueryContext(ctx, stmt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			admins = nil
			code = constants.CodeErrDataNotFound
			error = err
			return
		}

		admins = nil
		code = fiber.StatusBadRequest
		error = err
		return
	}
	for rows.Next() {
		var admin model.Admin
		scan := rows.Scan(
			&admin.IdAdmin,
			&admin.UsernameAdmin,
			&admin.PasswordAdmin,
			&admin.Position,
		)

		if scan != nil {
			code = fiber.StatusInternalServerError
		}

		admins = append(admins, &admin)
	}

	code = fiber.StatusOK
	error = nil
	return
}

type adminRepo struct {
	cfg      config.Config
	sqlDb    *sql.DB
	tagError string
}
