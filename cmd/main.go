// Main entry point for the application
package main

import (
	"go-calculator/internal/banner"
	"go-calculator/internal/calculator"
)

// main is the entry point of the application
func main() {
	banner.PrintBanner()
	calculator.Run()
}
