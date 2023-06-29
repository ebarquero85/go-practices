// Reference: https://www.youtube.com/watch?v=kaZOXRqFPCw
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

func main() {

	start := time.Now()

	ctx := context.WithValue(context.Background(), "userId", 10)

	val, err := fetchUserData(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))

}

func fetchUserData(ctx context.Context) (int, error) {

	// This is just for testing passing value in context
	userId := ctx.Value("userId")
	fmt.Println("UserId is:", userId)

	// Set maximum time for this app to get data
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel() // Not required but necesary if we want to aboid leaking contexts

	responchannel := make(chan Response)

	go func() {
		value, err := fetchThirdPartyStuffWhichCanBeSlow()
		responchannel <- Response{
			value: value,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")
		case resp := <-responchannel:
			return resp.value, nil
		}
	}

}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {

	time.Sleep(time.Millisecond * 500)

	return 666, nil

}
