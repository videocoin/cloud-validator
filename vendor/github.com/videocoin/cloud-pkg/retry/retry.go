package retry

import (
	"fmt"
	"time"
)

func RetryWithAttempts(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)
	}

	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
