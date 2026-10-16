package main

import (
	"fmt"
	"time"
)

type Event struct {
	Mensagem string
}

func main() {
	// Criamos um channel com buffer para 3 mensagens
	// Se gerarem 3 mensagens rápidas, elas entram sem travar a goroutine que envia.
	canal := make(chan Event, 3)

	// Goroutine simulando o sistema gerando eventos sem parar
	go func() {
		for i := 1; i <= 5; i++ {
			canal <- Event{Mensagem: fmt.Sprintf("Evento #%d", i)}
			fmt.Printf(">> Enviado para o buffer: Evento #%d\n", i)
		}
	}()

	// O cliente lendo as mensagens em um loop infinito/continuo
	// Dando um tempo para simular o recebimento gradual
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n-- Cliente consumindo do channel --")

	// Usamos um loop para esvaziar o channel conforme chegam
	for evento := range canal {
		fmt.Printf("<< Cliente recebeu: %s\n", evento.Mensagem)

		// Condição para parar o loop neste exemplo simples
		if evento.Mensagem == "Evento #5" {
			break
		}
	}
}

// Depois pessoas podem ver a parede RLock mas quando todas terminarem de ver o pintou vai dar um Lock
// para alterar a cor
