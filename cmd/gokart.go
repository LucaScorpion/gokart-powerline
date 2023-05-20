package main

import (
	"fmt"
	"gokart-prompt/internal"
	"gokart-prompt/internal/ansi"
	"gokart-prompt/internal/git"
	"gokart-prompt/internal/versions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected one argument: ps1 or ps2")
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "ps1" {
		ps1()
	} else if arg == "ps2" {
		ps2()
	} else {
		fmt.Println("Argument should be ps1 or ps2")
		os.Exit(1)
	}
}

func ps1() {
	fmt.Print("\n")
	fmt.Print(ansi.Bold())

	// Path and project info
	fmt.Print(internal.Path())
	fmt.Print(git.Git())

	// Languages
	fmt.Print(versions.Go.Version())
	fmt.Print(versions.Java.Version())
	fmt.Print(versions.Node.Version())
	fmt.Print(versions.Php.Version())
	fmt.Print(versions.Ruby.Version())
	fmt.Print(versions.Rust.Version())

	// Tools
	fmt.Print(versions.Docker.Version())

	fmt.Print(ansi.Reset())
}

func ps2() {
	fmt.Print(internal.ExitCode())
}
