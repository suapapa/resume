package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := loadData()
	if err != nil {
		return fmt.Errorf("failed to load data: %w", err)
	}

	if err := genResume(config, "_deploy/index.html"); err != nil {
		return fmt.Errorf("failed to generate resume: %w", err)
	}

	return nil
}
