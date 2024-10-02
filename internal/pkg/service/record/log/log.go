package log

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"

	"github.com/bastean/godo/pkg/infrastructure/record/log"
)

var (
	Log     = log.New()
	Debug   = Log.Debug
	Error   = Log.Error
	Fatal   = Log.Fatal
	Info    = Log.Info
	Success = Log.Success
)

var (
	White = color.New(color.FgWhite, color.Bold).Sprint
	Cyan  = color.New(color.FgCyan, color.Bold).Sprint
)

func Logo() {
	figureGo := figure.NewFigure("GO", "speed", true).Slicify()
	figureDo := figure.NewFigure("DO", "speed", true).Slicify()

	width := 0
	fixedWidth := 0

	for _, line := range figureGo {
		width = len(line)

		if width > fixedWidth {
			fixedWidth = width
		}
	}

	for i, line := range figureGo {
		width = len(line)

		if width < fixedWidth {
			line += strings.Repeat(" ", (fixedWidth - width))
		}

		fmt.Println(Cyan(line), White(figureDo[i]))
	}

	fmt.Println()
}

func Service(service string) string {
	return fmt.Sprintf("Service:%s", service)
}

func Module(module string) string {
	return fmt.Sprintf("Module:%s", module)
}

func Starting(service string) {
	Info(fmt.Sprintf("Starting %s...", service))
}

func Started(service string) {
	Success(fmt.Sprintf("%s started", service))
}

func CannotBeStarted(service string) {
	Error(fmt.Sprintf("%s cannot be started", service))
}

func Stopping(service string) {
	Info(fmt.Sprintf("Stopping %s...", service))
}

func Stopped(service string) {
	Success(fmt.Sprintf("%s stopped", service))
}

func CannotBeStopped(service string) {
	Error(fmt.Sprintf("%s cannot be stopped", service))
}
