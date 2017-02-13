package main

import (
	"regexp"
	"strings"
)

type transpiledDeclaration struct {
	source     string
	transpiled string
}

type transpiledCall struct {
	source     string
	transpiled string
}

func transpileDeclaration(call templateCall, source templateDeclaration) transpiledDeclaration {
	if len(call.usedTypes) != len(source.usedTypes) {
		panic("Mismatched number of supplied types and needed types for template [" + source.name + "], call [" + call.call + "] from origin [" + call.origin)
	}

	transpiledSource := source.declaration

	for index, usedType := range call.usedTypes {
		reg := regexp.MustCompile("(\\b|^)(" + source.usedTypes[index] + ")(\\b|$)")

		transpiledSource = reg.ReplaceAllString(transpiledSource, usedType)
	}

	regGetNonAlpha := regexp.MustCompile("([^a-zA-Z]+)")
	transpiledSource = strings.Replace(transpiledSource, "func "+source.name+"(", "func gosh"+regGetNonAlpha.ReplaceAllString(source.name+strings.Join(call.usedTypes, ""), "_")+"(", 1)

	return transpiledDeclaration{source: source.declaration, transpiled: transpiledSource}
}

func transpileTemplates(declarations []templateDeclaration, calls []templateCall) ([]transpiledDeclaration, []transpiledCall) {
	transpiledDeclarations := []transpiledDeclaration{}
	transpiledCalls := []transpiledCall{}

	for _, declaration := range declarations {
		transpiledResult := ""

		for _, call := range calls {

			// create the transpiledDeclarations
			if declaration.name != call.name || len(declaration.usedTypes) != len(call.usedTypes) {
				continue
			}

			result := transpileDeclaration(call, declaration)
			if transpiledResult == "" {
				transpiledResult = transpiledResult + result.transpiled
			} else {
				transpiledResult = transpiledResult + "\n\n" + result.transpiled
			}
		}

		transpiledDeclarations = append(transpiledDeclarations, transpiledDeclaration{source: declaration.declaration, transpiled: transpiledResult})
	}

	regGetNonAlpha := regexp.MustCompile("([^a-zA-Z]+)")
	for _, call := range calls {
		// create the transpiledCalls

		transpiledCallString := "gosh" + call.name + strings.Join(call.usedTypes, "")
		transpiledCallString = regGetNonAlpha.ReplaceAllString(transpiledCallString, "_")
		transpiledCallString = strings.Replace(call.call, call.name, transpiledCallString, 1)
		log(transpiledCallString)

		transpiledCalls = append(transpiledCalls, transpiledCall{source: call.call, transpiled: transpiledCallString})
	}

	return transpiledDeclarations, transpiledCalls
}

func commentTemplateFlags(source string) string {
	return strings.Replace(strings.Replace(source, "@template", "//@ template", -1), ")@", ") // @", -1)
}

func transpileFile(file string, declarations []transpiledDeclaration, calls []transpiledCall) string {

	for _, declaration := range declarations {
		file = strings.Replace(file, declaration.source, declaration.transpiled, -1)
	}

	for _, call := range calls {
		file = strings.Replace(file, call.source, call.transpiled, -1)
		log("-")
		log(call.transpiled)
	}

	return commentTemplateFlags(file)
}
