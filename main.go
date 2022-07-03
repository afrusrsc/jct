package main

import (
	"log"
	"os"

	"github.com/afrusrsc/jct/cmd"
)

func main() {
	if err := cmd.JCT.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
