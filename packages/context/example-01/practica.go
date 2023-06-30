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

	ctx := context.Background()

	val, err := getUserData(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", val)
	fmt.Println("took:", time.Since(start))

}

func getUserData(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*550)
	defer cancel()

	responsch := make(chan Response)

	go func() {
		value, err := itWillTakeSomeMilliseconds()
		responsch <- Response{
			value: value,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("Tardo demasiado")
		case responsch := <-responsch:
			return responsch.value, nil
		}
	}

}

func itWillTakeSomeMilliseconds() (int, error) {

	time.Sleep(time.Millisecond * 500)

	return 1, nil

}
