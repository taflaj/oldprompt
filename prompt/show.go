// prompt/prompt.go

package prompt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	separator = ``
	ansiStart = "\\[\x1b["
	ansiEnd   = "m\\]"
	plain     = ansiStart + "0" + ansiEnd
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

func getTime() string {
	if options["time"] != "yes" {
		return ""
	}
	time, _ := strconv.ParseInt(os.Getenv("time"), 10, 64)
	if time == 0 {
		return ""
	}
	now, _ := strconv.ParseInt(os.Getenv("now"), 10, 64)
	ms := (now - time) / 1000000
	seconds := ms / 1000
	ms %= 1000
	minutes := seconds / 60
	seconds %= 60
	hours := minutes / 60
	minutes %= 60
	return fmt.Sprintf("󱑂 %d:%02d:%02d.%03d\n", hours, minutes, seconds, ms)
}

func getStatus() string {
	status := ""
	if os.Getenv("code") != "0" {
		status += ` `
	}
	jobs, _ := strconv.ParseInt(os.Getenv("jobs"), 10, 0)
	if jobs > 0 {
		status += `  `
	}
	virtual := os.Getenv("VIRTUAL_ENV")
	if virtual != "" {
		p := strings.LastIndex(virtual, "/") + 1
		status += virtual[p:] + " "
	}
	nixos := os.Getenv("IN_NIX_SHELL")
	if nixos != "" {
		status += nixos + " "
	}
	container := os.Getenv("container")
	cid := os.Getenv("CONTAINER_ID")
	toolbox := os.Getenv("TOOLBOX_NAME")
	if container != "" {
		if cid == "" && toolbox == "" {
			status += container + " "
		} else {
			status += container + ":"
			if cid != "" {
				status += cid
			} else {
				status += toolbox
			}
			status += " "
		}
	} else if cid != "" {
		status += cid + " "
	} else if toolbox != "" {
		status += toolbox + " "
	}
	_, next := getColors(options["user"])
	if len(status) > 0 {
		fore, back := getColors(options["status"])
		status = setBackground(back) + setForeground(fore) + status
	}
	return status + setForeground(next) + ``
}

func addSpaces(s string) string {
	if !cozy {
		return " " + s + " "
	}
	return s
}

func getUser() string {
	fore, back := getColors(options["user"])
	user := setForeground(fore) + setBackground(back) + addSpaces("\\u")
	_, next := getColors(options["host"])
	return user + setForeground(back) + setBackground(next) + separator
}

func getHost() string {
	hostname := "\\H"
	if options["fullname"] != "yes" {
		hostname = "\\h"
	}
	fore, back := getColors(options["host"])
	_, next := getColors(options["dir"])
	return setForeground(fore) + setBackground(back) + addSpaces(hostname) + setForeground(back) + setBackground(next) + separator
}

func getDir() string {
	dir, _ := os.Getwd()
	dir = strings.ReplaceAll(dir, os.Getenv("HOME"), ``)
	limit, err := strconv.ParseInt(options["limit"], 10, 0)
	l := len(dir)
	if err == nil && l > int(limit) {
		sub := dir[l-int(limit):]
		p := strings.Index(sub, "/")
		if p < 0 {
			p = 0
		}
		dir = ` ` + sub[p:]
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
	fmt.Print(reset + getTime() + getStatus() + getUser() + getHost() + getDir() + restOfLine())
}
