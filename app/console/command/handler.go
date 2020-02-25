package command

func Commands() map[string]map[string]string {
	return map[string]map[string]string{
		"say":  {"hello": "say hello to Goro"},
	}
}

func Handler(Row string,Command string,attr []string) {
	switch Row {
	case "say":
		say(Command, attr)
	}
}