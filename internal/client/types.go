package client;

type LatestVersion struct {
	Version     string `json:"version"`
	PublishedAt string `json:"published_at"`
}

// PackageInfo represents the overall structure of the package data.
type PackageInfo struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	RepoURL       string        `json:"repo_url"`
	License       string        `json:"license"`
	LatestVersion LatestVersion `json:"latest_version"`
}
