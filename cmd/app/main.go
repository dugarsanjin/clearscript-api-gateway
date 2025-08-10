package main

import (
	"clearscript-api-gateway/internal/config"
)

func main() {

	cfg := config.MustLoad()

	// todo init logger (slog)
	// todo init router (chi, "chi render")
	// todo run server
}
