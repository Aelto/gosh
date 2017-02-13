package main

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func encapsulate(str string, start string, end string) (string, int) {

	// startLength := utf8.RuneCountInString(start)
	endLength := utf8.RuneCountInString(end)

	// can't find any starting point in 'str'
	if strings.Index(str, start) == -1 {
		return "", -1
	}

	// can't find any end point in the str
	// return early from the first starting point met directly
	// to the end of 'str'
	if strings.Index(str, end) == -1 {
		return str[strings.Index(str, start):], utf8.RuneCountInString(str)
	}

	startIndex := strings.Index(str, start)
	endIndex := startIndex
	deep := 0

	for index := range str {
		textLeft := str[index:]

		if strings.Index(textLeft, start) == 0 {
			deep++
		}

		if strings.Index(textLeft, end) == 0 {
			endIndex = index + endLength

			if deep == 1 {
				break
			}

			deep--

		}

		if deep > 0 {
			endIndex++
		}

	}

	return str[startIndex:endIndex], endIndex

}

func goAfter(input string, after string) string {
	afterIndex := strings.Index(input, after)
	if afterIndex >= 0 {
		afterLength := utf8.RuneCountInString(after)

		return input[afterIndex+afterLength:]
	}

	return input
}

func removeEol(input string) string {
	re := regexp.MustCompile(`\r?`)

	return re.ReplaceAllString(input, "")
}

func goUntil(input string, until string) string {
	untilIndex := strings.Index(input, until)
	if untilIndex >= 0 {

		return input[:untilIndex]
	}

	return input
}

func goBetween(input, after, until string) string {
	return goUntil(goAfter(input, after), until)
}
