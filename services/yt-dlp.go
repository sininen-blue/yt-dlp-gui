package services

import (
	"bufio"
	"log"
	"os/exec"
)

var command exec.Cmd
var defaultArgs []string

func init() {
	defaultArgs = []string{
		"--progress-template",
		"download:[%(progress)s]",
	}
	command.Args = append(command.Args, defaultArgs...)
}

func CheckYtDlp() error {
	_, err := exec.LookPath("yt-dlp")
	if err != nil {
		return err
	}
	return nil
}

func AddArgument(arg string) {
	log.Println("adding argument", arg)
	command.Args = append(command.Args, arg)
}

func RemoveArgument() {
}

func Download() error {
	err := CheckYtDlp()
	if err != nil {
		log.Println("yt-dlp could not be found:", err)
		return err
	}
	command = *exec.Command("yt-dlp", command.Args...)

	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Println("Couldn't connect to stdout:", err)
		return err
	}

	err = command.Start()
	if err != nil {
		log.Println("Couldn't start:", err)
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading stdout:", err)
	}
	if err := command.Wait(); err != nil {
		log.Println("Error waiting for command:", err)
	}
	return nil

}
