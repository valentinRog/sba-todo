package main

import (
	store "github.com/valentinRog/sba-todo/store/generate"
	ui "github.com/valentinRog/sba-todo/ui/generate"
)

func main() {
	ui.Generate()
	store.Generate()
}
