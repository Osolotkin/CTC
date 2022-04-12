//package main
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	gogo "project/project/gogo"

	"google.golang.org/grpc"

	etcd "go.etcd.io/etcd/client/v3"
)

const ADDR string = ":";
const PORT string = "8080";
const DB_ENDPOINT string = ":2379";

var db *etcd.Client;




type Service struct {
	gogo.UnimplementedGoGoServiceServer
}

func (s *Service) Ping(ctx context.Context, in *gogo.Message) (*gogo.Message, error) {
	log.Printf("Received message: %s", in.Body);
	return &gogo.Message{Body: "Pong"}, nil;
}

func (s *Service) Get(ctx context.Context, in *gogo.Message) (*gogo.Message, error) {

	keyStr := in.Body;

	ctx, cancel := context.WithTimeout(context.Background(), time.Second);
    rsp, err := db.Get(ctx, keyStr);
    cancel();

    if err != nil {
		log.Printf("Get failed: %s", err);
        return &gogo.Message{Body: "Get failed!"}, nil;
    }

	var buffer bytes.Buffer;
	fmt.Printf("Values returned:\n");
    for _, ev := range rsp.Kvs {
		str := fmt.Sprintf("%s:%s\n", ev.Key, ev.Value);
        fmt.Printf(str);
		buffer.WriteString(str);
    }

	outBytes := buffer.Bytes();
	if (len(outBytes) > 0) {
		return &gogo.Message{Body: string(outBytes[0 : len(outBytes) - 1])}, nil;
	} else {
		return &gogo.Message{Body: ""}, nil;
	}

}

func (s *Service) Post(ctx context.Context, in *gogo.KeyValuePair) (*gogo.Message, error) {

	keyStr := in.Key;
	valStr := in.Val;

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    _, err := db.Put(ctx, keyStr, valStr);
	cancel();

    if err != nil {
		log.Printf("Post failed: %s", err);
        return &gogo.Message{Body: "Post failed!"}, nil;
    }

	log.Printf("Post success");
	return &gogo.Message{Body: "Post success!"}, nil;

}

func (s *Service) List(ctx context.Context, in *gogo.Message) (*gogo.Message, error) {

	keyStr := in.Body;

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    rsp, err := db.Get(ctx, keyStr, etcd.WithPrefix());
	cancel();

    if err != nil {
		log.Printf("List failed: %s", err);
        return &gogo.Message{Body: "List failed!"}, nil;
    }

	var buffer bytes.Buffer;
	fmt.Printf("Values returned:\n");
    for _, ev := range rsp.Kvs {
		str := fmt.Sprintf("%s:%s\n", ev.Key, ev.Value);
        fmt.Printf(str);
		buffer.WriteString(str);
    }

	outBytes := buffer.Bytes();
	if (len(outBytes) > 0) {
		return &gogo.Message{Body: string(outBytes[0 : len(outBytes) - 1])}, nil;
	} else {
		return &gogo.Message{Body: ""}, nil;
	}

}

func (s *Service) Delete(ctx context.Context, in *gogo.Message) (*gogo.Message, error) {

	keyStr := in.Body;

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    rsp, err := db.Delete(ctx, keyStr);
	cancel();

    if err != nil {
		log.Printf("Delete failed: %s", err);
        return &gogo.Message{Body: "Delete failed!"}, nil;
    }
	log.Printf("Delete success!");
	return &gogo.Message{Body: rsp.Header.String()}, nil;
	

}

func main() {
  
	// listetn to a port

	lis, err := net.Listen("tcp", ADDR + PORT);
  	if err != nil {
		log.Fatalf("Failed to listen: %v", err);
		return;
	}
	fmt.Printf("Listening %s%s...\n", ADDR, PORT);


	// start grpc service

	grpcSrv := grpc.NewServer();
	service := Service{};
	gogo.RegisterGoGoServiceServer(grpcSrv, &service);

	fmt.Printf("GRPC server started, service registered...\n");


	// start db

	db, err = etcd.New(etcd.Config{
        Endpoints:   []string{DB_ENDPOINT},
        DialTimeout: 5 * time.Second,
    });
	if err != nil {
        log.Fatalf("Failed to connect to the etcd, err:%v\n", err);
        return;
    }
	fmt.Printf("Databse etcd initialized on endpoint: %s...\n", DB_ENDPOINT);

	defer db.Close();



	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err);
		return;
	}

}
