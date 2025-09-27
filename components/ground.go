package c

import rl "github.com/gen2brain/raylib-go/raylib"

// BECNHOD
// ZAMEEN HAI
// NEEND AA RAHI HAI MUJHE
// 30+ GHANTE SE JAAG RHA HOO
type Ground struct {
	// FUCKING IMAGE
	FuckingTILE rl.Texture2D
}

func NewGround() Ground {
	return Ground{
		FuckingTILE: rl.LoadTexture("assets/BHADWETILES/Tileset.png"),
	}

}

// WORLD COORDINATE ME KARNA WARNA GAAND ME DANDA
func (g *Ground) DRAW_KARDO(X, Y int) {
	// MUJHE NAHI PATA ME KYA KAR RHA HOO
}

// DrawTile draws a tile from a tileset at a given position.
// tileX, tileY = tile index in the tileset grid (not pixels!)
// x, y = screen position to draw the tile.
func DrawTile(tileset rl.Texture2D, tileX, tileY int32, x, y int32) {
	// HAAN CHATGPT KO CHODA HAI
	const tileWidth = 64
	const tileHeight = 32

	src := rl.NewRectangle(
		float32(tileX*tileWidth),
		float32(tileY*tileHeight),
		float32(tileWidth),
		float32(tileHeight),
	)

	dest := rl.NewRectangle(float32(x), float32(y), float32(tileWidth), float32(tileHeight))
	origin := rl.NewVector2(0, 0)

	rl.DrawTexturePro(tileset, src, dest, origin, 0, rl.White)
}
