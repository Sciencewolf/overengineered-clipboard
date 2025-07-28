package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/getlantern/systray"
)

func onReady() {

}

func onExit() {
	systray.Quit()
}

func main() {
	systray.Run(onReady, onExit)
}
