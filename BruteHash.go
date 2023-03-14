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
	flags_bruteforce1 := []cli.Flag{
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

	flags_bruteforce2 := []cli.Flag{
		cli.StringFlag{
			Name:  "t",
			Usage: "Hash alvo",
		},
		cli.StringFlag{
			Name:  "type",
			Usage: "Tipo do Hash alvo",
		},
		cli.StringFlag{
			Name:  "min",
			Usage: "Minimo de caracteres",
		},
		cli.StringFlag{
			Name:  "max",
			Usage: "Máximo de caracteres",
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
			Name:   "bruteforce1",
			Usage:  "Brute Force de HASH com wordlist",
			Flags:  flags_bruteforce1,
			Action: run_bruteforce1,
		},
		{
			Name:   "bruteforce2",
			Usage:  "Brute Force de HASH com gerador de senhas",
			Flags:  flags_bruteforce2,
			Action: run_bruteforce2,
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

func run_bruteforce1(c *cli.Context) {
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

func run_bruteforce2(c *cli.Context) {
	// Exigindo os dois parametros
	if !c.IsSet("t") || !c.IsSet(("type")) || !c.IsSet("min") || !c.IsSet("max") {
		cli.ShowSubcommandHelp(c)
		log.Fatalf("Erro: é necessário fornecer o valor das flags -t, ---min, --max e --type")
	} else {
		core.Run_generator(
			c.String("t"),
			c.String("type"),
			c.Int("min"),
			c.Int("max"),
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
