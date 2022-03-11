package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

const (
	NextJS  = "npx create-next-app@latest"
	NextTS  = "npx create-next-app@latest --typescript"
	Angular = "npx ng new"
	Vite    = "npm create vite@latest"
	SolidJS = "npx degit solidjs/templates/js"
	SolidTS = "npx degit solidjs/templates/ts"
)

var (
	apps = map[string]string{
		"Next.js":            NextJS,
		"Next.js Typescript": NextTS,
		"Angular":            Angular,
		"Vite":               Vite,
		"Solidjs":            SolidJS,
		"Solidjs Typescript": SolidTS,
	}
)

func main() {
	choices := make([]string, 0, len(apps))
	for key := range apps {
		choices = append(choices, key)
	}
	questions := []*survey.Question{
		{
			Name: "App",
			Prompt: &survey.Select{
				Message: "Choose an app:",
				Options: choices,
			},
			Validate: survey.Required,
		},
	}
	answers := struct {
		App string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	run(answers.App)
}

func run(input string) error {
	fmt.Println(apps[input])
	cmd := exec.Command("bash", "-c", apps[input])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
