package main

import (
	"github.com/darfk/ts3"
	"log"
	"strings"
)

func doLock(client *ts3.Client, clid string) {
	data, err := client.Exec(ClientInfo(clid))
	if err != nil {
		log.Fatal(err)
	}

	cldbid := data.Params[0]["client_database_id"]
	groups := strings.Split(data.Params[0]["client_servergroups"], ",")
	for i := range groups {
		if groups[i] == "9" {
			_, err := client.Exec(ServerGroupDelClient("9",cldbid))
			if err != nil {
				log.Fatal(err)
			}
			_, err = client.Exec(SendMessageUser(false, "Removed Lock group from you!", clid))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
