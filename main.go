// main.go

package main

import (
	"fmt"
	"os"

	"github.com/taflaj/prompt/prompt"
)

const version = "1.0.0"

func init() {}

func doHelp() {
	fmt.Printf("Usage: %v <command>\n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Println("  help     Displays this message.")
	fmt.Println("  init     Displays text to be used inside a shell script.")
	fmt.Println("  show     Displays the prompt according to the parameters.")
	fmt.Printf("  version  Displays the current version (%v).\n", version)
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
		case "version":
			fmt.Printf("%v %v\n", os.Args[0], version)
		default:
			fmt.Println("Invalid command")
			doHelp()
		}
	}
}
