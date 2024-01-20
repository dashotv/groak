package scraper

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Metube struct {
	URL string `json:"url"`
}

func NewMetube(url string) *Metube {
	return &Metube{
		URL: url,
	}
}

func (m *Metube) Download(name, url string) error {
	client := resty.New()
	resp, err := client.R().
		SetBody(&MetubeDownload{url, false, "best", "any", name}).
		Post(m.URL)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("request failed: %d: %s: %s", resp.StatusCode(), resp.Status(), resp.String())
	}
	return nil
}

type MetubeDownload struct {
	URL       string `json:"url"`
	AutoStart bool   `json:"auto_start"`
	Quality   string `json:"quality"`
	Format    string `json:"format"`
	Name      string `json:"custom_name_prefix"`
}

type MetubeResponse struct {
	Status string `json:"status"`
}
