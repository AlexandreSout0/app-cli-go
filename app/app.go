package app

import "github.com/urfave/cli"

// func Gerar() - Retorna a aplicação CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI - App"
	app.Usage = "Search IP's and Server on internet"

	return app
}
