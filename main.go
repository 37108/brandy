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
		panic(err)
	}
	profile := profiles[idx]

	client, err := config.Client(profile)
	if err != nil {
		panic(err)
	}
	credentials, _ := client.Credentials.Retrieve()
	fmt.Printf("export AWS_ACCESS_KEY_ID=%s export AWS_SECRET_ACCESS_KEY=%s export AWS_SESSION_TOKEN=%s", credentials.AccessKeyID, credentials.SecretAccessKey, credentials.SessionToken)
}
