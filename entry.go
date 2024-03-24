package main

import (
	"embed"
	"fmt"

	"github.com/hellflame/arithmetic-focus/api"
	"github.com/hellflame/arithmetic-focus/utils"
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
