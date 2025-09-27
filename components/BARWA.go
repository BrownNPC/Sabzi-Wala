package c

import rl "github.com/gen2brain/raylib-go/raylib"

// MUJHE MAAR DO
type GANDU struct {
	ANIMATIONS []rl.Texture2D // MARDO ISKO BHAAAII
	HITBOX     rl.RectangleInt32
	POSITION   rl.Vector2
}

// NAYA GANDU DEDO BHAI
func NewGandu() GANDU {
	return GANDU{
		ANIMATIONS: []rl.Texture2D{
			rl.LoadTexture("assets/MADARCHODEKIANIMATIONS/FRAME1.png"),
			rl.LoadTexture("assets/MADARCHODEKIANIMATIONS/FRAME2.png"),
		},
	}
}

// KARDO BENCHOD
func (g *GANDU) HITBOXDRAWKARDO() {
	rl.DrawRectangle(
		int32(g.POSITION.X),
		int32(g.POSITION.Y),
		g.HITBOX.Width,
		g.HITBOX.Height,
		rl.Red,
	)
}
