/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package artifacts

import (
	"fmt"
	"io"

	"net/http"
	"net/url"
	"strings"

	"log/slog"
)

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

func NewClient(url, token string) *Client {

	return &Client{
		BaseURL:    url,
		Token:      token,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) Upload(a Artifact) error {

	form := url.Values{}
	form.Add("maven2.generate", "false")
	form.Add("maven2.asset1", fmt.Sprintf("@%s/%s", a.Path, a.Pom))
	form.Add("maven2.asset1.extension", "pom")
	form.Add("maven2.asset2", fmt.Sprintf("@$s/$s;type=application/java-archive", a.Path, a.Package))
	form.Add("maven2.asset2.extension", "jar")

	req, err := http.NewRequest("POST", c.BaseURL, strings.NewReader(form.Encode()))
	if err != nil {
		slog.Error("error", slog.Any("error", err))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		slog.Error("error", slog.Any("error", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error", slog.Any("error", err))
	}

	fmt.Println(string(body))
	return nil
}
