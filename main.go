package main

import (
	"flag"
	"os"
	"log"
	"path"
	"bufio"
	"os/exec"
)

func main() {
	flag.Parse()

	// If there are no arguments show usage message
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	// Capture command to run
	cmdToRun := flag.Arg(0)
	var cmdOptionsToPass []string
	if flag.NArg() > 1 {
		cmdOptionsToPass = flag.Args()[1:]
	}

	// Open .env file
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldn't get current working directory: %v", err)
	}
	envFilePath := path.Join(cwd, ".env")
	file, err := os.Open(envFilePath)
	if err != nil {
		log.Fatalf("error opening .env file: %v", err)
	}
	defer file.Close()

	// Extract environment variables from .env file.
	scanner := bufio.NewScanner(file)
	var envs []string
	for scanner.Scan() {
		envs = append(envs, scanner.Text())
	}

	cmd := exec.Command(cmdToRun, cmdOptionsToPass...)
	cmd.Env = envs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Fatalf("error executing command: %v", err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("error executing command: %v", err)
	}
}
