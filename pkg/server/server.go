package server

import (
	"context"
	"errors"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/gallak87/goskel/proto/userpb"
	"net"
	"net/http"
)

var (
	ErrNotFound = errors.New("not found")
)

// TODO: replace with db and use go-pg

type User struct {
	ID   int64
	Name string
}
type UserTable map[string]User
var UserRepo UserTable = make(UserTable)
type userServer struct{}

func (s *userServer) GetUser(ctx context.Context, request *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		Id: "1",
		Name: "Username",
	}, nil
}

func (s *userServer) RegisterUser(ctx context.Context, request *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	return &userpb.RegisterUserResponse{
		Success: false,
	}, nil
}

func (s *userServer) DeleteUser(ctx context.Context, request *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	return &userpb.DeleteUserResponse{
		Success: false,
	}, nil
}

func newUserServer() userpb.UserServiceServer {
	return new(userServer)
}

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, newUserServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}



// RunInProcessGateway starts the invoke in process http gateway.
// TODO: Implement grpc gateway
func RunInProcessGateway(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)

	//userpb.RegisterEchoServiceHandlerServer(ctx, mux, newEchoServer())

	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http gateway server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http gateway server: %v", err)
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil

}