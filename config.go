package main

import "github.com/cekys/gopkg"

type config []struct {
	Enabled bool   `json:"enabled"`
	Command string `json:"command"`
	Args    string `json:"args"`
	Type    string `json:"type"`
}

func (conf *config) ReadFromFile(filename string) error {
	j, err := gopkg.NewJsoner(conf)
	if err != nil {
		return err
	}

	err = j.ReadFromFile(filename)
	if err != nil {
		return err
	}
	return nil
}

func (conf *config) WriteToFile(filename string) error {
	j, err := gopkg.NewJsoner(conf)
	if err != nil {
		return err
	}

	err = j.WriteToFile(filename)
	if err != nil {
		return err
	}
	return nil
}
