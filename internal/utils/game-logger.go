package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type GameLoggerType struct {
    *log.Logger
	logDir string
}

var GameLogger *GameLoggerType

func init() {
	logFile := "general-log.log"

	// load .env
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file:", err)
    }

    logDir := os.Getenv("LOG_DIR")
    if logDir == "" {
        logDir = "./log"
    }

	// Get the absolute path of the log directory
	absPath, err := filepath.Abs(logDir)
	if err != nil {
		fmt.Println("Error reading given path:", err)
		os.Exit(1)
	}

	// Check if the directory exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		err = os.MkdirAll(absPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}
	}

	// Open or create the log file
	generalLog, err := os.OpenFile(filepath.Join(absPath, logFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// Initialize the logger
    GameLogger = &GameLoggerType{
        Logger: log.New(generalLog, "Game Logger:\t", log.Ldate|log.Ltime|log.Lshortfile),
		logDir: logDir,
    }
}

func (gl *GameLoggerType) PrintS(sessionId string, message string) {
	fileName := fmt.Sprintf("%s.log", sessionId)
	logFilePath := filepath.Join(gl.logDir, fileName)

	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", logFilePath, err)
		return
	}
	defer file.Close()

    tmpLogger := log.New(file, "Session Logger:\t", log.Ldate|log.Ltime)
    tmpLogger.Printf("%s", message)
}
