package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"thanhbk113/pkg/admin/server/initialize"
	servicepost "thanhbk113/pkg/admin/services/posts"

	"github.com/logrusorgru/aurora"
)

func BenchmarkTestMu(b *testing.B) {
	var (
		expected int32
		// wg       sync.WaitGroup
		postId = "0289713c-d334-44aa-8731-7bac8667c77f"
		mu     sync.Mutex
	)
	fmt.Println(aurora.Blue("BenchmarkMuTest"))
	ctx := context.Background()

	initialize.Init()

	expected = 1
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mu.Lock()
		_ = servicepost.NewPostsService(ctx).TransactionLikePost(ctx, postId)
		mu.Unlock()

	}

	result := servicepost.NewPostsService(ctx).GetTotalLikePost(ctx, postId)

	if result != expected {
		b.Errorf("Expected %d, but got %d", expected, result)
	}
}
