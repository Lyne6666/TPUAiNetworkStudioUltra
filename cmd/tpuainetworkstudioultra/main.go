// cmd/tpuainetworkstudioultra/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"tpuainetworkstudioultra/internal/tpuainetworkstudioultra"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := tpuainetworkstudioultra.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
