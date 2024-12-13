package main

import (
	"url"
)

type Module struct {
	Source  string
	Version string
	Name    string
}

func NewModule(sourceStr) &Module {
	u, err := url.Parse(sourceStr)
	return &Module{
		Source: url,
	}
}
