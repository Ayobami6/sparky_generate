package main

import (
	"os"
	"os/exec"
)

func generateProject(folderName string, projectChoice string) {
	// Create folder
	os.Mkdir(folderName, 0755)
	os.Chdir(folderName)
	switch projectChoice {
	case "Go":
		command := exec.Command("git", "clone", "https://github.com/Ayobami6/go_starter", ".")
		_, err := command.CombinedOutput()
		if err != nil {
			panic(err)
		}
	case "Python Django":
		command := exec.Command("git", "clone", "https://github.com/Ayobami6/django_starter_template", ".")
		_, err := command.CombinedOutput()
		if err != nil {
			panic(err)
		}
	case "NodeJS Nestjs":
		command := exec.Command("git", "clone", "https://github.com/nestjs/typescript-starter.git", ".")
		_, err := command.CombinedOutput()
		if err != nil {
			panic(err)
		}
	case "Java SpringBoot":
		command := exec.Command("git", "clone", "https://github.com/Ayobami6/java_springboot_starter_template", ".")
		_, err := command.CombinedOutput()
		if err != nil {
			panic(err)
		}
	case "Python FastApi":
		command := exec.Command("git", "clone", "https://github.com/Ayobami6/fast_api_starter", ".")
		_, err := command.CombinedOutput()
		if err != nil {
			panic(err)
		}
	default:
		// clo
	}
	command := exec.Command("git", "remote", "remove", "origin")
	_, err := command.CombinedOutput()
	if err != nil {
		panic(err)
	}

}
