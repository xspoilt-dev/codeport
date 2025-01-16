package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"encoding/json"
)

const (
	apiKey  = "142bfe7a264c3e621ecd6e3a7b1cba3d"
	baseURL = "https://pasteit.ftoolz.xyz"
)
func info() {
    fmt.Println("\033[1;36mCodePort(1.0.0) - A CLI Tool for Sharing Code\033[0m")
    fmt.Println("\033[1;33mDeveloped by: \033[1;35mMinhajul Islam \033[1;33m& \033[1;35mFarhan Ali\033[0m")
    fmt.Println("\033[1;32mWeb: \033[1;34mhttps://pasteit.ftoolz.xyz/\033[0m")
    fmt.Println("-----------------------------------------------")
}

func uploadCode(filePath, language, title, description, password string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("\033[1;31mFailed to read file: %w\033[0m", err)
	}

	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		URL     string `json:"url"`
	}

	encodedContent := url.QueryEscape(string(content))
	encodedLanguage := url.QueryEscape(language)
	encodedTitle := url.QueryEscape(title)
	encodedDescription := url.QueryEscape(description)
	encodedPassword := url.QueryEscape(password)

	apiURL := fmt.Sprintf("%s/api.php?api_key=%s&content=%s&syntax=%s&title=%s&description=%s&password=%s",
		baseURL, apiKey, encodedContent, encodedLanguage, encodedTitle, encodedDescription, encodedPassword)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("\033[1;31mFailed to upload code: %w\033[0m", err)
	}
	defer resp.Body.Close()

	var responseBody strings.Builder
	_, err = io.Copy(&responseBody, resp.Body)
	if err != nil {
		return "", fmt.Errorf("\033[1;31mFailed to read response: %w\033[0m", err)
	}

	var response Response
	err = json.Unmarshal([]byte(responseBody.String()), &response)
	if err != nil {
		return "", fmt.Errorf("\033[1;31mFailed to parse JSON response: %w\033[0m", err)
	}


	if response.Status != "success" {
		return "", fmt.Errorf("\033[1;31mupload failed: %s\033[0m", response.Message)
	}

	id := response.URL[strings.LastIndex(response.URL, "/")+1:]
	return id, nil
}

func downloadCode(id, outputFile string) error {
	pasteURL := fmt.Sprintf("%s/paste/%s", baseURL, id)
	resp, err := http.Get(pasteURL)
	if err != nil {
		return fmt.Errorf("\033[1;31mFailed to fetch paste page: %w\033[0m", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("\033[1;31mFailed to read HTML response: %w\033[0m", err)
	}
	html := string(body)

	downloadPattern := regexp.MustCompile(`href='([^"]+)'`)
	match := downloadPattern.FindStringSubmatch(html)
	if len(match) < 2 {
		return fmt.Errorf("\033[1;31mFailed to extract download URL from HTML\033[0m")
	}
	downloadURL := match[1]
	resp, err = http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("\033[1;31mFailed to download code file: %w\033[0m", err)
	}
	defer resp.Body.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("\033[1;31mFailed to create output file: %w\033[0m", err)
	}
	defer output.Close()

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		return fmt.Errorf("\033[1;31mFailed to save content to file: %w\033[0m", err)
	}

	fmt.Printf("\033[1;32mCode successfully saved: \033[1;34m%s\033[0m\n", outputFile)
	return nil
}

func main() {
	filePath := flag.String("f", "", "\033[1;33mPath to the file to upload\033[0m")
	downloadID := flag.String("g", "", "\033[1;33mPaste ID to download the file\033[0m")
	outputFile := flag.String("o", "output.txt", "\033[1;33mPath to save the downloaded file\033[0m")
	password := flag.String("p", "", "\033[1;33mPassword to protect the paste (optional)\033[0m")
	description := flag.String("d", "No description", "\033[1;33mDescription of the paste (optional)\033[0m")
	title := flag.String("t", "Untitled", "\033[1;33mTitle for the paste (optional)\033[0m")
	language := flag.String("l", "python", "\033[1;33mProgramming language for syntax highlighting (default: python)\033[0m")

	flag.Parse()
	info()
	if *filePath != "" && *downloadID == "" {
		fmt.Println("\033[1;32mUploading file to the server...\033[0m")
		pasteID, err := uploadCode(*filePath, *language, *title, *description, *password)
		if err != nil {
			fmt.Printf("\033[1;31mError during upload: %s\033[0m\n", err)
			return
		}
		fmt.Printf("\033[1;32mYour code has been successfully uploaded\033[0m\n\033[1;33mShare-Id: %s\033[0m\n\033[1;34mYou can get the code using 'codeport -g %s'\033[0m\n", pasteID, pasteID)
	} else if *filePath == "" && *downloadID != "" {
		fmt.Println("\033[1;32mDownloading file from the server...\033[0m")
		err := downloadCode(*downloadID, *outputFile)
		if err != nil {
			fmt.Printf("\033[1;31mError during download: %s\033[0m\n", err)
		}
	} else {
		fmt.Println("\033[1;31mInvalid usage. Use -f to upload a file or -g to download using an ID.\033[0m")
		flag.Usage()
	}
}
