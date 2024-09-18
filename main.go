package main

import (
	"context"
	"os"
)

func main() {
	component := SearchBox()
	component.Render(context.Background(), os.Stdout)
}
