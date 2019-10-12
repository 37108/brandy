package main

import (
	"fmt"

	"github.com/37108/brandy/config"
	"github.com/37108/brandy/finder"
)

func main() {
	profiles := config.Parser()
	idx, err := finder.Finder(profiles)
	if err != nil {
		fmt.Println(err)
	}
	profile := profiles[idx]

	sess, err := config.ObtainSession(profile)
	if err != nil {
		fmt.Println(err)
	}

	config.SessionToEnv(sess)
}
