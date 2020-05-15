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
}

// InitServer InitServer
func InitServer() *Server {
	return &Server{DelaySeconds: config.RuntimeViper.GetInt("delaySeconds")}
}

// MessageReceived 新建聚合，返回聚合Id
func (s *Server) MessageReceived(ctx context.Context, req *pb.MessageReceivedRequest) (*pb.MessageReceivedResponse, error) {
	// stmt, _ := db.Db.Prepare(`INSERT INTO aggregate (aggregate_type, data) VALUES (?, ?)`)
	// defer stmt.Close()
	// ret, err := stmt.Exec(req.GetAggregateType(), req.GetData())
	// if err != nil {
	// 	fmt.Printf("insert data error: %v\n", err)
	// 	return nil, err
	// }
	// LastInsertId, _ := ret.LastInsertId()
	return &pb.MessageReceivedResponse{}, nil
}
