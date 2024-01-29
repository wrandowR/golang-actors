package main

import (
 "fmt"
 "sync"
)

// Actor represents an actor with its own state and a channel for receiving messages.
type Actor struct {
 state int
 mailbox chan int
}

// NewActor creates a new actor with an initial state.
func NewActor(initialState int) *Actor {
 return &Actor{
  state: initialState,
  mailbox: make(chan int),
 }
}

// ProcessMessage processes a message by updating the actor's state.
func (a *Actor) ProcessMessage(message int) {
 fmt.Printf("Actor %d processing message: %d\n", a.state, message)
 a.state += message
}

// Run simulates the actor's runtime by continuously processing messages from the mailbox.
func (a *Actor) Run(wg *sync.WaitGroup) {
 defer wg.Done()
 for {
  message := <-a.mailbox
  a.ProcessMessage(message)
 }
}

// System represents the actor system managing multiple actors.
type System struct {
 actors []*Actor
}

// NewSystem creates a new actor system with a given number of actors.
func NewSystem(numActors int) *System {
 system := &System{}
 for i := 1; i <= numActors; i++ {
  actor := NewActor(i)
  system.actors = append(system.actors, actor)
  go actor.Run(nil)
 }
 return system
}

// SendMessage sends a message to a randomly selected actor in the system.
func (s *System) SendMessage(message int) {
 actorIndex := message % len(s.actors)
 s.actors[actorIndex].mailbox <- message
}

func main() {
 // Create an actor system with 3 actors.
 actorSystem := NewSystem(3)

 // Send messages to the actors concurrently.
 var wg sync.WaitGroup
 for i := 1; i <= 5; i++ {
  wg.Add(1)
  go func(message int) {
   defer wg.Done()
   actorSystem.SendMessage(message)
  }(i)
 }

 // Wait for all messages to be processed.
 wg.Wait()
}
