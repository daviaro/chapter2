package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}
func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		log.Info("ingresando a Random Float")
		var valor = rand.Float64()
		fmt.Fprintln(w, valor)
		log.Info("randon Float generado ", valor)
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		log.Info("ingresando a Random Int")
		var valor = rand.Intn(100)
		fmt.Fprintln(w, valor)
		log.	( log.Fields{
			"valor": valor,
		}).Info("randon Int generado ")
	})

	http.ListenAndServe(":8000", newMux)
}
