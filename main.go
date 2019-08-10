package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	host, port, timeout := os.Getenv("KAFKA_HOST"), os.Getenv("PORT"), os.Getenv("TIMEOUT")
	fmt.Println(fmt.Sprintf("begin ping kafka,%v:%v", host, port))
	t, err := strconv.ParseInt(timeout, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	count := int(t) / 2 //ping every two seconds
	for index := 0; index < count; index++ {
		if err := dailKafka(host, port); err != nil {
			time.Sleep(2 * time.Second)
			if index%30 == 0 {
				fmt.Printf("ping is err:%v,is trying... \n", err)
			}
			continue
		}
	}

	fmt.Println("ping kafka is ok")
	time.Sleep(100 * time.Hour)
}

func dailKafka(host, port string) (err error) {
	_, err = kafka.DialLeader(context.Background(), "tcp", host+":"+port, "ping", 0)
	if err != nil {
		return
	}
	return
}
