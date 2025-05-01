package main

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kbinani/screenshot"
)

func main() {
	// Create a folder to save screenshots
	godotenv.Load()
	fmt.Println(os.Getenv("AWS_ACCESS_KEY_ID"))
	outputDir := "screenshots"
	os.Create("orincec")
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Println("📸 Screenshot tracker started...")

	for {
		// Generate a random duration between 1 to 5 minutes
		randomMinutes := rand.Intn(5) + 1
		waitDuration := time.Duration(randomMinutes) * time.Minute
		fmt.Printf("⏳ Waiting for %v before taking next screenshot...\n", waitDuration)
		time.Sleep(waitDuration)

		// Capture and save screenshot
		err := captureAndSave(outputDir)
		if err != nil {
			fmt.Println("❌ Error capturing screenshot:", err)
		}
	}
}

func captureAndSave(outputDir string) error {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		return fmt.Errorf("no active displays found")
	}

	for i := 0; i < n; i++ {
		img, err := screenshot.CaptureDisplay(i)
		if err != nil {
			return err
		}

		// Create filename with timestamp
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		filename := fmt.Sprintf("screenshot_display%d_%s.png", i, timestamp)
		filepath := filepath.Join(outputDir, filename)

		// Save screenshot
		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer file.Close()
		err = png.Encode(file, img)
		if err != nil {
			return err
		}

		fmt.Printf("✅ Screenshot saved: %s\n", filepath)
	}
	return nil
}
