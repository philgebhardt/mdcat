# mdcat (Package)
This is a fork of [samfoo/mdcat](https://github.com/samfoo/mdcat), for the purposes of wrapping `mdcat`'s functionality into a package.

The following example would yeild the same result as invoking `mdcat` on a file containing `# Hello World!`.
```go

import "github.com/philgebhardt/mdcat"

func main() {
  markdown_string := fmt.Sprintf("# %s\n", "Hello World!") 
  mdcat.Print(strings.NewReader(markdown_string))
}
```

