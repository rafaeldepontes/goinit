package builder

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"
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

func toPascalCase(src string) string {
	if strings.TrimSpace(src) == "" {
		return ""
	}

	var sb strings.Builder
	capNext := true

	for _, r := range src {
		if r == '-' || r == '_' || r == '.' {
			sb.WriteRune(r)
			capNext = true
			continue
		}

		if capNext {
			sb.WriteString(strings.ToUpper(string(r)))
			capNext = false
			continue
		}

		sb.WriteRune(r)
	}

	return sb.String()
}
