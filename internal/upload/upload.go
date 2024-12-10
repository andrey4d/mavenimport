/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package upload

import (
	"bytes"
	"io"

	"mime/multipart"
	"os"
	"sync"

	"net/http"
	"strings"

	"log/slog"

	"github.com/andrey4d/mavenimport/internal/artifacts"
)

type Client struct {
	URL        string
	Token      string
	HTTPClient *http.Client
	log        slog.Logger
}

func NewClient(log slog.Logger, url, repository, token string) *Client {
	return &Client{
		URL:        url + "/service/rest/v1/components?repository=" + repository,
		Token:      token,
		HTTPClient: &http.Client{},
		log:        log,
	}
}

func (c *Client) UploadGoWG(a artifacts.Artifact, wg *sync.WaitGroup, ch chan error, index int) {
	c.log.Debug("UploadGoWG()", slog.Int("gorutine index", index))
	defer wg.Done()

	err := c.Upload(a)

	if err != nil {
		ch <- err
	}
}

func (c *Client) Upload(a artifacts.Artifact) error {

	pom, err := os.Open(a.Pom)
	if err != nil {
		return err
	}
	jar, err := os.Open(a.Package)
	if err != nil {
		return err
	}

	values := map[string]io.Reader{
		"maven2.asset1":           pom,
		"maven2.asset2":           jar,
		"maven2.generate-pom":     strings.NewReader("false"),
		"maven2.asset1.extension": strings.NewReader("pom"),
		"maven2.asset2.extension": strings.NewReader("jar"),
	}

	b, FormDataContentType, err := c.multipartBody(values)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.URL, b)
	if err != nil {
		c.log.Error("Upload()", slog.Any("error", err))
	}
	req.Header.Set("Content-Type", FormDataContentType)
	req.Header.Add("Authorization", "Basic "+c.Token)
	c.log.Debug("Upload()", slog.Any("request header", req.Header))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		c.log.Error("Upload()", slog.Any("error", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.log.Error("Upload()", slog.Any("error", err))
	}

	if resp.StatusCode == 200 || resp.StatusCode == 204 {
		c.log.Info("Upload()", slog.Any("artifact", a.Package), slog.Any("Status", resp.Status))
	} else {
		c.log.Error("Upload()", slog.Any("artifact", a.Package), slog.Any("Status", resp.Status))
	}

	c.log.Debug("Upload()", slog.Any("response header", resp.Header), slog.Any("response body", string(body)))

	return nil
}

func (c *Client) multipartBody(values map[string]io.Reader) (*bytes.Buffer, string, error) {
	var body bytes.Buffer

	w := multipart.NewWriter(&body)
	defer w.Close()

	for key, r := range values {

		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}

		// Add a file
		var err error
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, "", err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, "", err
			}
		}

		if _, err = io.Copy(fw, r); err != nil {
			return nil, "", err
		}
	}
	return &body, w.FormDataContentType(), nil
}
