package main

import (
	"time"

	"github.com/firework-a/easy-banner"
)

func main() {
	err := banner.Print(map[string]string{
		"Version":   "1.0.0",
		"BuildTime": time.Now().Format("2006-01-02"),
		"Author":    "Your Name",
	})
	if err != nil {
		return
	}
}
