package infra

import (
	"gorm.io/gorm"
	"projectIraUw/infra/database"
	"projectIraUw/infra/database/postgresql"
	"projectIraUw/user_case/handler"
	"projectIraUw/user_case/repository"
	"projectIraUw/user_case/service"
)

type ContainerDI struct {
	Config     Config
	PostgresDB *gorm.DB
	Repository repository.ListRepository
	Service    service.ListService
	Handler    handler.ListHandler
}

func NewContainerDI(config Config) *ContainerDI {
	container := &ContainerDI{
		Config: config,
	}

	container.buildDB()
	container.buildRepository()
	container.buildService()
	container.buildHandler()

	return container
}

func (c *ContainerDI) buildDB() {
	configPostgresDB := database.Config{
		Hostname: c.Config.PostgresHost,
		Port:     c.Config.PostgresPort,
		UserName: c.Config.PostgresUser,
		Password: c.Config.PostgresPassword,
		Database: c.Config.PostgresDatabase,
	}

	c.PostgresDB = postgresql.InitGorm(&configPostgresDB)
}
func (c *ContainerDI) buildRepository() {
	c.Repository = repository.NewListRepository(c.PostgresDB)
}
func (c *ContainerDI) buildService() {
	c.Service = service.NewListService(c.Repository)
}
func (c *ContainerDI) buildHandler() {
	c.Handler = handler.NewListHandler(c.Service)
}
func (c *ContainerDI) ShutDown() {}
