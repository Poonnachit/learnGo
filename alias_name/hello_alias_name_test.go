package main

import (
	"testing"
)

func TestAliasName(t *testing.T) {
	firstname := "Katy"
	lastname := "Perry"

	aliasName := getAliasName(firstname, lastname)
	if aliasName != "KaPe" {
		t.Fatalf(`Expected KaPe, got %s`, aliasName)
	}

	firstname = "Alex"
	lastname = "Goot"
	aliasName = getAliasName(firstname, lastname)
	if aliasName != "AlGo" {
		t.Fatalf(`Expected AlGo, got %s`, aliasName)
	}
}
