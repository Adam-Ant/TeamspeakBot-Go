package main

import (
	"github.com/darfk/ts3"
	"log"
	"strings"
)

func doLockandGuest(client *ts3.Client, clid string) {
	var groupAssigned bool
	data, err := client.Exec(ClientInfo(clid))
	if err != nil {
		log.Fatal(err)
	}

	cldbid := data.Params[0]["client_database_id"]
	groups := strings.Split(data.Params[0]["client_servergroups"], ",")
	nickname := data.Params[0]["client_nickname"]

	for i := range groups {
		if groups[i] == "9" {
			_, err := client.Exec(ServerGroupDelClient("9", cldbid))
			if err != nil {
				log.Fatal(err)
			}

			log.Println("Removed lock group from client " + nickname)

			//_, err = client.Exec(SendMessageChannel(false, "Removed Lock group from you!"))
			//if err != nil {
			//	log.Println(err)
			//}

		}
		if groups[i] == "8" {
			groupAssigned = true
		}
	}

	log.Println(groupAssigned)

	if !groupAssigned {
		// Guest group missing, needs assigning.
		_, err := client.Exec(ServerGroupAddClient("8", cldbid))
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Gave guest group to client " + nickname)
	}

}
