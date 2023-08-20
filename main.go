package main

import (
	"bufio"
	"eventLoop/eventComponents"
	"fmt"
	"net"
)

var clientID int

func main() {
	eventQueue := &eventComponents.Queue{} // Create a new eventComponents queue

	go eventComponents.EventLoop(eventQueue) // Start the eventComponents loop, passing in the eventComponents queue,
	// this is a single thread which will loop forever and process events

	ln, err := net.Listen("tcp", ":8080") // Listen on port 8080
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening on port 8080...") // Print to the server console that the server is listening on port 8080
	for {
		conn, err := ln.Accept() // Accept a connection
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(conn, eventQueue) // Handle the client in a new thread,
		// the eventComponents queue is passed in so that the client can enqueue events, multiple threads can enqueue events at the same time
	}
}

func handleClient(client net.Conn, queue *eventComponents.Queue) {
	defer client.Close() // Close the client when the function ends

	clientID++                  // Increment the client ID
	currentClientID := clientID // Assign the current client ID to the client ID variable

	fmt.Printf("Client %d connected: %s\n", currentClientID, client.RemoteAddr())

	scanner := bufio.NewScanner(client) // Create a scanner for the client

	for scanner.Scan() {
		message := scanner.Text() // Get the message from the client
		if message == "EXIT" {
			fmt.Fprintf(client, "DISCONNECTED\n")                   // Notify the client of disconnection
			fmt.Printf("Client %d disconnected\n", currentClientID) // Print to the server console of disconnection
			return                                                  // Break out of the loop and disconnect the client
		}
		queue.Enqueue(eventComponents.Event{Message: message, Client: client, ClientID: currentClientID}) // Enqueue the eventComponents
	}
}
