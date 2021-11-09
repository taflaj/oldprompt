# prompt
A simple, yet powerful, command line prompt displayer

## Installation ##

1. Download the source and compile it, or otherwise download the proper executable file from the `bin` directory.
2. Place the binary as `prompt` on your path.
3. Add the following lines to `.bashrc`:

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
* command=`<fg>/<bg>` specifies the color after the end of the prompt (where the user types the command).
* cozy=`yes`|`no`, if set to `no`, adds one extra space between sections. Default is `no`.
* weight=`normal`|`bold`, if set to `bold`, displays in bold font. Default is `normal`.
* fullname=`yes`|`no`, if set to `yes`, displays the fully qualified hostname. Default is `no`.
* limit=`<n>`, where `<n>` specifies the maximum length of the directory section. Default is unlimited.

The foreground `<fg>` and background `<bg>` colors are calculated using the formula `16 + 36 * r + 6 * g + b`, where `r`, `g`, and `b` are respectively the red, green, and blue components, each on a scale of 0 to 5. If a color is not specified, then it assumes the default foreground or background color.

For example, `PROMPT='cozy=no; status=88/226; user=226/58; host=58/226; dir=58/231; command=226/; limit=25; weight=bold; fullname=no'` shows the prompt in different shades of yellow using bold fonts and red status indicators. It only displays the top portion of the hostname and limits the directory section to 25 characters. Commands typed by the user are displayed in bold bright yellow on a transparent background.

## Standalone Usage ##

Use `prompt help`.
