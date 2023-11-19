package model

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := Initialize()

	// Test that the counter starts at 1
	if val := counter.GetCounter(); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Test that the counter increments correctly
	if val := counter.GetCounter(); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	// Test that the counter increments correctly
	if val := counter.GetCounter(); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	// Test that the counter increments correctly
	if val := counter.GetCounter(); val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}

	// Test that the counter increments correctly
	if val := counter.GetCounter(); val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}

	// Test that the counter increments correctly
	if val := counter.GetCounter(); val != 6 {
		t.Errorf("Expected 6, got %d", val)
	}

	// Test that the counter stops correctly
	counter.Stop()
	time.Sleep(time.Second) // Give the counter time to stop

	// Test that the counter does not increment after stopping
	if val := counter.GetCurrentCounter(); val != 6 {
		t.Errorf("Expected 6, got %d", val)
	}
}

func TestCounterWithFiveCallers(t *testing.T) {
	counter := Initialize()

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			counter.GetCounter()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			counter.GetCounter()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			counter.GetCounter()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			counter.GetCounter()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			counter.GetCounter()
		}
	}()

	wg.Wait()

	// Test that the counter increments correctly
	if val := counter.GetCurrentCounter(); val != 500 {
		t.Errorf("Expected 500, got %d", val)
	}

	// Test that the counter stops correctly
	counter.Stop()
	time.Sleep(time.Second) // Give the counter time to stop

	// Test that the counter does not increment after stopping
	if val := counter.GetCurrentCounter(); val != 500 {
		t.Errorf("Expected 500, got %d", val)
	}
}
