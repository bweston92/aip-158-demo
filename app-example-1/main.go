package main

import (
	"context"
	"flag"
	"log"

	"github.com/bweston92/aip-158-demo/svc-contacts/api"
	"google.golang.org/grpc"
)

var (
	contactsClient api.ContactsClient

	contactsAddr = flag.String("contacts-addr", ":3101", "Address that contacts is listening on")
)

func main() {
	ctx := context.Background()

	{
		conn, err := grpc.DialContext(ctx, *contactsAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("invalid configuration for contacts svc: %s", err)
		}
		contactsClient = api.NewContactsClient(conn)
	}

	res, err := contactsClient.ListContacts(ctx, &api.ListContactsRequest{})
	if err != nil {
		log.Fatalf("unable to get a list of contacts: %s", err)
	}

	log.Println("-----")
	for _, r := range res.Contacts {
		log.Printf("\tContact name: %s\n", r.Forename)
	}
	log.Println("-----")
	log.Printf("Total: %d\n", res.Total)
	log.Printf("Next page token: %s\n", res.NextPageToken)
}
