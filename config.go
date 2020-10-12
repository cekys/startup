package main

import "github.com/cekys/gopkg"

type config []struct {
	Enabled bool   `json:"enabled"`
	Command string `json:"command"`
	Args    string `json:"args"`
	Type    string `json:"type"`
}

func (conf *config) ReadFromFile(filename string) error {

	jsoner := gopkg.Jsoner{Jst: &conf}
	err := jsoner.ReadFromFile(filename)
	if err != nil {
		return err
	}
	return nil
}

func (conf *config) WriteToFile(filename string) error {
	jsoner := gopkg.Jsoner{Jst: &conf}
	err := jsoner.WriteToFile(filename)
	if err != nil {
		return err
	}
	return nil
}
