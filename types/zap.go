package types

import (
	"errors"
	"path"
)

type ZapDlAsset struct {
	Name     string `json:"name"`
	Download string `json:"download"`
	Size     string `size:"size"`
}

func (asset ZapDlAsset) GetBaseName() string {
	return path.Base(asset.Name)
}

// ZapIndex is used to marshal the index.min.json
type ZapIndex struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Image      string      `json:"image"`
	Maintainer string      `json:"maintainer"`
	Summary    string      `json:"summary"`
	Links      []ZapSource `json:"links"`
}

type ZapRelease struct {
	Roll        int
	Id          string                `json:"id"`
	Author      string                `json:"author"`
	PreRelease  bool                  `json:"prerelease"`
	Releases    string                `json:"releases"`
	Assets      map[string]ZapDlAsset `json:"assets"`
	Tag         string                `json:"tag"`
	PublishedAt string                `json:"published_at"`
}

type ZapSource struct {
	Type string
	Url  string
}

type ZapReleases struct {
	Releases map[int]ZapRelease
	Author   string
	Source   ZapSource
}

func (r ZapReleases) GetReleasesArray() []string {
	arr := make([]string, 0, len(r.Releases))
	// keys := make([]int, 0, len(r.Releases))
	// logger.Debug(r.Releases)
	for i := 0; i < len(r.Releases); i++ {
		arr = append(arr, r.Releases[i].Tag)
	}

	return arr
}

func (r ZapReleases) SortByReleaseDate(i, j int) bool {
	return r.Releases[i].Roll < r.Releases[j].Roll
}

func (r ZapReleases) GetLatestRelease() string {
	return r.Releases[0].Tag
}

func (r ZapReleases) GetAssetsFromTag(tag string) (map[string]ZapDlAsset, error) {
	for i := range r.Releases {
		if r.Releases[i].Tag == tag {
			return r.Releases[i].Assets, nil
		}
	}
	return map[string]ZapDlAsset{}, errors.New("could not find tag in release")
}
