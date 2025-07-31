package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

var windowChannel = make(chan struct{})

func main() {
	go func() {
		systray.Run(onReady, func() {})
	}()

	myApp := app.NewWithID("com.sciencewolf.overengineered-clipboard")

	go func() {
		for range windowChannel {
			w := myApp.NewWindow("Copied Items")
			w.Resize(fyne.NewSize(400, 200))
			w.Show()
		}
	}()

	myApp.Run()
}

func onReady() {
	systray.SetIcon(icon.Data)

	exitItem := systray.AddMenuItem("Exit", "")
	showItem := systray.AddMenuItem("Show", "")

	go func() {
		for {
			select {
			case <-exitItem.ClickedCh:
				systray.Quit()
			case <-showItem.ClickedCh:
				windowChannel <- struct{}{}
			}
		}
	}()
}
