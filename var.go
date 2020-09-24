package main

const (
	DefaultConfig  = "startup.json"
	DefaultSimple  = "simple"
	DefaultForking = "forking"
)

var (
	conf        config
	defaultConf = make(config, 2)
)
