package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// Разрешаем функции работать не более 2 секунд. В результате получим или ответ за время до 2 сек. или ошибку
	fmt.Println(TooSlowWrap(ctx, time.Second*2))
}

// TooSlowWrap разрешает функции работать не более переданного duration. В результате получим или ответ
// за допустимое время или ошибку, что функция работает слишком долго.
func TooSlowWrap(ctx context.Context, timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ch := make(chan int)
	defer close(ch)

	go func() {
		ch <- TooSlowFunc()
	}()

	select {
	case <-ctx.Done():
		return 0, errors.New("too slow")
	case v := <-ch:
		return v, nil
	}
}

// TooSlowFunc имитация медленной функции.
func TooSlowFunc() int {
	time.Sleep(time.Second * 4)

	return 100
}
