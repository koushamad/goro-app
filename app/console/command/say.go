package command

import "github.com/fatih/color"

func say(Command string, args []string) {
	switch Command {
	case "hello":
		sayHello(args)
	}
}

func sayHello(args []string) {
	color.Red("hello from goro")
}
