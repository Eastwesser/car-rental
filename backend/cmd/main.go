package main

import (
	"log"

	"github.com/eastwesser/car-rental/backend/run"
)

func main() {
    if err := run.Run(); err != nil {
        log.Fatal(err)
    }
}
