package server

import (
	"bolao-2026/internal/utils"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Setup

	server, errChan := Listen()

	port := utils.GetPort()
	address := fmt.Sprintf("http://localhost:%d/health", port)

	if err := waitForServer(address); err != nil {
		log.Fatalf("Failed to start test server: %v", err)
	}

	select {
	case err := <-errChan:
		log.Fatalf("Server crashed: %v", err)
	default:
		// Everything is fine
	}

	// Run TestXXX's in package

	code := m.Run()

	// Teardown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	os.Exit(code)
}

func waitForServer(url string) error {
	retries := 20
	wait := time.Duration(100)

	for i := 0; i < retries; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		time.Sleep(wait * time.Millisecond)
	}

	return fmt.Errorf("timeout waiting for server")
}

func TestListen(t *testing.T) {
	expected := "Healthy"

	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		t.Errorf("Fetch failed: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Read failed: %v", err)
		return
	}

	bodyStr := string(body)

	passes := strings.HasPrefix(bodyStr, expected)
	if !passes {
		t.Errorf("Check failed: %v", bodyStr)
	}

	t.Log(bodyStr)
}
