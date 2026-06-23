package main

import (
    "log"
    "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"os"
)

func main() {
    
    logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)
    srv := server.CreateServer(logger)
    if err := srv.Server.ListenAndServe(); err != nil {
        logger.Fatal(err)
    }
}
