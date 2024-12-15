package services

import (
	"bufio"
	"log"
	"os/exec"
)

var command exec.Cmd
var argMap map[string]string

func init() {
	argMap = make(map[string]string)
	argMap["--progress-template"] = "download:[%(progress)s]"
}

func CheckYtDlp() error {
	_, err := exec.LookPath("yt-dlp")
	if err != nil {
		return err
	}
	return nil
}

func GetOptions(url string) error {
	optionsCommand := exec.Command("yt-dlp", url, "-j")

	stdout, err := optionsCommand.StdoutPipe()
	if err != nil {
		log.Println("Couldn't connect to stdout:", err)
		return err
	}

	log.Println(optionsCommand.Args)
	err = optionsCommand.Start()
	if err != nil {
		log.Println("Couldn't start:", err)
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// do parsing here
		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading stdout:", err)
	}

	if err := optionsCommand.Wait(); err != nil {
		log.Println("Error waiting for optionsCommand:", err)
	}
	return nil
}

func SetArgument(argName string, arg string) {
	argMap[argName] = arg
}

func Download() error {
	err := CheckYtDlp()
	if err != nil {
		log.Println("yt-dlp could not be found:", err)
		return err
	}
	command = *exec.Command("yt-dlp")

	for argName, argValue := range argMap {
		command.Args = append(command.Args, argName)
		command.Args = append(command.Args, argValue)
	}

	command.Args = append(command.Args, argMap["url"])

	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Println("Couldn't connect to stdout:", err)
		return err
	}

	log.Println(command.Args)
	err = command.Start()
	if err != nil {
		log.Println("Couldn't start:", err)
		return err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// do parsing here
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
