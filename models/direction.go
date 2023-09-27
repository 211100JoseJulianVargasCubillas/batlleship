package spacegame

// Direction es un tipo enumerado que representa las direcciones posibles en el juego.
type Direction int

const (
	// Idle representa la dirección en reposo o sin movimiento.
	Idle Direction = iota

	// LeftDirection representa la dirección hacia la izquierda.
	LeftDirection

	// RightDirection representa la dirección hacia la derecha.
	RightDirection
)
