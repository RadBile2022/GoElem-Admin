package registries

import (
	"elementary-admin/config"
	"elementary-admin/src/domain/repositories/mysql"
	"sync"
)

func InitRepository(cfg config.Config) RepositoryRegistry {
	var repoRegistry RepositoryRegistry
	var loadOnce sync.Once

	loadOnce.Do(func() {
		repoRegistry = &repositoryRegistry{cfg: cfg}
	})

	return repoRegistry
}

type RepositoryRegistry interface {
	Admin() mysql.AdminRepoImpl
	// TODO --- write more ---
}

func (r repositoryRegistry) Admin() mysql.AdminRepoImpl {
	var adminRepo mysql.AdminRepoImpl
	var loadOnce sync.Once

	loadOnce.Do(func() {
		adminRepo = mysql.InitAdminRepository(r.cfg)
	})
	return adminRepo
}

type repositoryRegistry struct {
	cfg config.Config
}
