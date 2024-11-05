package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func deploy(dryRun bool) ([]string, error) {
	app := "helm"
	var command = "install"
	if dryRun {
		command = "template"
	}
	args := []string{
		command,
		"httpbin-staging",
		"manifest",
		"--values", "manifest/staging/values.yaml",
		"--namespace", "httpbin-helm",
		"--version", "0.6.2",
	}
	fmt.Println("Running command: ", app, args)
	os.Chdir("/Users/brianw/git/httpbin-helm")
	cmd := exec.Command(app, args...)
	// Turn into list of strings using buffer
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(stdout))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	} else {
		return lines, nil
	}
}

func main() {
	output, err := deploy(false)
	fmt.Println(output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
