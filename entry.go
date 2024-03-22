package main

import (
	"arithmetic/api"
	"arithmetic/utils"
	"embed"
	"fmt"
)

//go:embed pages
var pages embed.FS

func main() {
	args := utils.ParseArgs()
	if args != nil {
		address := args.GetAddress()
		if !args.NoBrowser {
			utils.OpenExplore(fmt.Sprintf("http://%s", address))
		}
		api.RegisterAPI(&pages).Start(address)
	}
}
