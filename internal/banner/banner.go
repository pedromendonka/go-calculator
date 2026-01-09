// Package banner Description: This file contains a function to print an ASCII art banner for the application.
// It displays the application name and version in a stylized format when called.
package banner

import "fmt"

func PrintBanner() {
	var (
		appName    = "Go Calculator"
		appVersion = "1.0.0"
	)

	banner := `
╔═══════════════════════════════════╗
║                                   ║
║  ███████  █████  ██     ████████  ║
║  ██      ██╔══██ ██     ██]       ║
║  ██      ██║  ██ ██     ██]       ║
║  ██      ██║████ ██     ██]       ║
║  ███████ ██║  ██ ██████ ████████  ║
║                                   ║
╚═══════════════════════════════════╝`
	fmt.Println(banner)
	fmt.Printf("\n  %s v%s\n\n", appName, appVersion)
}
