package main

const (
	//DefaultConfig Default Config file path
	DefaultConfig = "startup.json"
)

var (
	conf        configs
	defaultConf = make(configs, 2)
)
