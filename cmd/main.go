package main

import (
	"github.com/TulioGuaraldoB/pizoo/internal/config/env"
	"github.com/TulioGuaraldoB/pizoo/internal/gateway/postgres"
)

func main() {
	// Environment Variable
	env.GetEnvironmentVariable()

	// Driven
	postgres.OpenConnection()


}
