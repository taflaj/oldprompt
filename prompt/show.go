// prompt/prompt.go

package prompt

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	separator = "\ue0b0" // right triangle
	plain     = "\\[\x1b[0m\\]"
)

var (
	options map[string]string
	reset   string
	spaces  bool
)

func init() {
	log.SetFlags(log.Flags() | log.Lmicroseconds)
}

func getOptions() {
	options = make(map[string]string)
	splits := strings.Split(os.Getenv("options"), ";")
	for _, split := range splits {
		subsplits := strings.Split(split, "=")
		options[strings.Trim(subsplits[0], " ")] = strings.Trim(subsplits[1], " ")
	}
}

func setForeground(color string) string {
	if color == "" {
		return color
	}
	return "\\[\x1b[38;5;" + color + "m\\]"
}

func setBackground(color string) string {
	if color == "" {
		return color
	}
	return "\\[\x1b[48;5;" + color + "m\\]"
}

func getColors(s string) (string, string) {
	var fore, back string
	if s != "" {
		colors := strings.Split(s, "/")
		fore = colors[0]
		back = colors[1]
	}
	return fore, back
}

func getStatus() string {
	status := ""
	if os.Getenv("code") != "0" {
		status += "\uf12a" // exclamation
		if spaces {
			status += " "
		}
	}
	if os.Getenv("jobs") != "0" {
		status += "\uf252" // hourglass
		if spaces {
			status += " "
		}
	}
	if len(status) > 0 {
		fore, back := getColors(options["status"])
		status = setForeground(fore) + setBackground(back) + status
	}
	_, back := getColors(options["user"])
	return reset + status + setForeground(back) + "\ue0b6"
}

func addSpaces(s string) string {
	if spaces {
		return " " + s + " "
	}
	return s
}

func getUser() string {
	fore, back := getColors(options["user"])
	user := setForeground(fore) + setBackground(back) + addSpaces(os.Getenv("USER"))
	_, next := getColors(options["host"])
	return user + setForeground(back) + setBackground(next) + separator
}

func getHost() string {
	hostname, _ := os.Hostname()
	if options["hostname"] != "full" {
		hostname = strings.Split(hostname, ".")[0]
	}
	fore, back := getColors(options["host"])
	_, next := getColors(options["dir"])
	return setForeground(fore) + setBackground(back) + addSpaces(hostname) + setForeground(back) + setBackground(next) + separator
}

func getDir() string {
	dir, _ := os.Getwd()
	dir = strings.ReplaceAll(dir, os.Getenv("HOME"), "\uf015")
	limit, err := strconv.ParseInt(options["limit"], 10, 64)
	l := len(dir)
	if err == nil && l > int(limit) {
		// log.Printf("%v", l)
		sub := dir[l-int(limit):]
		p := strings.Index(sub, "/")
		dir = "\uf065" + sub[p:]
	}
	fore, back := getColors(options["dir"])
	return setForeground(fore) + setBackground(back) + addSpaces(dir) + reset + setForeground(back) + separator
}

// Show displays the prompt according to the parameters
func Show() {
	getOptions()
	reset = plain
	if options["weight"] == "bold" {
		reset = "\\[\x1b[0;1m\\]"
	}
	spaces = options["spaces"] == "yes"
	fmt.Print(getStatus() + getUser() + getHost() + getDir() + plain + " ")
}
