package main

type config []struct {
	Enabled bool   `json:"enabled"`
	Command string `json:"command"`
	Args    string `json:"args"`
	Type    string `json:"type"`
}
