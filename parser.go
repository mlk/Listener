package main

import "strings"

type Line struct {
	EventType Type
	User string
	Raw string
}

func (this Line) String() string {
	return this.EventType.String() + " " + this.User;
}

type Type int

func (this Type) String() string {
	switch this {
	case JOINED:
		return "[JOINED]"
	case LEFT:
		return "[LEFT]"
	case UNKNOWN:
		return "[UNKNOWN]"
	default:
		return "[default]"
	}
}

func getType(content []string) Type {
	if len(content) > 3 {
		if content[2] == "[JOIN]" {
			return JOINED
		}
		if content[2] == "[LEAVE]" {
			return LEFT
		}
	}

	return UNKNOWN
}

func getUser(currentType Type, content []string) string {
	if currentType != UNKNOWN {
		return content[3]
	}

	return ""
}

const (
	JOINED Type = iota
	LEFT Type = iota
	UNKNOWN Type = iota
)

func Parse(content string) Line {
	words := strings.Fields(content)
	theType := getType(words)
	return Line{
		EventType: theType,
		User: getUser(theType, words),
		Raw: content,
	}
}