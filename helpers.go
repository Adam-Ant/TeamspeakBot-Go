package main

import "github.com/darfk/ts3"

func ClientGetDbIdFromUid(uid string) ts3.Command {
	return ts3.Command{
		Command: "´clientgetdbidfromuid",
		Params: map[string][]string{
			"cluid": []string{uid},
		},
	}
}

func ClientGetIds(uid string) ts3.Command {
	return ts3.Command{
		Command: "´clientgetids",
		Params: map[string][]string{
			"cluid": []string{uid},
		},
	}
}

func ClientGetNameFromUid(uid string) ts3.Command {
	return ts3.Command{
		Command: "´clientgetnamefromuid",
		Params: map[string][]string{
			"cluid": []string{uid},
		},
	}
}

func ClientGetNameFromDbId(dbid string) ts3.Command {
	return ts3.Command{
		Command: "´clientgetnamefromdbid",
		Params: map[string][]string{
			"cldbid": []string{dbid},
		},
	}
}

func ClientInfo(clid string) ts3.Command {
	return ts3.Command{
		Command: "clientinfo",
		Params: map[string][]string{
			"clid": []string{clid},
		},
	}
}

func ServerGroupDelClient(sgid string, cldbid string) ts3.Command {
	return ts3.Command{
		Command: "servergroupdelclient",
		Params: map[string][]string{
			"sgid":   []string{sgid},
			"cldbid": []string{cldbid},
		},
	}
}

func ServerGroupAddClient(sgid string, cldbid string) ts3.Command {
	return ts3.Command{
		Command: "servergroupaddclient",
		Params: map[string][]string{
			"sgid":   []string{sgid},
			"cldbid": []string{cldbid},
		},
	}
}

func SendMessageChannel(isError bool, message string) ts3.Command {
	//if isError {
	//	message = "'[B][COLOR=#ff0000]" + message + "[/COLOR][/B]"
	//}

	return ts3.Command{
		Command: "sendtextmessage",
		Params: map[string][]string{
			"targetmode": []string{"2"},
			"target":     []string{"1"},
			"msg":        []string{"Cockwomble"},
		},
	}
}

func SendMessageUser(isError bool, message string, clid string) ts3.Command {
	//if isError {
	//	message = "'[B][COLOR=#ff0000]" + message + "[/COLOR][/B]"
	//}

	return ts3.Command{
		Command: "sendtextmessage",
		Params: map[string][]string{
			"targetmode": []string{"1"},
			"target":     []string{clid},
			"msg":        []string{message},
		},
	}
}
