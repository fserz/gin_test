package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

func testTx() {
	// 此处rdb为初始化的redis连接客户端
	const routineCount = 100

	// 设置5秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// increment 是一个自定义对key进行递增（+1）的函数
	// 使用 GET + SET + WATCH 实现，类似 INCR
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			// 获得当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 实际操作（乐观锁定中的本地操作）
			n++

			// 仅在监视的Key保持不变的情况下运行
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				// pipe 处理错误情况
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// 最多重试100次
		for retries := routineCount; retries > 0; retries-- {
			err := rdb.Watch(ctx, txf, key)
			if err != redis.TxFailedErr {
				return err
			}
			// 乐观锁丢失
		}
		return errors.New("increment reached maximum number of retries")
	}

	// 开启100个goroutine并发调用increment
	// 相当于对key执行100次递增
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(ctx, "counter3").Int()
	fmt.Println("最终结果：", n, err)

}
