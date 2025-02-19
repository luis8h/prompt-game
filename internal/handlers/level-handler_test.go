package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"prompt-game/external/openai"
	"prompt-game/internal/stores"
	"runtime"
	"testing"

	"github.com/joho/godotenv")

func TestLevel1(t *testing.T) {
	messages := []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level1_v1_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level1_v1_a_1})
	testLevel(t, 0, true, true, messages)

	messages = []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level1_v2_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level1_v2_a_1})
	testLevel(t, 0, false, true, messages)
}

func TestLevel2(t *testing.T) {
	messages := []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level2_v1_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level2_v1_a_1})
	testLevel(t, 1, true, true, messages)
}

func TestLevel3(t *testing.T) {
	messages := []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level3_v1_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level3_v1_a_1})
	testLevel(t, 2, false, false, messages)

	messages = append(messages, openai.Message{Role: "user", Content: level3_v1_u_2})
	messages = append(messages, openai.Message{Role: "assistant", Content: level3_v1_a_2})
	testLevel(t, 2, true, true, messages)
}

func TestLevel4(t *testing.T) {
	messages := []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level4_v1_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level4_v1_a_1})
	testLevel(t, 3, true, true, messages)
}

func TestLevel5(t *testing.T) {
	messages := []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level5_v1_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level5_v1_a_1})
	testLevel(t, 4, true, true, messages)
}

func TestLevel6(t *testing.T) {
	messages := []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level6_v1_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level6_v1_a_1})
	testLevel(t, 5, true, true, messages)
}

func TestLevel7(t *testing.T) {
	messages := []openai.Message{}
	messages = append(messages, openai.Message{Role: "user", Content: level7_v1_u_1})
	messages = append(messages, openai.Message{Role: "assistant", Content: level7_v1_a_1})
	testLevel(t, 6, true, true, messages)
}

func testLevel(t *testing.T, levelId int, expectecAnswer bool, expectedStrategy bool, messages []openai.Message) {
	initEnv()
    levelHandler := NewLevelHandler(os.Getenv("OPENAI_API_KEY"))

	answer, strategy := levelHandler.validateLevel(nil, messages, stores.GetLevel(levelId, "en"))

	if answer != expectecAnswer {
		t.Fatalf(`Level %d: answer = %t, expected = %t, messages: %v`, levelId + 1, answer, expectecAnswer, messages);
	}

	if strategy != expectedStrategy {
		t.Fatalf(`Level %d: strategy = %t, expected = %t, messages: %v`, levelId + 1, strategy, expectedStrategy, messages);
	}
}

func getProjectRoot() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("unable to get current file path")
	}
	dir := filepath.Dir(filename)

	// Traverse upwards to find the go.mod file
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", fmt.Errorf("project root (with go.mod) not found")
		}
		dir = parentDir
	}
}

// initEnv loads the .env file from the project root.
func initEnv() {
	root, err := getProjectRoot()
	if err != nil {
		fmt.Println("Error finding project root:", err)
		return
	}

	envPath := filepath.Join(root, ".env")
	if err := godotenv.Load(envPath); err != nil {
		fmt.Println("No .env file found, relying on system environment variables.")
	}
}
