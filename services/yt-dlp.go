package services

import (
	"bufio"
	"encoding/json"
	"log"
	"os/exec"
)

var command exec.Cmd
var argMap map[string]string

type Options struct {
	URL     string   `json:"original_url"`
	Formats []Format `json:"formats"`
}

type Format struct {
	ID         string `json:"format_id"`
	Ext        string `json:"ext"`
	VideoExt   string `json:"video_ext"`
	Resolution string `json:"resolution"`
}

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

func GetOptions(url string) (Options, error) {
	var options Options

	optionsCommand := exec.Command("yt-dlp", url, "-j")
	output, err := optionsCommand.Output()
	if err != nil {
		log.Println("Couldn't start:", err)
		return options, err
	}

	err = json.Unmarshal(output, &options)
	if err != nil {
		log.Println("Unmarshal err: ", err)
		return options, nil
	}

	return options, nil
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
