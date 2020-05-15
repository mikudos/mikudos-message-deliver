package clients

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mikudos/mikudos_message_deliver/config"
	deliver "github.com/mikudos/mikudos_message_deliver/proto/message-deliver"
	pb "github.com/mikudos/mikudos_message_deliver/proto/message-pusher"
	"google.golang.org/grpc"
)

var (
	conns = make(map[string]*grpc.ClientConn)
	// MessagePusherClient MessagePusherClient
	MessagePusherClient pb.MessagePusherClient
	err                 error
	callIndex           = 1
)

// MessagePusherServiceName MessagePusherServiceName
const MessagePusherServiceName = "message_pusher"

func init() {
	log.Println("Init all grpc client: ai, learn, users, messages")
	setUpClientConn(MessagePusherServiceName)
}

func setUpClientConn(connName string) {
	confLoc := fmt.Sprintf("grpcClients.%s", connName)
	grpcAddr := config.RuntimeViper.GetString(confLoc)
	if grpcAddr == "" {
		log.Fatalln("address for " + confLoc + "must be set")
	}
	// Set up a connection to the server.
	conns[connName], err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	MessagePusherClient = pb.NewMessagePusherClient(conns[connName])
}

// GenRemoveRequest GenRemoveRequest
func GenRemoveRequest(req *deliver.MessageReceivedRequest) *pb.DeliverRemoveRequest {
	return &pb.DeliverRemoveRequest{ChannelId: req.GetChannelId(), MsgId: req.GetMsgId()}
}

// RemoveMessage RemoveMessage
func RemoveMessage(removeRequest *pb.DeliverRemoveRequest) {
	state := conns[MessagePusherServiceName].GetState()
	if state.String() != "READY" {
		conns[MessagePusherServiceName].Close()
		setUpClientConn(MessagePusherServiceName)
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := MessagePusherClient.DeliverRemoveMessage(ctx, removeRequest)
	log.Printf("SayHello called %d times", callIndex)
	callIndex++
	if err != nil {
		log.Printf("could not call method on %s: %v", MessagePusherServiceName, err)
	} else {
		log.Printf("call return: %v", r.GetResult())
	}
}
