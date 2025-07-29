package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func loopWithErrGroup() {
	eg, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 5; i++ {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if i == 3 {
					return fmt.Errorf("タスク %d でエラーが発生", i)
				}
				fmt.Printf("タスク %d が完了\n", i)
				time.Sleep(100 * time.Millisecond)
				return nil
			}
		})
	}

	// いずれかのゴルーチンがエラーを返した場合、g.Wait()はそのエラーを返す
	if err := eg.Wait(); err != nil {
		fmt.Printf("エラー発生: %v\n", err)
	} else {
		fmt.Println("すべてのタスクが成功")
	}
}

func loopWithoutErrGroup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var once sync.Once
	var firstErr error
	var mu sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		id := i
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				if id == 3 {
					err := fmt.Errorf("タスク %d でエラーが発生", id)

					once.Do(func() {
						mu.Lock()
						firstErr = err
						mu.Unlock()
						cancel()
					})

					return
				}

				fmt.Printf("タスク %d が完了\n", id)
				time.Sleep(100 * time.Millisecond)
			}
		}(id)
	}

	wg.Wait()

	if firstErr != nil {
		fmt.Printf("エラー発生: %v\n", firstErr)
	} else {
		fmt.Println("すべてのタスクが成功")
	}
}
