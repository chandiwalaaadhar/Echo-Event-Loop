package eventComponents

import (
	"fmt"
)

func EventLoop(queue *Queue) {
	for {
		event := queue.Dequeue() // Dequeue an eventComponents

		if event.Message == "" {
			continue // If the eventComponents is empty, continue
		}

		fmt.Fprintf(event.Client, "Echo: %s\n", event.Message)                 // Echo the message back to the client
		fmt.Printf("Echoed to Client %d: %s\n", event.ClientID, event.Message) // Print to the server console that the message was echoed to the client
	}
}
