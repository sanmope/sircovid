package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

//Game es la estructura del juego
type Game struct {
	count int
	nube
	numPlayers     int
	electPlayer    int
	siguienteNivel (player)
}

// Game1 es el juego
var Game1 Game
var intro1 intro

var (
	Relato                    bool
	ModeGame                  bool
	ModeTitleLevel            bool
	Level                     int
	ModeWin                   bool
	ModeTitle                 bool
	ElectNumPlayers           int
	ElectPlayer               int
	ModeGameOver              bool
	count1                    int
	ModePause                 bool
	elecCompras               int
	farmacia, farmacia1       bool
	mart, mart1               bool
	bakery, bakery1           bool
	supermarket, supermarket1 bool

	// imágenes
	imgTiles *ebiten.Image

	// sonido
	// audioContext *audio.Context
	// deadSound    *audio.Player
	// deadSound2   *audio.Player
	// sonidoFondo  *audio.InfiniteLoop
	// fondo        *audio.Player
	// sonidoIntro  *audio.InfiniteLoop
	// sIntro       *audio.Player

	//para start y game over
	arcadeFont      font.Face
	smallArcadeFont font.Face
	texts           = []string{}

	err error
)

const (
	// game
	screenWidth  = 66 * 16
	screenHeight = 33 * 16

	// tiles
	tileSize = 16
	tileXNum = 48

	//para start y game Over
	fontSize      = 32
	smallFontSize = fontSize / 2
)

