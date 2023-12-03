package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"thanhbk113/pkg/admin/server/initialize"
	"thanhbk113/pkg/admin/server/initialize/kafka"
	servicepost "thanhbk113/pkg/admin/services/posts"
	"time"

	"github.com/logrusorgru/aurora"
)

//go test -bench=. ./test/ben_kafka.go ( command to run benchmark test)

func BenchmarkTestKafka(b *testing.B) {
	var (
		expected int32
		// wg       sync.WaitGroup
		postId  = "0289713c-d334-44aa-8731-7bac8667c77f"
		topic   = "like-post"
		holdMes = ""
	)
	ctx := context.Background()

	initialize.Init()

	kafka.CreateTopic(topic)

	kafka.ListTopic()

	expected = 1
	fmt.Println(aurora.Blue("BenchmarkTestKafka"))
	b.ResetTimer()

	intTimeNow := time.Now().Unix()
	// wg.Add(2)
	for i := 0; i < b.N; i++ {
		// go func(i int, time int64) {
		fmt.Println("time", intTimeNow+int64(i))
		message := fmt.Sprintf("like%s%d", postId, intTimeNow+int64(i))
		key := fmt.Sprintf("like%s%d", postId, intTimeNow+int64(i))
		kafka.SendMessage(topic, message, key)

		if holdMes == "" {
			holdMes = message
		}

		// wg.Done()
		// }(i, intTimeNow)
	}

	// wg.Wait()

	if kafka.MatchMessage(ctx, topic, holdMes) {
		_ = servicepost.NewPostsService(ctx).TransactionLikePost(ctx, postId)
	}
	result := servicepost.NewPostsService(ctx).GetTotalLikePost(ctx, postId)

	if result != expected {
		b.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestAdd(t *testing.T) {
	var (
		expected int32
		wg       sync.WaitGroup
		postId   = "0289713c-d334-44aa-8731-7bac8667c77f"
		topic    = "like-post"
		holdMes  = ""
	)
	ctx := context.Background()

	initialize.Init()

	kafka.CreateTopic(topic)

	kafka.ListTopic()

	expected = 1

	intTimeNow := time.Now().Unix()
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(i int, time int64) {
			fmt.Println("time", time+int64(i))
			message := fmt.Sprintf("like%s%d", postId, time+int64(i))
			key := fmt.Sprintf("like%s%d", postId, time+int64(i))
			kafka.SendMessage(topic, message, key)

			if holdMes == "" {
				holdMes = message
			}

			wg.Done()
		}(i, intTimeNow)
	}

	wg.Wait()

	if kafka.MatchMessage(ctx, topic, holdMes) {
		_ = servicepost.NewPostsService(ctx).TransactionLikePost(ctx, postId)
	}
	result := servicepost.NewPostsService(ctx).GetTotalLikePost(ctx, postId)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
