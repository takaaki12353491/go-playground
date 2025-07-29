package main

import (
	"context"
	"fmt"
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
