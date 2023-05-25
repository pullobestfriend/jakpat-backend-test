package main

import (
	"fmt"
	"strings"
)

type Room struct {
	Name    string
	Message []Message
}
type Message struct {
	User string
	Text string
}

var chatRoom = []Room{
	{
		Name: "backend",
		Message: []Message{
			{
				User: "Ari",
				Text: "Halo",
			},
			{
				User: "Ariyanto",
				Text: "Halo juga",
			},
		},
	},
	{
		Name: "frontend",
		Message: []Message{
			{
				User: "suryanto",
				Text: "Lorem ipsum",
			},
			{
				User: "Arindo",
				Text: "Helo",
			},
		},
	},
}

func soal2(input string) []string {
	var output []string
	for k, v := range searchIndex {
		if strings.Contains(v, strings.ToLower(input)) {
			output = append(output, k)
		}
	}
	return output
}

func main() {
	fmt.Println(strings.Join(soal2("Ari"), ", "))
	fmt.Println(strings.Join(soal2("lorem"), ", "))
}

var searchIndex = func() map[string]string {
	dict := make(map[string]string)
	for _, v := range chatRoom {
		var searchText string
		for _, msg := range v.Message {
			searchText += fmt.Sprintf("%s%s", strings.ToLower(msg.User), strings.ToLower(msg.Text))
		}
		dict[v.Name] = searchText
	}

	return dict
}()
