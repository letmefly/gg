package main

import (
	"context"
	//"errors"
	"log"
	"reflect"
	"sync"
)
import (
	"actors"
	"pb"
	"pb/auth"
	"services"

	"google.golang.org/grpc/metadata"
)

type Server struct {
	auth.GrpcServer
	users sync.Map
}

func (s *Server) Init(ctx context.Context) {
}

func (s *Server) CreateUser(userId string, out chan interface{}) *User {
	user, ok := s.GetUser(userId)
	if !ok {
		user = NewUser(userId, out)
		s.users.Store(userId, user)
	}
	return user
}

func (s *Server) GetUser(userId string) (*User, bool) {
	user, ok := s.users.Load(userId)
	if !ok {
		return nil, false
	}
	return user.(*User), true
}

// ------------------------------------- grpc api -------------------------------------
func (s *Server) CreateStream(stream pb.Stream_CreateStreamServer) error {
	log.Println("new stream is coming..")
	meta, _ := metadata.FromIncomingContext(stream.Context())
	userId := meta["user-id"][0]
	codec := meta["codec"][0]
	out := make(chan interface{}, 1000)
	user := s.CreateUser(userId, out)

	go func() {
		for {
			select {
			case outMsg := <-out:
				msgName := reflect.TypeOf(outMsg)
				msgId := services.ToMsgId(msgName.String())
				data, _ := auth.EncodeMessage(codec, outMsg)
				stream.Send(&pb.StreamFrame{
					Type:    pb.StreamFrameType_Message,
					Codec:   codec,
					MsgId:   int32(msgId),
					MsgData: data,
				})
			case <-stream.Context().Done():
				return
			}
		}
	}()

	defer func() {
		close(out)
	}()

	// receive loop
	for {
		frame, err := stream.Recv()
		if err != nil {
			log.Println(err)
			return err
		}
		codec := frame.Codec
		msgName := services.ToMsgName(uint32(frame.MsgId))
		msgData := frame.MsgData
		msg, decodeErr := auth.DecodeMessage(codec, msgName, msgData)
		if decodeErr != nil {
			log.Println(decodeErr)
		} else {
			actors.AsynCall(user.ActorId(), (*User).HandleMessages, msgName, msg)
		}

		switch frame.Type {
		case pb.StreamFrameType_Message:
		case pb.StreamFrameType_Ping:
		case pb.StreamFrameType_Kick:
		}
	}
	return nil
}

func (s *Server) Login(ctx context.Context, param *auth.LoginParam) (*auth.LoginRet, error) {
	log.Println("Login")
	//return nil, errors.New("this api not support")
	return &auth.LoginRet{Error: "ok", UserId: 100000}, nil
}
