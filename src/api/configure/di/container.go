package di

import (
	r "github.com/inagacky/go_gin_api/src/api/infrastructure/repository"
	co "github.com/inagacky/go_gin_api/src/api/interface/controllers"
	"github.com/inagacky/go_gin_api/src/api/usecase"
	"go.uber.org/dig"
)

//  DIによるコンテナ化の実施
func Init() (*dig.Container, error) {

	c := dig.New()
	if _, err := userInit(c); err != nil {
		return c, err
	}

	return c, nil
}

// Userインターフェースの初期化
func userInit(c *dig.Container) (*dig.Container, error) {
	if err := c.Provide(func () r.UserRepository {
		return r.NewUserRepository() }); err != nil {
		return nil, err
	}
	if err := c.Provide(func (repo r.UserRepository) usecase.UserService {
		return usecase.NewUserService(repo) }); err != nil {
		return nil, err
	}
	if err := c.Provide(func (service usecase.UserService) co.UserController {
		return co.NewUserController(service) }); err != nil {
		return nil, err
	}

	return c, nil
}
