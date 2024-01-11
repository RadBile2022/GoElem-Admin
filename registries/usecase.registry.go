package registries

import (
	"elementary-admin/config"
	"elementary-admin/src/http/usecases"
	"sync"
)

func InitUseCase(repo RepositoryRegistry, cfg config.Config) UseCaseRegistry {
	var ucr UseCaseRegistry
	var loadOnce sync.Once

	loadOnce.Do(func() {
		ucr = &useCaseRegistry{
			repo: repo,
			cfg:  cfg,
		}
	})
	return ucr
}

type UseCaseRegistry interface {
	Admin() usecases.AdminUseCaseImpl
	// TODO --- write more ---
}

func (di useCaseRegistry) Admin() usecases.AdminUseCaseImpl {
	var ucr usecases.AdminUseCaseImpl
	var loadOnce sync.Once

	loadOnce.Do(func() {
		ucr = usecases.InitAdminUseCase(di.repo.Admin())
	})
	return ucr
}

type useCaseRegistry struct {
	repo RepositoryRegistry
	cfg  config.Config
}
