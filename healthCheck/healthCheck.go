package main

import (
	"io"
	"net/http"
	"os"
	"time"

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

// HealthCheck API returns date time to client
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	log.Info("Entrando al metodo healthCheck y se obtiene la hora actual ", currentTime.String())
	io.WriteString(w, currentTime.String())
}

func main() {
	log.Warn("Se arranca el servicio en el puerto 8080")
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
