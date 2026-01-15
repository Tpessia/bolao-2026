package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetPort() int {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Invalid server port: '%d'", port)
	}
	return port
}
