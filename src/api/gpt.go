package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func GPTHandler(prompt string) (string, error) {
	url := "https://api.chatanywhere.tech/v1/chat/completions"
	fmt.Printf("uuuuuuuuuuuuuuu%s\n", prompt)

	jsonData := []byte(fmt.Sprintf(`{
		"model": "gpt-3.5-turbo",
		"messages": [
			{
				"role": "system",
				"content": "你是怪物猎人里的随从'呆猫',如果没有特别指定,用中文回答问题,但是在需要的时候也可以用其他语言回答问题"
			},
			{
				"role": "user",
				"content": "%s"
			}
		],
		"temperature": 0.7
	}`, prompt))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("create req failed: %w", err)
	}

	req.Header.Set("Authorization", "Bearer sk-10oKFBWQhDAy75xzCWSAnHFU7mOfQPTK0w9B5WEnLW4s8hYx")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send req failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var gptResp GPTResponse
	if err := json.Unmarshal(body, &gptResp); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	if len(gptResp.Choices) > 0 {
		return gptResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no content in response")
}
