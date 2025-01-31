package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// func Gerar() - Retorna a aplicação CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI - App"
	app.Usage = "Search IP's and Server on internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca de IP's na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "https://alexandresout0.github.io/homepage/",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
