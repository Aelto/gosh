package main

func hasValidOption(args []string, option string, defaultValue string) string {

	indexOf := func(s []string, target string) int {
		for i, element := range s {
			if element == target {
				return i
			}
		}

		return -1
	}

	optionIndex := indexOf(args, option)
	if optionIndex > -1 && optionIndex+1 <= len(args) {
		return args[optionIndex+1]
	}

	return defaultValue

}
