package main

import (
	"context"
	"flag"

	"github.com/dexidp/dex/api/v2"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

var address string
var username, password, email string

func init() {
	flag.StringVar(&address, "address", "127.0.0.1:5557", "target url")
	flag.StringVar(&email, "email", "tianqiuhuang@gmail.com", "email address")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&username, "username", "TQ", "username")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		klog.Fatal(err)
	}

	userID, err := uuid.NewV4()
	if err != nil {
		klog.Fatal(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		klog.Fatal(err)
	}

	client := api.NewDexClient(conn)
	req := api.CreatePasswordReq{
		Password: &api.Password{
			UserId:   userID.String(),
			Username: username,
			Hash:     hashedPassword,
			Email:    email,
		},
	}

	_, err = client.CreatePassword(context.Background(), &req)
	if err != nil {
		klog.Fatal(err)
	}
}
