package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

func initTextos() {

	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	smallArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func dibujarTextos(screen *ebiten.Image) {
	// dibujar vidas y monedas
	lifesP1 := fmt.Sprintf("Life:%02d  Coins:%d", player1.vidas, player1.Coins)
	text.Draw(screen, lifesP1, smallArcadeFont, 20, 515, color.Black)
	if Game1.numPlayers == 2 {
		lifesP2 := fmt.Sprintf("Life:%02d  Coins:%d", player2.vidas, player2.Coins)
		text.Draw(screen, lifesP2, smallArcadeFont, 785, 515, color.Black)
	}

	//dibujar inmunidad
	if player1.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, fontSize, 100, color.White)
	}
	if player2.Inmune {
		Inm := fmt.Sprintf("Inmune for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 840, 100, color.White)
	}
	//dibujar Fast
	if player1.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player1.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, fontSize, 130, color.White)
	}
	if player2.Fast {
		Inm := fmt.Sprintf("Fast for:%02d", player2.CountPoder/60)
		text.Draw(screen, Inm, smallArcadeFont, 840, 130, color.White)
	}
	switch {
	case ModeTitle:
		// intro draw
		intro1.drawIntro(screen, screenWidth, screenHeight)

	case ModeTitleLevel:
		nivel := fmt.Sprintf("    LEVEL %d", Level)
		text.Draw(screen, nivel, arcadeFont, 300, 200, color.Black)
		nivel1 := fmt.Sprintf("PRESS SPACE KEY")
		text.Draw(screen, nivel1, smallArcadeFont, 400, 290, color.Black)

		if Level == 1 {
			nivel := fmt.Sprintf("MISSION: GET MEDS")
			text.Draw(screen, nivel, smallArcadeFont, 390, 250, color.Black)
		}
		if Level == 2 {
			nivel := fmt.Sprintf("MISSION: GET FOOD")
			text.Draw(screen, nivel, smallArcadeFont, 390, 250, color.Black)
		}
		if Level == 3 {
			nivel := fmt.Sprintf("MISSION GET CLOTHES")
			text.Draw(screen, nivel, smallArcadeFont, 380, 250, color.Black)
		}
		if Level == 4 {
			nivel := fmt.Sprintf("MISSION: GET TOLIET PAPER")
			text.Draw(screen, nivel, smallArcadeFont, 350, 250, color.Black)
		}
		if Level == 5 {
			nivel := fmt.Sprintf("MISSION: GET MEDS AND FOOD")
			text.Draw(screen, nivel, smallArcadeFont, 330, 250, color.Black)
		}
		if Level == 6 {
			nivel := fmt.Sprintf("MISSION: GET TOLET PAPER AND MEDS")
			text.Draw(screen, nivel, smallArcadeFont, 280, 250, color.Black)
		}
		if Level == 7 {
			nivel := fmt.Sprintf("MISSION: GET GLOTHES AND FOOD")
			text.Draw(screen, nivel, smallArcadeFont, 300, 250, color.Black)
		}
		if Level == 8 {
			nivel := fmt.Sprintf("MISSION: GET MEDS AND CLOTHES")
			text.Draw(screen, nivel, smallArcadeFont, 320, 250, color.Black)
		}
		if Level == 9 {
			nivel := fmt.Sprintf("MISSION GET TOILET PAPER AND CLOTHES")
			text.Draw(screen, nivel, smallArcadeFont, 240, 250, color.Black)
		}
		if Level == 10 {
			nivel := fmt.Sprintf("MISSION: GET BREAD AND TOILET PAPER")
			text.Draw(screen, nivel, smallArcadeFont, 260, 250, color.Black)
		}
		if Level > 10 {
			nivel := fmt.Sprintf("MISSION: GET VACCINE")
			text.Draw(screen, nivel, smallArcadeFont, 370, 250, color.Black)
		}

	case ModeGameOver:
		lost := fmt.Sprintf("  GAME OVER!\n\n  TRAY AGAIN?\n\nPRESS SPACE KEY")
		noMoney := fmt.Sprintf("Mission failed!\n\n  TRAY AGAIN?\n\nPRESS SPACE KEY")

		if player1.Coins < 2 && player2.Coins < 2 && monedas.X == 1500 {
			text.Draw(screen, noMoney, arcadeFont, 310, 200, color.Black)
		} else {
			text.Draw(screen, lost, arcadeFont, 310, 200, color.Black)
		}
	case ModeWin == true:
		win := fmt.Sprintf("YOU WIN")
		text.Draw(screen, win, arcadeFont, 400, 300, color.White)
	}
	switch {
	case ModePause && count1 < 40:
		jugadores := fmt.Sprintf("PAUSE")
		text.Draw(screen, jugadores, arcadeFont, 440, 220, color.Black)

		//elegir numero de jugadores
	case ElectNumPlayers == 0 && Game1.numPlayers == 1:
		jugadores := fmt.Sprintf(">1 PLAYER\n 2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 350, 300, color.Black)
	case ElectNumPlayers == 0 && Game1.numPlayers == 2:
		jugadores := fmt.Sprintf(" 1 PLAYER\n>2 PLAYERS")
		text.Draw(screen, jugadores, arcadeFont, 350, 300, color.Black)

		//elegir jugador
	case ElectPlayer == 0:
		jugadores := fmt.Sprintf("CHOOSE PLAYER")
		text.Draw(screen, jugadores, arcadeFont, 320, 300, color.Black)

	case ElectNumPlayers == 1:
		for i, l := range texts {
			x := (screenWidth - len(l)*fontSize) / 2
			text.Draw(screen, l, arcadeFont, x, (i+5)*fontSize, color.White)
		}
	}
	//dibujar Comandos
	if Commands {
		plasmaVida := fmt.Sprintf("Player 1 movements:\n\n Up: Arrow Up\n Dow: Arrow Down\n Rigth: Arrow Rigth\n Left: Arrow Left")
		text.Draw(screen, plasmaVida, smallArcadeFont, 325, 150, color.White)
		alcholInmune := fmt.Sprintf("\nPlayer 2 movements:\n\n Up: W\n Dow: S\n Rigth: D\n Left: A")
		text.Draw(screen, alcholInmune, smallArcadeFont, 325, 250, color.White)
		barbijoInmune := fmt.Sprintf("\nPause: Space Bar\n\nMute: X\n\nVolume Up: Ctrl+Arrow Up\nVolume Down: Ctrl+Arrow Down")
		text.Draw(screen, barbijoInmune, smallArcadeFont, 325, 360, color.White)

	}
}
