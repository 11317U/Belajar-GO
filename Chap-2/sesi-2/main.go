package main

import (
	router "sesi-2/Router"
)

const PORT = ":8080"

func main() {

	router.StartServer().Run(PORT)
}
