package services

import (
	"bufio"
	"log"
	"os/exec"
)

var command exec.Cmd
var defaultArgs []string

func init() {
	defaultArgs = []string{}
}

func CheckYtDlp() (string, error) {
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		return "", err
	}
	return path, nil
}

func AddArgument(arg string) {
	log.Println("adding argument", arg)
	command.Args = append(command.Args, arg)
}

func RemoveArgument() {
}

func Download() error {
	ytdlpPath, err := CheckYtDlp()
	if err != nil {
		log.Println(err)
		return err
	}

	command.Path = ytdlpPath

	log.Println("starting download with", command.Path, command.Args)
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Println("Couldn't connect to stdout:", err)
		return err
	}

	err = command.Run()
	if err != nil {
		log.Println("Couldn't start:", err)
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	// Check for any scanner errors.
	if err := scanner.Err(); err != nil {
		log.Println("Error reading stdout:", err)
	}
	return nil
}
