package builder

import (
	"bufio"
	"context"
	"errors"
	"os"
)

// scanLine uses channels and a context to verify if the user chose to
// cancel the build, it also uses a Scanner from bufio to work.
func scanLine(ctx context.Context) (string, error) {
	ch := make(chan string, 1)
	errCh := make(chan error, 1)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			ch <- scanner.Text()
			return
		}
		if err := scanner.Err(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		return "", errors.New("Reverting changes...")
	case err := <-errCh:
		return "", err
	case line := <-ch:
		return line, nil
	}
}
