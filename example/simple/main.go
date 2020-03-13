package main

import (
	"fmt"
	"github.com/davidji99/rollbar-go/rollbar"
)

func main() {
	client, newClientErr := rollbar.New(rollbar.AuthAAT("some_account_access_token"),
		rollbar.UserAgent("rollbar-go-custom"))

	if newClientErr != nil {
		fmt.Printf("Error: %v\n", newClientErr)
		return
	}

	fmt.Println(client)
}
