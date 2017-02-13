package main

import (
	"regexp"
	"strings"
)

type templateDeclaration struct {
	name        string
	declaration string
	usedTypes   []string

	// stores the origin of the template declaration
	// it should be a path to the file containing the
	// template declaration
	origin string
}

type templateCall struct {
	name      string
	usedTypes []string
	call      string

	// stores the origin of the template call
	// it should be a path to the file containing the
	// template call
	origin string
}

func searchTemplateDeclarationNext(source string, sourcePath string) []templateDeclaration {
	templates := []templateDeclaration{}

	currentSlice := source[:]
	for {
		nextTemplateIndex := strings.Index(currentSlice, "@template")
		if nextTemplateIndex == -1 {
			break
		}

		templateBody, _ := encapsulate(goAfter(currentSlice[nextTemplateIndex:], " {\n"), "{", "}")
		templateDeclarationString := goUntil(currentSlice[nextTemplateIndex:], templateBody) + templateBody
		templateName := goBetween(templateDeclarationString, "func ", "(")

		templateUsedTypes := strings.Split(removeEol(goBetween(templateDeclarationString, "@template ", "\n")), " ")

		templates = append(templates, templateDeclaration{name: templateName, declaration: templateDeclarationString, origin: sourcePath, usedTypes: templateUsedTypes})

		currentSlice = goAfter(currentSlice, templateDeclarationString)
	}

	return templates
}

func containsTemplateCall(list []templateCall, call templateCall) bool {
	for _, callInList := range list {
		if sliceStringEqual(callInList.usedTypes, call.usedTypes) && callInList.name == call.name {
			return true
		}
	}

	return false
}

func searchTemplateCallNext(source string) []templateCall {
	calls := []templateCall{}

	lines := strings.Split(source, "\n")
	for _, line := range lines {
		if strings.Index(line, ")@") == -1 {
			continue
		}

		// the regex does however stops at a '(' but it also includes it
		// removing it is needed
		regexSearchAnyWord, _ := regexp.Compile("(\\S+\\()")
		templateName := goUntil(regexSearchAnyWord.FindString(line), "(")
		templateCallString := templateName + goAfter(line, templateName)
		templateTypes := strings.Split(removeEol(goAfter(templateCallString, "@")), " ")

		newCall := templateCall{name: templateName, usedTypes: templateTypes, call: templateCallString}

		// skip this call if a call with the same
		// types has already been declared
		if containsTemplateCall(calls, newCall) {
			continue
		}

		calls = append(calls, newCall)
	}

	return calls
}
