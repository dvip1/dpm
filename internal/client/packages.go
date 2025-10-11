package client

import "net/http"

func (c *Client) GetPackages() (*http.Response, error) {
	url := c.BaseURL + "/" + AppRoutes.Packages.Root
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}
func (c *Client) GetPackageByName(name string) (*http.Response, error) {
	url := c.BaseURL + "/" + AppRoutes.Packages.Root + "/" + name
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}
