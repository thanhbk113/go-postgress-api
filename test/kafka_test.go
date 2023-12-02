package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"thanhbk113/pkg/admin/server/initialize"
	servicepost "thanhbk113/pkg/admin/services/posts"
	"time"
)

func TestAdd(t *testing.T) {
	var (
		expected int32
		wg       sync.WaitGroup
		postId   = "0289713c-d334-44aa-8731-7bac8667c77f"
		topic    = "like-post"
	)
	ctx := context.Background()

	initialize.Init()

	initialize.CreateTopic(topic)

	initialize.ListTopic()

	expected = 1

	intTimeNow := time.Now().Unix()
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func(i int, time int64) {
			fmt.Println("time", time+int64(i))
			message := fmt.Sprintf("like%s%d", postId, time+int64(i))
			key := fmt.Sprintf("like%s%d", postId, time+int64(i))
			initialize.SendMessage(topic, message, key)

			if initialize.MatchMessage(ctx, topic, key) {
				_ = servicepost.NewPostsService(ctx).TransactionLikePost(ctx, postId)
			}

			wg.Done()
		}(i, intTimeNow)
	}
	wg.Wait()

	result := servicepost.NewPostsService(ctx).GetTotalLikePost(ctx, postId)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
