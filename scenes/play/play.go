package play

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	screen   render.Screen
	cam      rl.Camera2D
	ZAMEENBC c.Ground
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	scene.screen = render.NewScreen(ctx.VirtualResolution)
	scene.cam = rl.NewCamera2D(
		ctx.VirtualResolution.Scale(0.5).R(),
		c.V2Z.R(),
		0,
		1.0,
	)
}

// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	scene.screen.BeginDrawing()
	rl.BeginMode2D(scene.cam)
	rl.DrawEllipse(0, 0, 30, 60, rl.Red)
	rl.DrawRectangle(100, 50, 200, 100, rl.Blue) // world position
	rl.EndMode2D()
	scene.screen.EndDrawing()
	return false // if true is returned, Unload is called
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "someOtherSceneId" // the engine will switch to the scene that is registered with this id
}
