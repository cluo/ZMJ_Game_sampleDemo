package cg

import (
	"fmt"
)

type Player struct {
	Name  string `json:"Name"`
	Level int    `json:"level"`
	Exp   int    `json:"exp"`
	Room  int    `json:"room"`

	mq chan *Message `json:"message"` //等待接受消息...
}

func NewPlayer() *Player {
	m := make(chan *Message)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "recieved message:", msg.Content)
		}
	}(player)

	return player
}
