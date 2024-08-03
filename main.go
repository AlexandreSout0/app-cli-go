package main

import (
	"app-cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("init app")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
