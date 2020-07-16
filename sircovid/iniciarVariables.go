package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func iniciarVariables() {
	count = 0

	//nube
	nube1.nubeX = float64(rand.Intn(screenWidth) + 300)
	nube1.nubeY = float64(rand.Intn(1500))

	//defino el juego
	Game1.nube = nube1
	//Game1.siguienteNivel = player1

	//paso de nivel
	ModeGame = 0
	ModeTitle = 0
	ModeGameOver = 0
	ElectNumPlayers = 0
	ElectPlayer = 0

}
func pasarNivel() {
	count = 0

	enemigo1.FrameOX = 0
	enemigo1.FrameOY = 48
	enemigo1.FrameNum = 1
	enemigo1.X = float64(200)
	enemigo1.Y = float64(200)
	enemigo1.FrameWidth = 32
	enemigo1.FrameHeight = 48
	enemigo1.num = rand.Intn(5)
	enemigo1.cambio = rand.Intn(100) + 20

	enemigo2.FrameOX = 0
	enemigo2.FrameOY = 48
	enemigo2.FrameNum = 1
	enemigo2.X = float64(screenWidth - 50)
	enemigo2.Y = float64(290)
	enemigo2.FrameWidth = 32
	enemigo2.FrameHeight = 48
	enemigo2.num = rand.Intn(5)
	enemigo2.cambio = rand.Intn(50) + 50

	//reinciar enemigos
	count1 = 0

	// viejo
	viejo.FrameOX = 0
	viejo.FrameOY = 96
	viejo.FrameNum = 1
	viejo.X = float64(25)
	viejo.Y = float64(375)
	viejo.FrameWidth = 32
	viejo.FrameHeight = 48

	//palyer 2
	chica.FrameOX = 0
	chica.FrameOY = 96
	chica.FrameNum = 1
	chica.X = float64(25)
	chica.Y = float64(415)
	chica.FrameWidth = 32
	chica.FrameHeight = 48

	//nube
	nube1.nubeX = float64(rand.Intn(screenWidth) + 300)
	nube1.nubeY = float64(rand.Intn(1500))

	//defino el juego
	Game1.nube = nube1
	//Game1.siguienteNivel = player1

	//paso de nivel
	ModeTitle = 2
	ModeGame = 1

	//para movimiento de player
	player1.a, player1.b, player1.c, player1.d = 0, 0, 0, 0
	player2.a, player2.b, player2.c, player2.d = 0, 0, 0, 0
}
