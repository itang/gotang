package gotang

import (
	"fmt"
	"time"
)

// 执行IO操作并设定超时时间，超时返回超时错误
func DoIOWithTimeout(iof func() error, t time.Duration) error {
	timeout := time.NewTicker(t)
	defer timeout.Stop()

	done := make(chan error)
	go func() {
		done <- iof()
	}()

	select {
	case <-timeout.C:
		return fmt.Errorf("Do IO timeout: %v", t)
	case err := <-done:
		return err
	}
}
