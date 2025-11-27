package client

import (
	"net/http"
    "encoding/json"
	"github.com/dvip1/dpm/internal/tool"
	"io"
)


func (c *Client) GetPackages() (Package, error) {
	url := c.BaseURL + "/" + AppRoutes.Packages.Root
	var pkg Package

	req, err := http.NewRequest("GET", url, nil)

	if tool.Error(err, "Failed to create HTTP request") {
		return pkg, err
	}

	resp, err := c.httpClient.Do(req)
	if tool.Error(err, "Couldn't make HTTP request") {
		return pkg, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		tool.Warn("Server returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if tool.Error(err, "Couldn't read response body") {
		return pkg, err
	}

	if err := json.Unmarshal(body, &pkg); tool.Error(err, "Couldn't unmarshal JSON") {
		return pkg, err
	}

	tool.Info("Successfully fetched package data from %s", url)
	return pkg, nil
}

func (c *Client) GetPackageByName(name string) (Package, error) {
	url := c.BaseURL + "/" + AppRoutes.Packages.Root + "/" + name
	req, err := http.NewRequest("GET", url, nil)
	var pkg Package

	if tool.Error(err, "Failed to create HTTP request") {
		return pkg, err
	}

	resp, err := c.httpClient.Do(req)
	if tool.Error(err, "Couldn't make HTTP request") {
		return pkg, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		tool.Warn("Server returned status %d", resp.StatusCode)
		return pkg, err;
	}

	body, err := io.ReadAll(resp.Body)
	if tool.Error(err, "Couldn't read response body") {
		return pkg, err
	}

	if err := json.Unmarshal(body, &pkg); tool.Error(err, "Couldn't unmarshal JSON") {
		return pkg, err
	}

	tool.Info("Successfully fetched package data from %s", url)
	return pkg, nil
}
