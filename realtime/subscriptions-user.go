package realtime

import (
	"log"
)

// Subscribes to stream-notify-logged
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *Client) SubUser(name string, args ...interface{}) (chan string, error) {

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
