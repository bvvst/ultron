package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func QueryChatGPT(query string) string {
	url := "https://chat.openai.com/backend-api/conversation"

	queryString := fmt.Sprintf("{\n   \"action\":\"next\",\n   \"messages\":[\n      {\n         \"role\":\"user\",\n         \"content\":{\n            \"content_type\":\"text\",\n            \"parts\":[\n               \"%s\"\n            ]\n         }\n      }\n   ],\n   \"parent_message_id\":\"\",\n   \"model\":\"text-davinci-002-render\"\n}", query)
	payload := strings.NewReader(queryString)

	// req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer INSERT BEARER TOKEN HERE")
	// Make a request to the URL
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, res.Body); err != nil {
		return "An error occured."
	}

	fmt.Println(string(buf.Bytes()))

	re := regexp.MustCompile(`\{[^\{\}]*(?:\{[^\{\}]*\}[^\{\}]*)*\}`)
	groups := re.FindAllString(string(buf.Bytes()), -1)
	fmt.Println(groups[len(groups)-2])

	jsonData := []byte(groups[len(groups)-2])

	var resp GPTResp
	err = json.Unmarshal(jsonData, &resp)
	if err != nil {
		fmt.Println(err)
		return "An error occured."
	}

	return resp.Content.Parts[0]
}

type GPTResp struct {
	ID         string      `json:"id"`
	Role       string      `json:"role"`
	User       interface{} `json:"user"`
	CreateTime interface{} `json:"create_time"`
	UpdateTime interface{} `json:"update_time"`
	Content    struct {
		ContentType string   `json:"content_type"`
		Parts       []string `json:"parts"`
	} `json:"content"`
	EndTurn  interface{} `json:"end_turn"`
	Weight   float64     `json:"weight"`
	Metadata struct {
	} `json:"metadata"`
	Recipient string `json:"recipient"`
}

type ChatGPTReq struct {
	Action          string    `json:"action"`
	Messages        []Message `json:"messages"`
	ParentMessageID string    `json:"parent_message_id"`
	Model           string    `json:"model"`
}

type Message struct {
	ID      string `json:"id"`
	Role    string `json:"role"`
	Content Content
}

type Content struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}
