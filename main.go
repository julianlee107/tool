package main

import (
	"log"

	"github.com/julianlee107/tool/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.excute err:%v", err)
	}
}
