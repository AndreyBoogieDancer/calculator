package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {

	calcApp := app.New()
	window := calcApp.NewWindow("Calculator")
	window.Resize(fyne.NewSize(300, 200))

	var input string
	display := widget.NewLabel("0")

	appendInput := func(char string) {
		if char == "C" {
			input = ""
		} else if char == "=" {
			result, err := eval(input)
			if err != nil {
				display.SetText("Error")
				return
			}
			input = fmt.Sprintf("%g", result)

		} else {
			input += char
		}

		display.SetText(input)
	}

	buttons := []string{
		"7", "8", "9", "+",
		"4", "5", "6", "-",
		"1", "2", "3", "*",
		"C", "0", "=", "/",
	}

	grid := container.NewGridWithColumns(4)
	for _, button := range buttons {
		b := button
		grid.Add(widget.NewButton(b, func() {
			appendInput(b)
		}))
	}

	contant := container.NewVBox(
		display,
		grid,
	)
	window.SetContent(contant)

	window.ShowAndRun()

}

func eval(expr string) (any, error) {

	expr_obj, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return 0, fmt.Errorf("invalid expression")
	}

	result, err := expr_obj.Evaluate(nil)
	if err == nil {
		return result, nil
	} else {
		return 0, fmt.Errorf("invalid expression")
	}
}
