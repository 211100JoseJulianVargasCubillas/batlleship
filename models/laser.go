package spacegame

import (
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
)

type Laser struct {
	pic       pixel.Picture
	sfxPath   string
	pos       *pixel.Vec
	vel       float64
	sprite    *pixel.Sprite
	isVisible bool
	world     *World
}

// NewBaseLaser crea una nueva instancia de Laser con información básica.
func NewBaseLaser(path, sfxPath string, vel float64, world *World) (*Laser, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}

	return &Laser{
		pic:     pic,
		vel:     vel,
		world:   world,
		sfxPath: sfxPath,
	}, nil
}

// NewLaser crea una nueva instancia de Laser en una posición específica.
func (l *Laser) NewLaser(pos pixel.Vec) *Laser {
	spr := pixel.NewSprite(l.pic, l.pic.Bounds())

	return &Laser{
		pos:       &pos,
		vel:       l.vel,
		sprite:    spr,
		isVisible: true,
		world:     l.world,
		sfxPath:   l.sfxPath,
	}
}

// Draw dibuja el Laser en la pantalla si es visible.
func (l Laser) Draw(t pixel.Target) {
	if l.isVisible == true {
		l.sprite.Draw(t, pixel.IM.Moved(*l.pos))
	}
}

// Update actualiza la posición del Laser y lo oculta cuando sale de la pantalla.
func (l *Laser) Update() {
	l.pos.Y += l.vel
	if l.pos.Y > l.world.height {
		l.isVisible = false
	}
}

// Shoot reproduce el sonido del Laser y espera a que la reproducción termine.
func (l Laser) Shoot() {
	sfx, err := loadSound(l.sfxPath)
	if err != nil {
		log.Fatal(err)
	}

	// Inicializa el sistema de audio con la tasa de muestreo del sonido.
	speaker.Init(sfx.format.SampleRate, sfx.format.SampleRate.N(time.Second/10))

	// Cierre diferido del streamer de sonido.
	defer sfx.streamer.Close()

	// Crea un canal 'done' para señalar cuando la reproducción de sonido ha terminado.
	done := make(chan bool)

	// Reproduce el sonido utilizando un secuenciador de sonido y una función de devolución de llamada.
	speaker.Play(beep.Seq(sfx.streamer, beep.Callback(func() {
		done <- true
	})))

	// Espera hasta que la señal 'done' indique que la reproducción de sonido ha finalizado.
	<-done
}
