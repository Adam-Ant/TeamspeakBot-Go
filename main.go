package main

import (
	"github.com/darfk/ts3"
	"log"
)

// Username and password in Constants.go, not on github :D

func processNotify(client *ts3.Client, notification chan ts3.Notification) {
	for i := range notification {
		//log.Println(i)
		if i.Type == "notifyclientmoved" {
			// We need to check for lock group
			doLock(client, i.Params[0]["clid"])
		}
	}
}

func main() {

	client, err := ts3.NewClient("ts3.cwgaming.co.uk:10011")
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Exec(ts3.Login(user, pass))
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Exec(ts3.Use(1))
	if err != nil {
		log.Fatal(err)
	}

	// Start listening to client join and move events
	_, err = client.Exec(ts3.Command{
		Command: "servernotifyregister",
		Params: map[string][]string{
			"event": []string{"channel"},
			"id":    []string{"0"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	notification := make(chan ts3.Notification)

	go processNotify(client, notification)

	client.NotifyHandler(func(n ts3.Notification) {
		notification <- n
	})

	for {
	}
}