//introduccion es la eleccion de los players, etc
func introduccion() {
	// intro update
	intro1.updateIntro(screenWidth, screenHeight)

	switch {
	case ElectNumPlayers == 0:
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			Game1.numPlayers = 2
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			Game1.numPlayers = 1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) && (Game1.numPlayers == 1 || Game1.numPlayers == 2) {
			ElectNumPlayers = 1
		}
	case ElectPlayer == 0 && Game1.numPlayers == 1 || Game1.numPlayers == 2:
		if Game1.numPlayers == 1 || Game1.numPlayers == 2 {
			if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
				player1.humanos.img = humano1.img
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
				player1.humanos.img = humano2.img
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyW) {
				player2.humanos.img = humano1.img
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyS) {
				player2.humanos.img = humano2.img
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			ElectPlayer = 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && ElectPlayer == 1 {
		ModeTitle = false
		player1.X[0] = 15
		player1.Y[0] = -40
		player2.X[0] = 130
		player2.Y[0] = -40
		ModeTitleLevel = true
	}

}

func siguienteNivel(p player) player {
	if p.CompleteLevel && (p.X[0] >= home.X && p.X[0] <= home.X+40 && p.Y[0] == -40 || p.X[0] >= home1.X && p.X[0] <= home1.X+40 && p.Y[0] == -40 && Game1.numPlayers == 2) {
		player1.CompleteLevel = false
		player2.CompleteLevel = false
		farmacia, mart, supermarket, bakery = false, false, false, false
		pasarNivel()
		fondo.Pause()
		fondo.Rewind()
		sLevelUp.Play()
		sLevelUp.Rewind()

	}
	return p
}
func compar(p player) player {
	//compras

	if (inpututil.IsKeyJustPressed(ebiten.KeyDown) && p.señalador == 0) || (inpututil.IsKeyJustPressed(ebiten.KeyS) && p.señalador == 1) {
		elecCompras++
	}
	if (inpututil.IsKeyJustPressed(ebiten.KeyUp) && p.señalador == 0) || (inpututil.IsKeyJustPressed(ebiten.KeyW) && p.señalador == 1) {
		elecCompras--
	}
	if elecCompras > 2 {
		elecCompras = 2
	}
	if elecCompras < 0 {
		elecCompras = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if (farmacia || bakery) && elecCompras == 0 {
			p.Coins = p.Coins - 3
			p.vidas++
		}
		if mart && elecCompras == 1 {
			p.Coins = p.Coins - 5
			pasarNivel()
		}
		if mart && elecCompras == 0 {
			p.Coins = p.Coins - 2
			p.Inmune = true
			p.CountPoder = 600
		}
		if elecCompras == 1 && (farmacia || supermarket || bakery) {
			p.Coins = p.Coins - 2
			p.Fast = true
			p.CountPoder = 600
		}
		switch {
		case elecCompras == 2 && bakery:
			p.Coins = p.Coins - 2
			bakery1 = true
		case elecCompras == 2 && farmacia:
			p.Coins = p.Coins - 2
			farmacia1 = true
		case elecCompras == 2 && mart:
			p.Coins = p.Coins - 2
			mart1 = true
		case elecCompras == 2 && supermarket:
			p.Coins = p.Coins - 2
			supermarket1 = true
		}
		switch {
		case farmacia && elecCompras == 2 && Level == 1:
			p.CompleteLevel = true
		case bakery && elecCompras == 2 && Level == 2:
			p.CompleteLevel = true
		case mart && elecCompras == 2 && Level == 3:
			p.CompleteLevel = true
		case supermarket && elecCompras == 2 && Level == 4:
			p.CompleteLevel = true
		case farmacia1 && bakery1 && Level == 5:
			p.CompleteLevel = true
		case farmacia1 && supermarket1 && Level == 6:
			p.CompleteLevel = true
		case mart1 && bakery1 && Level == 7:
			p.CompleteLevel = true
		case farmacia1 && mart1 && Level == 8:
			p.CompleteLevel = true
		case supermarket1 && mart1 && Level == 9:
			p.CompleteLevel = true
		case supermarket1 && bakery1 && Level == 10:
			p.CompleteLevel = true
		}

		farmacia, mart, supermarket, bakery = false, false, false, false
		p.Compras = false
	}
	return p
}
func dibujarTextoCompras(p player, screen *ebiten.Image) {

	if p.Compras {
		if p.Coins < 2 {
			jugadores := fmt.Sprintf("YOU DONT HAVE MONEY\n  COME BACK SOON")
			text.Draw(screen, jugadores, arcadeFont, 230, 200, color.White)
		}
		//EN FARMACIA
		switch {
		case farmacia && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$3-PLASMA -GET LIFE-\n $2-ASPIRIN -GO FAST-\n $2-MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
		case farmacia && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-PLASMA -GET LIFE-\n>$2-ASPIRIN -GO FAST-\n $2-MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)
		case farmacia && elecCompras == 2 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-PLASMA -GET LIFE-\n $2-ASPIRIN -GO FAST-\n>$2-MEDICINE")
			text.Draw(screen, jugadores, arcadeFont, 300, 250, color.White)

			//EN BAKERY
		case bakery && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$3-CRIOLLOS-GET LIFE-\n $2-CAFE -GO FAST-\n $2-BREAD")
			text.Draw(screen, jugadores, arcadeFont, 250, 250, color.White)
		case bakery && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-CRIOLLOS-GET LIFE-\n>$2-CAFE -GO FAST-\n $2-BREAD")
			text.Draw(screen, jugadores, arcadeFont, 250, 250, color.White)
		case bakery && elecCompras == 2 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-CRIOLLOS-GET LIFE-\n $2-CAFE -GO FAST-\n>$2-BREAD")
			text.Draw(screen, jugadores, arcadeFont, 250, 250, color.White)

			//EN MART
		case mart && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$2-MOUTH COVER-GET INMUNE-\n $5-HAT\n $2-CLOTHES")
			text.Draw(screen, jugadores, arcadeFont, 150, 250, color.White)
		case mart && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $2-MOUTH COVER-GET INMUNE-\n>$5-HAT\n $2-CLOTHES")
			text.Draw(screen, jugadores, arcadeFont, 150, 250, color.White)
		case mart && elecCompras == 2 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $2-MOUTH COVER-GET INMUNE-\n $5-HAT\n>$2-CLOTHES")
			text.Draw(screen, jugadores, arcadeFont, 150, 250, color.White)

			//en SUPERMARKET
		case supermarket && elecCompras == 0 && p.Coins >= 2:
			jugadores := fmt.Sprintf(">$3-FOOD-GET LIFE-\n $2-ENERGIZING -GO FAST-\n $2-WATER")
			text.Draw(screen, jugadores, arcadeFont, 200, 250, color.White)
		case supermarket && elecCompras == 1 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-FOOD-GET LIFE-\n>$2-ENERGIZING -GO FAST-\n $2-WATER")
			text.Draw(screen, jugadores, arcadeFont, 200, 250, color.White)
		case supermarket && elecCompras == 2 && p.Coins >= 2:
			jugadores := fmt.Sprintf(" $3-FOOD-GET LIFE-\n $2-ENERGIZING -GO FAST-\n>$2-WATER")
			text.Draw(screen, jugadores, arcadeFont, 200, 250, color.White)
		}
	}
}
