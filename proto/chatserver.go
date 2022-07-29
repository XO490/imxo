package proto

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type message struct {
	ClientUid         string
	MessageText       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type messageHandle struct {
	MessageQueue []message
	mutex        sync.Mutex
}

var messageHandleObject = messageHandle{}

type ChatServer struct {
	SenderServer
	//proto.UnimplementedSenderServer
}

func (s *ChatServer) SendMessage(in Sender_SendMessageServer) error {
	clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	// receive message - init a goroutine
	go receiveFromStream(in, clientUniqueCode, errch)

	// send message - init a goroutine
	go sendToStream(in, clientUniqueCode, errch)

	return <-errch
}

// receive messages
func receiveFromStream(in_ Sender_SendMessageServer, clientUniqueCode_ int, errch_ chan error) {

	// implement a loop
	for {
		msg, err := in_.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
			errch_ <- err
		} else {

			messageHandleObject.mutex.Lock()

			messageHandleObject.MessageQueue = append(messageHandleObject.MessageQueue, message{
				ClientUid:         msg.Uid,
				MessageText:       msg.Text,
				MessageUniqueCode: rand.Intn(1e8),
				ClientUniqueCode:  clientUniqueCode_,
			})

			log.Printf("%v", messageHandleObject.MessageQueue[len(messageHandleObject.MessageQueue)-1])

			messageHandleObject.mutex.Unlock()

		}
	}
}

// send message
func sendToStream(in_ Sender_SendMessageServer, clientUniqueCode_ int, errch_ chan error) {

	// implement a loop
	for {

		// loop through messages in MessageQueue
		for {

			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mutex.Lock()

			if len(messageHandleObject.MessageQueue) == 0 {
				messageHandleObject.mutex.Unlock()
				break
			}

			senderUniqueCode := messageHandleObject.MessageQueue[0].ClientUniqueCode
			senderName4Client := messageHandleObject.MessageQueue[0].ClientUid
			message4Client := messageHandleObject.MessageQueue[0].MessageText

			messageHandleObject.mutex.Unlock()

			//send message to designated client (do not send to the same client)
			if senderUniqueCode != clientUniqueCode_ {

				err := in_.Send(&FromServer{Uid: senderName4Client, Text: message4Client})

				if err != nil {
					errch_ <- err
				}

				messageHandleObject.mutex.Lock()

				if len(messageHandleObject.MessageQueue) > 1 {
					messageHandleObject.MessageQueue = messageHandleObject.MessageQueue[1:] // delete the message at index 0 after sending to receiver
				} else {
					messageHandleObject.MessageQueue = []message{}
				}

				messageHandleObject.mutex.Unlock()

			}

		}

		time.Sleep(100 * time.Millisecond)
	}
}

//func (s *server) SendMessage(ctx, in *pb.FromClient) (*pb.FromServer, error) {
//	log.Printf("Message from %v > %v", in.GetUid(), in.GetText())
//	return &pb.FromServer{Uid: in.GetUid() + " > " + in.GetText()}, nil
//}
