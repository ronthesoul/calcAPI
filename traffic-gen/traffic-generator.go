// traffic-gen/traffic-generator.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	rand "math/rand/v2" // Go 1.22+: no seeding needed
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	baseURL = "http://localhost:8080"
	client  = &http.Client{Timeout: 5 * time.Second}
)

type tokenResp struct {
	APIKey string `json:"API-Key"`
}

func getToken(ctx context.Context) (string, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/token", nil)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("token %s: %s", resp.Status, string(b))
	}
	var tr tokenResp
	if err := json.Unmarshal(b, &tr); err != nil {
		return "", err
	}
	if tr.APIKey == "" {
		return "", fmt.Errorf("no API-Key in token response: %s", string(b))
	}
	return tr.APIKey, nil
}

func callOp(ctx context.Context, op string, a, b int, token string) (int, []byte, error) {
	url := fmt.Sprintf("%s/%s/%d/%d", baseURL, op, a, b)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if token != "" {
		req.Header.Set("X-API-Key", token)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, body, nil
}

func looksInvalid(body []byte) bool {
	var m map[string]any
	if json.Unmarshal(body, &m) == nil {
		if s, ok := m["error"].(string); ok {
			s = strings.ToLower(s)
			return strings.Contains(s, "invalid") || strings.Contains(s, "expired")
		}
	}
	return false
}

func Start(ctx context.Context) {
	var token string
	ops := []string{"add", "sub", "multiply", "divide"}

	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			op := ops[rand.IntN(len(ops))]
			a := rand.IntN(100) + 1
			b := rand.IntN(100) + 1
			if op == "divide" && b == 0 {
				b = 1
			}

			if token == "" {
				if tok, err := getToken(ctx); err == nil {
					token = tok
				} else {
					fmt.Println("token err:", err)
					continue
				}
			}

			cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			status, body, err := callOp(cctx, op, a, b, token)
			cancel()
			if err != nil {
				fmt.Println("call err:", err)
				continue
			}

			if status == http.StatusUnauthorized || looksInvalid(body) {
				if tok, err := getToken(ctx); err == nil {
					token = tok
					cctx, cancel = context.WithTimeout(ctx, 3*time.Second)
					_, _, err = callOp(cctx, op, a, b, token)
					cancel()
					if err != nil {
						fmt.Println("retry err:", err)
					}
				} else {
					fmt.Println("refresh err:", err)
				}
			}

			// Uncomment to see responses:
			// fmt.Printf("%s %d %d -> %s\n", op, a, b, string(body))
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	Start(ctx)
}
