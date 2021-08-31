package realtime

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gopackage/ddp"
)

// Subscribes to stream-notify-logged
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *Client) Sub(name string, args ...interface{}) (chan string, error) {

	if args == nil {
		log.Println("no args passed")
		if err := c.ddp.Sub(name); err != nil {
			return nil, err
		}
	} else {
		if err := c.ddp.Sub(name, args[0], false); err != nil {
			return nil, err
		}
	}

	msgChannel := make(chan string, default_buffer_size)
	c.ddp.CollectionByName(name).AddUpdateListener(genericExtractor{msgChannel, "update"})

	return msgChannel, nil
}

type genericExtractor struct {
	messageChannel chan string
	operation      string
}

func (u genericExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	if operation == u.operation {
		b, _ := json.Marshal(doc)
		u.messageChannel <- string(b) //fmt.Sprintf("%s -> update", collection)
	}
	// for _, listener := range c.listeners {
	// 	listener.CollectionUpdate(c.Name, operation, id, doc)
	// }
}

type genericExtractor2 struct {
	messageChannel chan string
}

func (u genericExtractor2) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	u.messageChannel <- fmt.Sprintf("%s -> update", collection)
}
