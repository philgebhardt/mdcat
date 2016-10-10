package mdcat

import (
	// "bufio"
	// "flag"
	// "fmt"
	"github.com/russross/blackfriday"
	// "io/ioutil"
	"strings"
	"os"
)

func Print(reader *strings.Reader) {


	renderer := &Console{}
	extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS


	// reader := bufio.NewReader(os.Stdin)

	var input []byte
	buffer := make([]byte, 2<<20)

	for {
		count, err := reader.Read(buffer)

		if count == 0 {
			break
		}

		if err != nil {
			os.Stderr.WriteString("mdcat: unable to read from pipe\n")
			os.Exit(1)
		}

		input = append(input, buffer...)
	}

	output := blackfriday.Markdown(input, renderer, extensions)
	os.Stdout.Write(output)

}
