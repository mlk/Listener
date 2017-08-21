package main

import "testing"


func (actual Type) is(expected Type, t *testing.T) {
	if actual != expected {
		t.Errorf("Event Type was incorrect, got: '%s', want: '%s'.", actual, expected)
	}
}


func (actual Line) withUser(expected string, t *testing.T) {
	if actual.User != expected {
		t.Errorf("User was incorrect, got: '%s', want: '%s'.", actual.User, expected)
	}
}

func TestJoinEvent(t *testing.T) {
	actual := Parse("2017-08-16 11:04:17 [JOIN] mlk joined the game\n")
	actual.EventType.is(JOINED, t)
	actual.withUser("mlk", t)
}

func TestLeaveEvent(t *testing.T) {
	actual := Parse("2017-08-16 11:56:15 [LEAVE] mlk left the game\n")
	actual.EventType.is(LEFT, t)
	actual.withUser("mlk", t)
}

func TestUnknownString(t *testing.T) {
	actual := Parse("Biggles\n")
	actual.EventType.is(UNKNOWN, t)
	actual.withUser("", t)
}
