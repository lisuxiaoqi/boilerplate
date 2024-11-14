package main

import (
	"context"
	"fmt"
	client "github.com/gatechain/gatenode-openrpc"
	"github.com/gatechain/gatenode-openrpc/types/blob"
	"github.com/gatechain/gatenode-openrpc/types/share"
	"testing"
)

func TestSubmitBlob(t *testing.T) {
	ctx := context.Background()
	// bridge node client
	//url := "http://localhost:26658"
	url := "http://124.243.187.49:26658"
	apiToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.gMp8-YzfPT5fpf8YXS5U36mntQFRvKUFzrDmgowa6tY"
	//apiToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.0Trbr1oIKZlb39Q3ArLH-Ly89bVSy33Vw9xE3qb8vI8"

	client, err := client.NewClient(ctx, url, apiToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// let's post to 0xDEADBEEF namespace
	namespace, err := share.NewBlobNamespaceV0([]byte{0xDE, 0xAD, 0xBE, 0xEF})
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a blob
	helloWorldBlob, err := blob.NewBlobV0(namespace, []byte("Hello, World!"))
	if err != nil {
		return
	}

	// submit the blob to the network
	height, err := client.Blob.Submit(ctx, []*blob.Blob{helloWorldBlob}, blob.NewSubmitOptions())
	if err != nil {
		fmt.Println(err)
		return
	}

	//_ = height
	extendedHeader, err := client.Header.WaitForHeight(context.Background(), height)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Submit: height %d\n", height)
	fmt.Println("extended Header\n", extendedHeader)

	// bonus: fetch the blob back from the network
	retrievedBlobs, err := client.Blob.GetAll(ctx, height, []share.Namespace{namespace})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("GetAll:", retrievedBlobs)
	//// retrieve data back from DA
	daBlob, err := client.Blob.Get(ctx, height, namespace, helloWorldBlob.Commitment)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get:", daBlob)

}
