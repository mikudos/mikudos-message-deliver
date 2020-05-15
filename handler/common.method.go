package handler

import (
	"time"

	"github.com/mikudos/mikudos_message_deliver/clients"
	pb "github.com/mikudos/mikudos_message_deliver/proto/message-deliver"
)

func (s *Server) handleDelChan() {
	for {
		req := <-s.deleteChan
		clients.RemoveMessage(clients.GenRemoveRequest(req))
	}
}

func (s *Server) waitAndPushToDelete(req *pb.MessageReceivedRequest) {
	time.Sleep(5 * time.Second)
	s.deleteChan <- req
}
