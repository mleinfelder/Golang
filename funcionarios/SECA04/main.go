package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	agora := time.Now()
	url := os.Args[1]
	get, err := http.Get(url)
	if err != nil {
		fmt.Println("Ocorreu um erro ao executar o get url")
		panic(err)
	}
	decorrido := time.Since(agora).Seconds()
	// HTTP 200,201
	status := get.StatusCode
	fmt.Printf("Status:  [ %d] Tempod de carga [%f] \n", status, decorrido)
}
