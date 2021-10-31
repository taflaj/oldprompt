# prompt
A simple, yet powerful, command line prompt displayer

**Note**: this is my original program but is not being maintained anymore; I'm currently working on a port to a different platform.

## Installation ##

Download the source, compile it, and place the binary `prompt` on your path. Add the following lines to `.bashrc`:

``` bash
PROMPT='<format string>'
eval "$(prompt init)"
```

### Format String ###

The format string contains a number or arguments in the form `<attribute>`=`<value>`, separated by a semicolon `;`. These are the currently supported arguments:

* status=`<fg>/<bg>` specifies the color of the status indicators to the left of the prompt.
* user=`<fg>/<bg>` specifies the color in the user section of the prompt.
* host=`<fg>/<bg>` specifies the color in the host section of the prompt.
* dir=`<fg>/<bg>` specifies the color in the directory section of the prompt.
* cozy=`yes`|`no`, if set to `no`, adds one extra space between sections.
* weight=`normal`|`bold`, if set to `bold`, displays in bold font.
* hostname=`full`|`top`, if set to `full`, displays the fully qualified hostname.
* limit=`<n>`, where `<n>` specifies the maximum length of the driectory section.

The foreground `<fg>` and background `<bg>` colors are calculated using the formula `16 + 36 * r + 6 * g + b`, where `r`, `g`, and `b` are respectively the red, green, and blue components, each on a scale of 0 to 5.

For example, `PROMPT="cozy=no; status=160/0; user=51/23; host=23/51; dir=23/231; limit=25; weight=bold; hostname=top"` shows the prompt in different shades of cyan using bold fonts and red status indicators. It only displays the top portion of the hostname and limits the directory section to 25 characters.

## Standalone Usage ##

Use `prompt help`.
