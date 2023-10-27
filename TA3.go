package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	TamanioTablero = 58
	numJugadores   = 2
	fichasMax      = 4
	PosWin         = 57
)

var (
	tablero [TamanioTablero]int
	wg      sync.WaitGroup
	mu      sync.Mutex
)

type Jugador struct {
	ID             int
	fichas         [fichasMax]int
	ganador        bool
	siguienteFicha int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < TamanioTablero; i++ {
		tablero[i] = -1
	}

	jugadores := make([]Jugador, numJugadores)

	wg.Add(numJugadores)

	for i := 0; i < numJugadores; i++ {
		jugadores[i] = Jugador{ID: i}
		go jugarLudo(&jugadores[i])
	}

	wg.Wait()
}

func jugarLudo(jugador *Jugador) {
	defer wg.Done()

	for !jugador.ganador {
		dados := tirarDados()

		if jugador.fichas[jugador.siguienteFicha] >= PosWin {
			jugador.siguienteFicha++
			if jugador.siguienteFicha == fichasMax {
				jugador.ganador = true
				fmt.Printf("JUGADOR %d HA GANADO!!!!!!!\n", jugador.ID+1)
				os.Exit(0)
			}
			continue
		}

		fichaPos := jugador.fichas[jugador.siguienteFicha]
		nuevaPos := fichaPos + dados

		if nuevaPos > PosWin {
			nuevaPos = PosWin
		}

		jugador.fichas[jugador.siguienteFicha] = nuevaPos
		fmt.Printf("TURNO DEL JUGADOR %d\n", jugador.ID+1)
		fmt.Printf("Jugador %d, ha sacado un %d.\n", jugador.ID+1, dados)
		fmt.Printf("Jugador %d, peón número %d, posición %d.\n\n", jugador.ID+1, jugador.siguienteFicha+1, nuevaPos)

		if casillaEspecial(nuevaPos) {
			fmt.Printf("Jugador %d pierde un turno, cayó en una casilla múltiplo de 10\n\n", jugador.ID+1)
		}

		if nuevaPos >= PosWin {
			jugador.siguienteFicha++
			if jugador.siguienteFicha == fichasMax {
				jugador.ganador = true
				fmt.Printf("JUGADOR %d HA GANADO!!!!!!!!\n", jugador.ID+1)
				os.Exit(0)
			}
		}

		time.Sleep(time.Millisecond)
	}
}

func tirarDados() int {
	return rand.Intn(6) + 1
}

func casillaEspecial(fichaPos int) bool {
	return fichaPos%10 == 0
}
