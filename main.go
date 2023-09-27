package main

import (
	"fmt"
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	spacegame "spacegame/models"
	"golang.org/x/image/colornames"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	// Utilizamos la función pixelgl.Run(run) para iniciar la aplicación gráfica.
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Friends of Go: Space Game",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Creamos una instancia del mundo del juego.
	world := spacegame.NewWorld(windowWidth, windowHeight)
	if err := world.AddBackground("assets/background.png"); err != nil {
		log.Fatal(err)
	}

	// Creamos una instancia del jugador.
	player, err := spacegame.NewPlayer("assets/player.png", 5, world)
	if err != nil {
		log.Fatal(err)
	}

	direction := spacegame.Idle
	action := spacegame.NoneAction
	last := time.Now()

	// Inicia el bucle principal del juego.
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Black)
		world.Draw(win)

		// Captura la entrada del usuario para la dirección y la acción.
		if win.Pressed(pixelgl.KeyLeft) {
			direction = spacegame.LeftDirection
		}

		if win.Pressed(pixelgl.KeyRight) {
			direction = spacegame.RightDirection
		}

		if win.Pressed(pixelgl.KeySpace) {
			action = spacegame.ShootAction
		}

		// Actualiza el jugador con la dirección y la acción.
		player.Update(direction, action, dt)
		player.Draw(win)
		direction = spacegame.Idle
		action = spacegame.NoneAction

		// Calcula y muestra la velocidad de cuadros (FPS).
		fps := 1 / dt
		fmt.Println("FPS: ", int(fps))

		win.Update()
	}
}
