package main

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/dns"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

func main2() {
	client := cloudflare.NewClient(
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"), // defaults to os.LookupEnv("CLOUDFLARE_API_KEY")
		option.WithAPIEmail("user@example.com"),               // defaults to os.LookupEnv("CLOUDFLARE_EMAIL")
	)
	recordResponse, err := client.DNS.Records.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.RecordEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Record: dns.RecordParam{
				Name: "example.com",
			},
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", recordResponse)
}
