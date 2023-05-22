package bootstrap

import "github.com/gin-gonic/gin"

type application struct {
	Env    *env
	Client *dbClient
}

var App *application

func init() {
	App = new(application)
	App.Env = newEnv()
	App.Client = new(dbClient)

	if App.Env.Mode == gin.ReleaseMode {
		App.Client.KL = newDBClient(&App.Env.KL)
		App.Client.KC = newDBClient(&App.Env.KC)
		App.Client.KW = newDBClient(&App.Env.KW)
	} else {
		App.Client.KL = newDBClient(&App.Env.KL)
	}
}
