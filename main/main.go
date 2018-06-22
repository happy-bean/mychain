package main

import "mychain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("A")
	bc.SendData("B")
	bc.Print()
}
