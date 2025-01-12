package main

import (
	"fmt"
	"github.com/gsoares85/GoMosaic/server"
	"net/http"
)

func main() {
	config := server.LoadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := server.NewContext(w, r)
		ctx.JSON(http.StatusOK, map[string]string{"message": "Mosaic is working!"})
	})

	fmt.Println("Starting server")
	srv := server.New(":"+config.Port, mux)
	if err := srv.Start(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
