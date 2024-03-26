package main

import (
	"Gochat/gen"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type Connection struct {
	gen.UnimplementedBroadcastServer
	stream gen.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

type Pool struct {
	gen.UnimplementedBroadcastServer
	Connection []*Connection
}

func (p *Pool) CreateStream(pconn *gen.Connect, stream gen.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}

	p.Connection = append(p.Connection, conn)

	return <-conn.error
}

func (s *Pool) BroadcastMessage(ctx context.Context, msg *gen.Message) (*gen.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *gen.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				fmt.Printf("Sending message to: %v from %v", conn.id, msg.Id)

				if err != nil {
					fmt.Printf("Error with Stream: %v - Error: %v\n", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)

	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &gen.Close{}, nil
}

func main() {

	grpcServer := grpc.NewServer()

	var conn []*Connection

	pool := &Pool{
		Connection: conn,
	}

	gen.RegisterBroadcastServer(grpcServer, pool)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Error creating the server %v", err)
	}

	fmt.Println("Server started at port :8080")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error creating the server %v", err)
	}
}
