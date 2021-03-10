package main

import (
	"context"
	"flag"
	"github.com/alicefr/guestfs-server/libguestfs"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

var (
	serverAddr = flag.String("socket", "", "The server socket path")
	path       = flag.String("path", "", "Path to the image")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *serverAddr == "" {
		log.Fatalf("Please set the socket path for the server")
	}
	addr := "unix://" + *serverAddr
	if *path == "" {
		log.Fatalf("Please set the path")
	}
	opts = append(opts, grpc.WithInsecure())
	i := &libguestfs.Image{
		Path: *path,
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := libguestfs.NewVirtSparsifyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = client.Sparsify(ctx, i)
	if err != nil {
		log.Fatalf("Failed to sparsify image %s: %v", i.Path, err)
	}

	log.Infof("Successfully sparisfied image")
}
