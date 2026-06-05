package checkout

import (
	"fmt"
	"time"
)

func RunDemo() {
	procesarDatos()
	consultarDB()
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("⏱ %s tardó: %v\n", name, time.Since(start))
	}
}

func procesarDatos() {
	defer timer("procesarDatos")()

	// Simulando trabajo pesado
	time.Sleep(1200 * time.Millisecond)
	fmt.Println("✅ Procesamiento terminado")
}

func consultarDB() {
	defer timer("consultarDB")()

	time.Sleep(300 * time.Millisecond)
	fmt.Println("✅ Consulta DB terminada")
}
