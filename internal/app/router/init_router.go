package router

import (
	"net/http"

	"developer-orientenergy-golang/internal/pkg/config"
	"developer-orientenergy-golang/internal/pkg/database"
)

func InitRouterMapping() {
	Dispatch = make(map[string]http.HandlerFunc)

	info := config.NewConfigDatabase()
	_, err := database.NewConnection(database.Config{
		Host:     info.Host,
		Port:     info.Port,
		User:     info.User,
		Password: info.Password,
		Database: info.Database,
	})
	if err != nil {
		print(err)
	}
	Dispatch = AuthRouter(Dispatch)
	Dispatch = SchoolRouter(Dispatch)

	//add bookRouter
	Dispatch = BookRouter(Dispatch)
}
