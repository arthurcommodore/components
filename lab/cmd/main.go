package main

import (
	"fmt"
	"sync"
	"time"
)

type QuadroDeAvisos struct {
	mu    sync.RWMutex
	aviso string
}

func main() {
	quadro := &QuadroDeAvisos{
		aviso: "Bem-vindos ao projeto!",
	}

	var wg sync.WaitGroup

	// --- 1. SIMULANDO MÚLTIPLAS LEITURAS (RLock / RUnlock) ---
	// Vários funcionários podem LER o quadro ao MESMO TEMPO sem atrapalhar ninguém.
	fmt.Println("=== INÍCIO DAS LEITURAS SIMULTÂNEAS ===")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// RLock = Read Lock (Bloqueio de Leitura)
			// Vários RLock podem acontecer em paralelo. Só bloqueia se alguém estiver ESCREVENDO.
			quadro.mu.RLock()
			defer quadro.mu.RUnlock()

			fmt.Printf("[Funcionário %d] Lendo o quadro: '%s'\n", id, quadro.aviso)
			time.Sleep(500 * time.Millisecond) // Simula o tempo lendo
		}(i)
	}

	wg.Wait() // Espera todo mundo terminar de ler

	// --- 2. SIMULANDO ESCRITA EXCLUSIVA (Lock / Unlock) ---
	// Quando o chefe vai ESCREVER, ele tranca TUDO. Ninguém mais pode ler ou escrever até ele terminar.
	fmt.Println("\n=== O CHEFE VAI ATUALIZAR O QUADRO ===")
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Lock = Write Lock (Bloqueio Total / Exclusivo)
		// Ninguém mais entra aqui nem para ler (`RLock`) nem para escrever (`Lock`).
		quadro.mu.Lock()
		defer quadro.mu.Unlock()

		fmt.Println("[Chefe] Escrevendo no quadro...")
		quadro.aviso = "Aviso importante: Reunião às 14h!"
		time.Sleep(1 * time.Second)
		fmt.Println("[Chefe] Terminou de escrever.")
	}()

	wg.Wait()

	// --- 3. LENDO O NOVO AVISO ---
	quadro.mu.RLock()
	fmt.Printf("\n[Novo Estado] Quadro diz: '%s'\n", quadro.aviso)
	quadro.mu.RUnlock()
}
