package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/MYKatz/gojam"
)

func usePreset(preset string) (*gojam.Markov, error) {
	m := gojam.NewMarkov(1, " ")
	var filepath string
	switch strings.ToLower(preset) {
	case "kanye":
		filepath = "presets/kanye.txt"
	case "beemovie":
		filepath = "presets/beemovie.txt"
	case "discord":
		filepath = "presets/discord.txt"
	default:
		return m, fmt.Errorf("invalid")
	}
	file, err := os.Open(filepath)
	if err != nil {
		return m, err
	}
	txt, err := ioutil.ReadAll(file)
	if err != nil {
		return m, err
	}
	err = m.FromJSON([]byte(txt))
	if err != nil {
		return m, err
	}
	return m, nil
}
