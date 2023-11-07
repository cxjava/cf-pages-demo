package main

import (
	"github.com/jpillora/installer/handler"
	"github.com/syumai/workers"
)

var version = "0.0.0-src"

func main() {
	c := handler.DefaultConfig
	h := &handler.Handler{Config: c}
	workers.Serve(h) 
	
}
