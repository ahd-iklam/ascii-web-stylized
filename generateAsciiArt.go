package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func excuteAsciiArtResult(w http.ResponseWriter, r *http.Request) {
	// catch the value from the request query parameters
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if text == "" || banner == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	art, err := generateAsciiArt(text, banner)
	if err != 0 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data := struct {
		Art string
	}{
		Art: art,
	}
	// excute the template with the generated ascii art
	templates.ExecuteTemplate(w, "result.html", data)
}
func generateAsciiArt(text, bannerType string) (string, int) {
	bannerFilePath := filepath.Join("banners", bannerType+".txt")
	er := 0
	file, err := os.ReadFile(bannerFilePath)
	if err != nil {
		er = 1
		return "", er
	}
	// count the number of new lines in the entered argument
	count := strings.Count(text, "\r\n")
	// split the entred argument by new line
	testLines := strings.Split(text, "\r\n")
	var asciiChars []string
	// split the banner file by new line in case of thinkertoy banner
	if bannerType == "thinkertoy" {
		asciiChars = strings.Split(string(file), "\r\n\r\n")
	} else {
		asciiChars = strings.Split(string(file), "\n\n")
	}
	// store the outputed ascii art in 2D array format
	characters := make([][]string, len(asciiChars))
	for i, char := range asciiChars {
		if bannerType == "thinkertoy" {
			characters[i] = strings.Split(char, "\r\n")
		} else {
			characters[i] = strings.Split(char, "\n")
		}
	}
	// check if the entered text contains any non-printable characters
	for _, v := range text {
		if (v < 32 || v > 126) && (v != 10 && v != 13) {
			er = 1
			return "", er
		}
	}
	// print a result with all requirements
	result := ""
	counter := 1
	for _, line := range testLines {
		if line == "" {
			if counter <= count {
				result += "\r\n"
			}
			counter++
			continue
		}

		for l := 0; l < 8; l++ {
			for _, char := range line {
				if char == ' ' { // a space character
					result += characters[char-32][l+1]
				} else {
					// display the ascii art character line by line and column by column
					index := char - 32
					result += characters[index][l]
				}
			}
			result += "\r\n"
		}
	}
	// returen the result and error if any error occurred
	return result, er
}
