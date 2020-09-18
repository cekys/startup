package main

type configs []struct {
	Enabled bool   `json:"enabled"`
	Command string `json:"command"`
	Args    string `json:"args"`
}
