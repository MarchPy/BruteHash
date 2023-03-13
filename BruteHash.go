package main

import (
	core "broken/Core"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := flags_prog()
	if erro := app.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

func flags_prog() *cli.App {
	app := cli.NewApp()
	app.Name = "Broken HASH"
	app.Usage = "Descubra a senha a partir do hash"

	//Definindo flags
	flags_bruteforce := []cli.Flag{
		cli.StringFlag{
			Name:  "w",
			Usage: "Wordlist",
		},
		cli.StringFlag{
			Name:  "t",
			Usage: "Hash alvo",
		},
		cli.StringFlag{
			Name:  "type",
			Usage: "Tipo do Hash alvo",
		},
	}

	flags_identifier := []cli.Flag{
		cli.StringFlag{
			Name:  "t",
			Usage: "Hash alvo",
		},
	}

	// Definindo os comandos
	app.Commands = []cli.Command{
		{
			Name:   "bruteforce",
			Usage:  "Brute Force de HASH",
			Flags:  flags_bruteforce,
			Action: run_bruteforce,
		},
		{
			Name:   "indentifier",
			Usage:  "Identifique o tipo de HASH",
			Flags:  flags_identifier,
			Action: run_indentifier,
		},
	}

	return app
}

func run_bruteforce(c *cli.Context) {
	// Exigindo os dois parametros
	if !c.IsSet("w") || !c.IsSet("t") || !c.IsSet(("type")) {
		cli.ShowSubcommandHelp(c)
		log.Fatalf("Erro: é necessário fornecer o valor das flags -w, -t e --type")
	} else {
		core.Bruteforce(
			c.String("w"),
			c.String("type"),
			c.String("t"),
		)
	}
}

func run_indentifier(c *cli.Context) {

	if !c.IsSet("t") {
		cli.ShowSubcommandHelp(c)
		log.Fatalf("Erro: é necessário fornecer o valor das flags -t")
	} else {
		core.Indentifier(c.String("t"))
	}
}
