package main

import (
	"os"
	"testing"
)

func TestGetCep(t *testing.T) {
    os.Args = []string{"main", "01454000"}

    exp := "01454000"

    res := getCep()

    if res != exp {
        t.Errorf("Expected %s, got %s instead.\n", exp, res)
    }
}

