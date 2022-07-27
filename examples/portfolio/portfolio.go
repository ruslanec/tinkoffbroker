package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ruslanec/tinkoffbroker"
	"github.com/ruslanec/tinkoffbroker/service/operations"
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

	operations := operations.NewOperationsService(conn, *accountid)

	client := tinkoffbroker.NewClient(conn, *accountid, tinkoffbroker.WithOperations(operations))
	//defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute))
	defer func() {
		cancel()
	}()

	portfolio, err := client.Portfolio(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Общая стоимость валют в портфеле в рублях: ", portfolio.TotalAmountCurrencies.String())
	for _, position := range portfolio.Positions {
		// s, _ := json.MarshalIndent(position, "", "\t")
		// fmt.Println(string(s))
		fmt.Println(position.String())
	}

}
