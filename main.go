package main

import (
	"github.com/darfk/ts3"
	"log"
	"time"
)

// Username and password in Constants.go, not on github :D

func processNotify(client *ts3.Client, notification chan ts3.Notification) {
	for i := range notification {
		switch i.Type {
		case "notifyclientmoved", "notifycliententerview":
			// We need to check for lock group on every move, and guest check is best done at the same time.
			doLockandGuest(client, i.Params[0]["clid"])
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

	// Start listening to text chat
	_, err = client.Exec(ts3.Command{
		Command: "servernotifyregister",
		Params: map[string][]string{
			"event": []string{"textchannel"},
			"id":    []string{"1"},
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
		time.Sleep(500 * time.Millisecond)
	}
}
