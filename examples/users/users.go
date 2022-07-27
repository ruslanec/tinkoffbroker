package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ruslanec/tinkoffbroker/service/users"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

const investApiUrl = "invest-public-api.tinkoff.ru:443"

func main() {
	token := flag.String("t", "", "tinkoff API token")
	accountid := flag.String("a", "", "tinkoff account id")
	flag.Parse()

	if len(*token) == 0 {
		*token = os.Getenv("TKF_TOKEN")
	}

	if len(*accountid) == 0 {
		*accountid = os.Getenv("TKF_ACCOUNTID")
	}

	// create connection to broker Tinkoff
	conn, err := grpc.Dial(investApiUrl,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: *token,
		})))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute))
	defer func() {
		cancel()
	}()

	users := users.NewUsersService(conn, *accountid)
	accounts, err := users.Accounts(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, account := range accounts {
		s, _ := json.MarshalIndent(account, "", "\t")
		fmt.Println(string(s))
	}
}
