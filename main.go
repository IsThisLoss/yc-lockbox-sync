package main

import (
	"context"
	"flag"
	"log"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	"github.com/yandex-cloud/go-sdk/iamkey"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func main() {
	authKeys := flag.String("credentials", "", "Path to authorized keys of service account")
	secretID := flag.String("secret-id", "", "Your Yandex.Cloud Lockbox ID of the secret")
  flag.Parse()

  key, err := iamkey.ReadFromJSONFile(*authKeys)
  if err != nil {
    panic(err)
  }

  cred, err := ycsdk.ServiceAccountKey(key)
  if err != nil {
    panic(err)
  }

	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials:  cred,
	})
	if err != nil {
		panic(err)
	}

	p, err := sdk.LockboxPayload().Payload().Get(context.Background(), &lockbox.GetPayloadRequest{
		SecretId: *secretID,
	})
  if err != nil {
    panic(err)
  }


	log.Println(p)
}
