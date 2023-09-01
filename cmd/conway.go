package main

import (
	"log"

	"os"
	"pafaul/conway-go/internal/simulation"

	"github.com/gdamore/tcell"
)

const FPS = 2 

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	game := simulation.NewGame(screen)

	game.AddBlinker(0, 0)

	game.AddBeehive(4, 4)

	game.AddBeacon(10, 10)

	stopChannel := make(chan bool)

	go game.Run(stopChannel, FPS)
	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Rune() == 'p' {
				if game.Running {
					stopChannel <- true
				}
			} else if event.Rune() == 'u' {
				if !game.Running {
					go game.Run(stopChannel, FPS)
				}
			}
		}
	}

}
