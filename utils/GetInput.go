package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const inputUrlTemplate = "https://adventofcode.com/%d/day/%d/input"

// You must fill this value in with your session. To find it,
// go to https://adventofcode.com/2023/day/1/input and look in
// the cookie of the Get request.
const session = "CHANGE ME"

func GetInput(year, day int) string {
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
		Value: session,
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
