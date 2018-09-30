package main

import "os"
import "log"

import "github.com/mandusm/padlock-cloud/padlockcloud"

func main() {
	err := padlockcloud.NewCliApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
