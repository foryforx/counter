package api

import (
	"counter/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

func TestGetCounterHandler(t *testing.T) {
	seqGen := model.Initialize()
	handler := GetCounterHandler(seqGen)

	// Create a server to handle requests
	ts := httptest.NewServer(handler)
	defer ts.Close()

	var wg sync.WaitGroup
	wg.Add(200)

	numbers := make(map[int]bool)
	mu := &sync.Mutex{}

	for i := 0; i < 200; i++ {
		go func() {
			defer wg.Done()

			// Send a request to the server
			res, err := http.Get(ts.URL)
			if err != nil {
				t.Error(err)
				return
			}
			defer res.Body.Close()

			// Check the response
			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
				return
			}

			// The response should be a number
			num, err := strconv.Atoi(string(body))
			if err != nil {
				t.Errorf("Expected a number, got %s", body)
			}

			mu.Lock()
			if _, exists := numbers[num]; exists {
				t.Errorf("Duplicate number: %d", num)
			}
			numbers[num] = true
			mu.Unlock()
		}()
	}

	wg.Wait()

	// Check that all numbers from 1 to 200 were received
	for i := 1; i <= 200; i++ {
		if _, exists := numbers[i]; !exists {
			t.Errorf("Missing number: %d", i)
		}
	}
}
