package client;

// PackageInfo represents the overall structure of the package data.
type PackageInfo struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	RepoURL       string        `json:"repo_url"`
	License       string        `json:"license"`
	LatestVersion LatestVersion `json:"latest_version"`
}

type LatestVersion struct {
    GitTag      string `json:"git_tag"`
    PublishedAt string `json:"published_at"`
    Release     int    `json:"release"`
    SourceURL   string `json:"source_url"`
    Version     string `json:"version"`
}

type Package struct {
    CreatedBy     int           `json:"created_by"`
    Description   string        `json:"description"`
    Homepage      string        `json:"homepage"`
    ID            int           `json:"id"`
    LatestVersion LatestVersion `json:"latest_version"`
    License       string        `json:"license"`
    Name          string        `json:"name"`
    RepoURL       string        `json:"repo_url"`
}

