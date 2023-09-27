package spacegame

// Action es un tipo enumerado que representa las acciones posibles en el juego.
type Action int

const (
	// NoneAction representa ninguna acción.
	NoneAction Action = iota

	// ShootAction representa la acción de disparar.
	ShootAction
)
