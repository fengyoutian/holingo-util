package log

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	SetLevel(TRACE)

	dir, _ := os.Getwd()
	logger := New("LOGGER_TEST", dir)
	for i := 0; i < 10; i++ {
		t := i % 4
		if t == 0 {
			logger.Tracef("Test %d %% 4 = %d", i, t)
		} else if t == 1 {
			logger.Infof("Test %d %% 4 = %d", i, t)
		} else if t == 2 {
			logger.Warnf("Test %d %% 4 = %d", i, t)
		} else if t == 3 {
			logger.Errorf("Test %d %% 4 = %d", i, t)
		} else {
			logger.Panic(fmt.Sprintf("Test %d %% 4 = %d", i, t), fmt.Errorf("Error!"))
		}
	}
	logger.Fatal("Test fatal!")
}

func TestHLog_Trace(t *testing.T) {
	SetLevel(TRACE)

	dir, _ := os.Getwd()
	logger := New("LOGGER_TEST", dir)
	for i := 0; i < 10; i++ {
		logger.Tracef("Test i: %d", i)
	}
}

func TestHLog_Info(t *testing.T) {
	dir, _ := os.Getwd()

	logger := New("LOGGER_TEST", dir)
	for i := 0; i < 10; i++ {
		logger.Info(fmt.Sprintf("Test i: %d", i))
	}
}

func TestHLog_Warn(t *testing.T) {
	dir, _ := os.Getwd()

	logger := New("LOGGER_TEST", dir)
	for i := 0; i < 10; i++ {
		logger.Warn(fmt.Sprintf("Test i: %d", i))
	}
}

func TestHLog_Error(t *testing.T) {
	dir, _ := os.Getwd()

	logger := New("LOGGER_TEST", dir)
	for i := 0; i < 10; i++ {
		logger.Error(fmt.Sprintf("Test i: %d", i))
	}
}

func TestHLog_Fatal(t *testing.T) {
	dir, _ := os.Getwd()

	logger := New("LOGGER_TEST", dir)
	for i := 0; i < 10; i++ {
		logger.Fatal(fmt.Sprintf("Test i: %d", i))
	}
}

func TestHLog_Panic(t *testing.T) {
	dir, _ := os.Getwd()

	logger := New("LOGGER_TEST", dir)
	for i := 0; i < 10; i++ {
		logger.Panic(fmt.Sprintf("Test i: %d", i), fmt.Errorf("New error!"))
	}
}