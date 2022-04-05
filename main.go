package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

var (
	apps = map[string]string{
		"Next.js":            "npx create-next-app@latest",
		"Next.js Typescript": "npx create-next-app@latest --typescript",
		"Angular":            "npx ng new",
		"Vite":               "npm create vite@latest",
		"Solidjs":            "npx degit solidjs/templates/js",
		"Solidjs Typescript": "npx degit solidjs/templates/ts",
		"Astro":              "npm init astro",
	}
)

func main() {
	choices := make([]string, 0, len(apps))
	for key := range apps {
		choices = append(choices, key)
	}
	questions := []*survey.Question{
		{
			Name: "Name",
			Prompt: &survey.Input{
				Message: "Application name:",
			},
			Validate: survey.Required,
		},
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
		Name string
		App  string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	run(answers.Name, answers.App)
}

func run(name string, app string) error {
	commandInput := fmt.Sprintf("%s %s", apps[app], name)
	fmt.Printf("%s\n", commandInput)
	cmd := exec.Command("bash", "-c", commandInput)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
