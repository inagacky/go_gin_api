package di

import (
	co "github.com/inagacky/go_gin_api/src/api/controller"
	s "github.com/inagacky/go_gin_api/src/api/domain/service"
	r "github.com/inagacky/go_gin_api/src/api/infrastructure/repository"
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
	if err := c.Provide(func (repo r.UserRepository) s.UserService {
		return s.NewUserService(repo) }); err != nil {
		return nil, err
	}
	if err := c.Provide(func (service s.UserService) co.UserController {
		return co.NewUserController(service) }); err != nil {
		return nil, err
	}

	return c, nil
}