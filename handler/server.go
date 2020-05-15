package handler

import (
	"context"

	"github.com/mikudos/mikudos_message_deliver/config"
	pb "github.com/mikudos/mikudos_message_deliver/proto/message-deliver"
)

// Server 事件驱动服务间流程控制方法，提供基本的数据库操作方法
type Server struct {
	pb.MessageDeliverServer
	DelaySeconds int
	deleteChan   chan *pb.MessageReceivedRequest
}

// InitServer InitServer
func InitServer() *Server {
	server := Server{DelaySeconds: config.RuntimeViper.GetInt("delaySeconds"), deleteChan: make(chan *pb.MessageReceivedRequest)}
	go server.handleDelChan()
	return &server
}

// MessageReceived 新建聚合，返回聚合Id
func (s *Server) MessageReceived(ctx context.Context, req *pb.MessageReceivedRequest) (*pb.MessageReceivedResponse, error) {
	go s.waitAndPushToDelete(req)
	return &pb.MessageReceivedResponse{}, nil
}
