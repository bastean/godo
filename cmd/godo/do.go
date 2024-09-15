package main

import (
	"github.com/bastean/godo/internal/app/cli"
	"github.com/bastean/godo/internal/pkg/service"
)

func main() {
	service.Up()
	cli.Up()
}
