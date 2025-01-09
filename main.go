package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/lockbox/v1"
	"github.com/yandex-cloud/go-sdk/iamkey"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func panicOnError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
	authKeys := flag.String("credentials", "", "Path to authorized keys of service account")
	secretID := flag.String("secret-id", "", "Your Yandex.Cloud Lockbox ID of the secret")
	secretKey := flag.String("secret-key", "", "Key inside the secret to save")
	dst := flag.String("dst", "", "Destination file")
  flag.Parse()

  key, err := iamkey.ReadFromJSONFile(*authKeys)
  panicOnError(err)

  cred, err := ycsdk.ServiceAccountKey(key)
  panicOnError(err)

	sdk, err := ycsdk.Build(context.Background(), ycsdk.Config{
		Credentials:  cred,
	})
  panicOnError(err)

	payload, err := sdk.LockboxPayload().Payload().Get(context.Background(), &lockbox.GetPayloadRequest{
		SecretId: *secretID,
	})
  panicOnError(err)

  data := []byte{}
  for _, entry := range payload.Entries {
    if entry.Key == *secretKey {
      binVal := entry.GetBinaryValue() 
      if binVal != nil {
        data = binVal
      } else {
        data = []byte(entry.GetTextValue())
      }
      break
    }
  }

  if len(data) == 0 {
    panic("Key was not found")
  }

  out, err := os.Create(*dst)
  panicOnError(err)

  out.Write(data)

	fmt.Println("Done")
}
