// prompt/prompt.go

package prompt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	clock       = "\uf017 "
	exclamation = "\uf12a "
	home        = "\uf015"
	separator   = "\ue0b0"
	shortcut    = "\uf0c4 "
	start       = "\ue0b6"
	ansiStart   = "\\[\x1b["
	ansiEnd     = "m\\]"
	plain       = ansiStart + "0" + ansiEnd
)

var (
	options map[string]string
	cozy    bool
)

func init() {}

func getOptions() {
	options = make(map[string]string)
	splits := strings.Split(os.Getenv("options"), ";")
	for _, split := range splits {
		subsplits := strings.Split(split, "=")
		if len(subsplits) == 2 {
			options[strings.Trim(subsplits[0], " ")] = strings.Trim(subsplits[1], " ")
		}
	}
}

func setForeground(color string) string {
	if color == "" {
		return ansiStart + "39" + ansiEnd
	}
	return ansiStart + "38;5;" + color + ansiEnd
}

func setBackground(color string) string {
	if color == "" {
		return ansiStart + "49" + ansiEnd
	}
	return ansiStart + "48;5;" + color + ansiEnd
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
		status += exclamation
	}
	jobs, err := strconv.ParseInt(os.Getenv("jobs"), 10, 0)
	if err == nil && jobs > 1 {
		status += clock
	}
	virtual := os.Getenv("VIRTUAL_ENV")
	if virtual != "" {
		p := strings.LastIndex(virtual, "/") + 1
		status += virtual[p:] + " "
	}
	toolbox := os.Getenv("TOOLBOX_NAME")
	if toolbox != "" {
		status += toolbox + " "
	}
	nixos := os.Getenv("IN_NIX_SHELL")
	if nixos != "" {
		status += nixos + " "
	}
	container := os.Getenv("container")
	cid := os.Getenv("CONTAINER_ID")
	if container != "" {
		if cid == "" {
			status += container + " "
		} else {
			status += container + ":" + cid + " "
		}
	} else if cid != "" {
		status += cid + " "
	}
	if len(status) > 0 {
		fore, back := getColors(options["status"])
		status = setForeground(fore) + setBackground(back) + status
	}
	_, back := getColors(options["user"])
	return status + setForeground(back) + start
}

func addSpaces(s string) string {
	if !cozy {
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
	if options["fullname"] != "yes" {
		hostname = strings.Split(hostname, ".")[0]
	}
	fore, back := getColors(options["host"])
	_, next := getColors(options["dir"])
	return setForeground(fore) + setBackground(back) + addSpaces(hostname) + setForeground(back) + setBackground(next) + separator
}

func getDir() string {
	dir, _ := os.Getwd()
	dir = strings.ReplaceAll(dir, os.Getenv("HOME"), home)
	limit, err := strconv.ParseInt(options["limit"], 10, 0)
	l := len(dir)
	if err == nil && l > int(limit) {
		sub := dir[l-int(limit):]
		p := strings.Index(sub, "/")
		if p < 0 {
			p = 0
		}
		dir = shortcut + sub[p:]
	}
	fore, back := getColors(options["dir"])
	_, next := getColors(options["command"])
	return setForeground(fore) + setBackground(back) + addSpaces(dir) + setForeground(back) + setBackground(next) + separator
}

func restOfLine() string {
	fore, back := getColors(options["command"])
	return setForeground(fore) + setBackground(back) + " "
}

func Show() {
	getOptions()
	reset := plain
	if options["weight"] == "bold" {
		reset = ansiStart + "0;1" + ansiEnd
	}
	cozy = options["cozy"] == "yes"
	fmt.Print(reset + getStatus() + getUser() + getHost() + getDir() + restOfLine())
}
