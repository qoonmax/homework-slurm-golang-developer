package main

import "fmt"

type message struct {
	id       uint
	authorId uint
	body     *string
}

func main() {
	//Создадим новое сообщение
	body := "Hello world!"
	msg := message{
		1,
		5,
		&body,
	}

	// Канал для сообщений
	msgChan := make(chan message)

	// Отправим сообщение
	go func(msg *message) {
		defer close(msgChan)
		msgChan <- *msg
	}(&msg)

	for m := range msgChan {
		m.id++
		m.authorId++
		*m.body = "Modified message"

		fmt.Printf("Id: %d, AuthorId: %d, Body: %s \n", m.id, m.authorId, *m.body)
	}

	fmt.Printf("Id: %d, AuthorId: %d, Body: %s \n", msg.id, msg.authorId, *msg.body)
}
