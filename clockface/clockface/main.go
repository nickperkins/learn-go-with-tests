package main

import (
	"os"
	"time"

	"github.com/nickperkins/learn-go-with-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
