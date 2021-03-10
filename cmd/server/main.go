package main

import (
	"context"
	"flag"
	"github.com/alicefr/guestfs-server/libguestfs"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	path = flag.String("socket", "", "Socket to listen")
)

type server struct {
	libguestfs.UnimplementedVirtSparsifyServer
}

func main() {
	flag.Parse()
	if *path == "" {
		log.Fatalf("Specify the socket")
	}
	log.Infof("Server listening at %s", *path)

	lis, err := net.Listen("unix", *path)
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	libguestfs.RegisterVirtSparsifyServer(srv, &server{})

	log.Fatalln(srv.Serve(lis))
}

func (s *server) Sparsify(_ context.Context, image *libguestfs.Image) (*libguestfs.Response, error) {
	i := image.GetPath()
	err := libguestfs.SparsifyImage(i)
	if err != nil {
		log.Errorf("Faild to sparsify the image %s", i)
		return &libguestfs.Response{
			Success: false,
			Message: "Failed to sparsify image:" + i,
		}, err

	}
	log.Infof("Sparsify image")
	return &libguestfs.Response{Success: true}, nil
}
