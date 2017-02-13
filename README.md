# Gosh.go
transpile your go code (into go) to get c++ like templates

# How to install
`go install github.com/aelto/gosh`

# How to use
in the gosh directory `$ gosh -in ./example/src` to transpile the example folder. It uses relative paths

## Creating a template function
```go
@template T
func printSlice(list T) {
	for _, el := range list {
		fmt.Println(el)
	}
}
```

## using a template function
```go
mySliceOfString := []string{"Hello", "World"}
mySliceOfInt := []int{1, 4, 5, 9}

printSlice(mySliceOfString)@[]string
printSlice(mySliceOfInt)@[]int
```

## Transpile your code
> your code should be in a directory named `src` (this condition won't be there in the next versions)

In the current version of gosh, the first occurence on `src` in the supplied path to the source code is replaced by `dist`.
Imagine the example `gosh -in ./example/src`, that'll transpile everything in the supplied folder. The transpiled code will then be written in `example/dist`.


# What's planned
- [ ] make the transpiling work with statements like `variableName = templateFunctionCall(args...)@types` to allow storage of returned values in variables
- [ ] Allow a `-out` argument for greater flexibility on where the transpiled code should be written
- [ ] Watch mode so typing `gosh -in path/to/src` isn't needed every change