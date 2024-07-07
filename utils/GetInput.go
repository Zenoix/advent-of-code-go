package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

const inputUrlTemplate = "https://adventofcode.com/%d/day/%d/input"

func GetInput(year, day int) string {
	err := godotenv.Load(filepath.Join(GetCurrentDir(), `.env`))

	if err != nil {
		log.Fatal("Could not read .env file")
	}

	url := fmt.Sprintf(inputUrlTemplate, year, day)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	userAgent := "https://github.com/zenoix/advent-of-code-go/blob/main/utils/GetInput.go"
	req.Header.Set("User-Agent", userAgent)

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: os.Getenv("SESSION"),
	}

	req.AddCookie(&sessionCookie)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	output := strings.TrimRight(string(body), "\n")

	return output
}
