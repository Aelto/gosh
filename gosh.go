package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args

	inputDir := hasValidOption(args, "-in", "C:/")

	templates := []templateDeclaration{}
	templatesCalls := []templateCall{}
	filesContents := []string{}

	files := getFilesList(inputDir)

	for _, file := range files {
		log(file)
	}

	for _, file := range files {

		buf, err := ioutil.ReadFile(file)
		check(err)

		textContent := string(buf)
		filesContents = append(filesContents, textContent)

		for _, template := range searchTemplateDeclarationNext(textContent, file) {
			templates = append(templates, template)
		}
	}

	for _, fileContent := range filesContents {
		templatesCalls = append(templatesCalls, searchTemplateCallNext(fileContent)...)
	}

	transpiledDeclarations, transpiledCalls := transpileTemplates(templates, templatesCalls)

	for i, fileContent := range filesContents {
		filesContents[i] = transpileFile(fileContent, transpiledDeclarations, transpiledCalls)
	}

	for i, srcPath := range files {
		err := ioutil.WriteFile(strings.Replace(srcPath, "src", "dist", 1), []byte(filesContents[i]), 0644)
		check(err)
	}

}
