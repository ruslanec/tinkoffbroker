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
	"github.com/ruslanec/tinkoffbroker/domain"
	"github.com/ruslanec/tinkoffbroker/service/marketdatastream"
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

	marketdatastream := marketdatastream.NewMarketDataStreamService(conn)

	client := tinkoffbroker.NewClient(conn, *accountid, tinkoffbroker.WithMarketDataStream(marketdatastream))
	//defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute))
	defer func() {
		cancel()
	}()

	var instruments []*domain.InfoInstrument
	instruments = append(instruments, &domain.InfoInstrument{
		Figi: "BBG00T22WKV5",
	})
	err = client.SubscribeInfo(ctx, instruments)
	if err != nil {
		log.Fatal(err)
	}

	err = client.MySubscriptions(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		recv, err := client.Recv(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(recv)
	}

	// inter, err := client.Recv(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Total amount currencies: %s", inter)

	// inter, err = client.Recv(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Total amount currencies: %s", inter)

	// inter, err = client.Recv(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Total amount currencies: %s", inter)
}
