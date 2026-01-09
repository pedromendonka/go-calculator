// Main entry point for the application
package main

import (
	"go-calculator/internal/banner"
	"go-calculator/internal/calculator"
)

func main() {
	banner.PrintBanner()
	calculator.Run()
}
