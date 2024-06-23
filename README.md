## Project_Kat

## Project Progress

Below, see a summary of current tasks that either need doing or have been done.

### To Do
- Start building level one get gauntlet section.
- Improve the hitbox collisions.
	- I should also see if I can refactor the level collider so that it is on the level object rather than being an entirely seperate entity.
	- This might mean that the level needs to become an entity.

### Done
- Time travel world mechanic.
- Camera movement.
- Fix bug with falling animation when facing left.
- Do player walking animation.
- Update Player Sprite to New.
- Character Rendering
- Character Physics And Collisions With TileMap Data
- Get Tile Map Collision Data
- Pull Out Level Data Loader From `levelOne.go` And Stick It In `engine_levels.LevelDataLoader` interface.
- Pulled out the majority of the engine code into packages prefixed with `engine_`
- Tile Map Rendering
- Background Image Rendering - Image Render
- Background Image Rendering - Parallax Effect
- Camera Stub

## Build and Run
To debug the project in vscode, you can simply press F5 (but if you don't know this already, you probably should start elsewhere.).

For those who hate debuggers, or who use a real man's editor like NeoVim, you can boot it up the good old fashioned way by running:

```bash
go run cmd/Project_Kat_3/main.go
```

### Dependencies

When running on Ubuntu based distros, you will need to install the following dependencies:

```bash
apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev
```
