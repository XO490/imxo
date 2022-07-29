package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"imxo/proto"
	"log"
	"os"
	"strings"
)

var (
	addr = flag.String("addr", "localhost:4994", "the address to connect IMXO-server")
	uid  = flag.String("uid", "", "Your identifier")
	//	text = flag.String("text", "", "text message")
)

func main() {
	flag.Parse()

	//fmt.Println("Enter Server IP:Port >>")
	//reader := bufio.NewReader(os.Stdin)
	//serverID, err := reader.ReadString('\n')
	//if err != nil {
	//	log.Printf("Failed to read from console :: %v", err)
	//}
	//if serverID == "" {
	//	serverID = *addr
	//}
	//serverID = strings.Trim(serverID, "\r\n")

	serverID := *addr
	log.Println("Connecting: " + serverID)

	// connect to grpc server
	conn, err := grpc.Dial(serverID, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Faile to conncet to gRPC server :: %v", err)
	}
	defer conn.Close()

	// call ChatService to create a stream
	client := proto.NewSenderClient(conn)

	stream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}

	// implement communication with gRPC server
	ch := clientHandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	// blocker
	bl := make(chan bool)
	<-bl

}

// clientHandle
type clientHandle struct {
	stream     proto.Sender_SendMessageClient
	clientName string
}

func (ch *clientHandle) clientConfig() {

	//fmt.Printf("Your Name >> ")
	//reader := bufio.NewReader(os.Stdin)
	//username, err := reader.ReadString('\n')
	//if err != nil {
	//	log.Fatalf(" Failed to read from console :: %v", err)
	//}
	//if username == "" {
	//	username = *uid
	//}
	username := *uid
	ch.clientName = strings.Trim(username, "\r\n")

}

// send message
func (ch *clientHandle) sendMessage() {

	// create a loop
	for {

		//var clientMessage string
		//fmt.Print("[%s] << ", *uid)
		//fmt.Fscanf(os.Stdin, &clientMessage)

		fmt.Printf("[%s] << ", *uid)
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")
		log.Printf("[%s]: %s \n", *uid, clientMessage)

		clientMessageBox := &proto.FromClient{
			Uid:  ch.clientName,
			Text: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)
		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}

	}

}

// receive message
func (ch *clientHandle) receiveMessage() {

	// create a loop
	for {
		msg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server :: %v", err)
		}

		// print message to console
		log.SetFlags(2)
		fmt.Printf("[%s] >> %s \n", msg.Uid, msg.Text)

	}
}

//func main() {
//	flag.Parse()
//
//	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//
//	defer conn.Close()
//
//	c := pb.NewSenderClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	//var text string
//	//_, err = fmt.Scan(&text)
//	//if err != nil {
//	//	return
//	//}
//
//	r, err := c.SendMessage(ctx, &pr.FromClient{Uid: *uid, Text: *text})
//	if err != nil {
//		log.Fatalf("could not send: %v", err)
//	}
//	log.Printf("Message form %s", r.GetMessage())
//}
