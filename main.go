// main.go

package main

import (
	"fmt"
	"os"

	"github.com/taflaj/prompt/prompt"
)

func init() {}

func doHelp() {
	fmt.Printf("Usage: %v <command>\n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Println("  help  Displays this message and exits.")
	fmt.Println("  init  Displays text to be used inside a shell script.")
	fmt.Println("  show  Displays the prompt according to the parameters.")
}

func main() {
	if len(os.Args) == 1 {
		doHelp()
	} else {
		switch os.Args[1] {
		case "help":
			doHelp()
		case "init":
			fmt.Println("PROMPT_COMMAND=set_prompt")
			fmt.Println("set_prompt() {")
			// fmt.Println("  PS1=\"$(code=$? jobs=$(jobs -p | wc -l) options=$PROMPT go run github.com/taflaj/prompt show)\" ")
			fmt.Println("  PS1=\"$(code=$? jobs=$(jobs -p | wc -l) options=$PROMPT prompt show)\" ")
			fmt.Println("}")
		case "show":
			prompt.Show()
		default:
			fmt.Println("Invalid command")
			doHelp()
		}
	}
}
