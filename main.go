package main

import (
	"gochain/Blockchain"
	"gochain/CLI"
)

func main() {
	bc := Blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := CLI.CLI{bc}
	cli.Run()
}
