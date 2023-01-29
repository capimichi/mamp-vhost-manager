package main

import (
	"embed"
	"mamp-vhosts-manager/controller"
	"mamp-vhosts-manager/helper"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	defaultController := &controller.DefaultController{}
	vhostController := &controller.VhostController{}
	mampHelper := &helper.MampHelper{}

	mampHelper.InitializeVhosts()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Mamp vhost manager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        defaultController.Startup,
		Bind: []interface{}{
			defaultController,
			vhostController,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
