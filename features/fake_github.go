package features

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"syscall"

	"github.com/tedsuo/ifrit/http_server"

	"github.com/google/go-github/v25/github"
	"github.com/gorilla/mux"
	"github.com/tedsuo/ifrit"
)

func createString(x string) *string { return &x }
func createBool(x bool) *bool       { return &x }
func createInt(x int64) *int64      { return &x }

type FakeGitHub struct {
	Host          string
	Server        ifrit.Runner
	serverProcess ifrit.Process

	ReleaseAssets   []*github.ReleaseAsset
	Releases        []*github.RepositoryRelease
	ReleasesPerPage int
}

func (m *FakeGitHub) logAndHandle(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("GitHub Request: [%s] %s\n", r.Method, r.URL.String())
		handler(w, r)
	}
}

func (m *FakeGitHub) Start() {
	m.serverProcess = ifrit.Invoke(m.Server)
}

func (m *FakeGitHub) Stop() {
	m.serverProcess.Signal(syscall.SIGKILL)
}

func (m *FakeGitHub) makeFakeNeedsReleaseAsset(version string, platform string) *github.ReleaseAsset {
	asset := github.ReleaseAsset{
		ID:                 createInt(int64(len(m.ReleaseAssets) + 1)),
		URL:                createString(fmt.Sprintf("https://localhost:9877/repos/cf-platform-eng/needs/releases/assets/%s", version)),
		Name:               createString(fmt.Sprintf("%s-%s", version, platform)),
		BrowserDownloadURL: createString(fmt.Sprintf("https://localhost:9877/cf-platform-eng/needs/releases/download/%s/needs-%s-%s", version, version, platform)),
	}
	m.ReleaseAssets = append(m.ReleaseAssets, &asset)

	return &asset
}

func (m *FakeGitHub) MakeFakeNeedsRelease(version string, prerelease bool) *github.RepositoryRelease {
	assets := []github.ReleaseAsset{
		*m.makeFakeNeedsReleaseAsset(version, "linux"),
		*m.makeFakeNeedsReleaseAsset(version, "darwin"),
	}
	release := &github.RepositoryRelease{
		Prerelease: createBool(prerelease),
		URL:        createString(fmt.Sprintf("https://localhost:9877/repos/cf-platform-eng/needs/releases/%s", version)),
		TagName:    createString(version),
		AssetsURL:  createString(fmt.Sprintf("https://localhost:9877/repos/cf-platform-eng/needs/releases/%s/assets", version)),
		Assets:     assets,
		Name:       createString(version),
	}

	return release
}

func makePaginatedUrl(source *url.URL, page float64) *url.URL {
	result, _ := url.Parse(source.String())
	queryParams := result.Query()
	queryParams.Set("page", fmt.Sprintf("%d", int(page)))
	result.RawQuery = queryParams.Encode()
	return result
}

func NewFakeGitHub() *FakeGitHub {
	fake := &FakeGitHub{
		ReleaseAssets:   []*github.ReleaseAsset{},
		Releases:        []*github.RepositoryRelease{},
		ReleasesPerPage: 3,
	}

	fake.Releases = append(
		fake.Releases,
		fake.MakeFakeNeedsRelease("2.0.0", false),
		fake.MakeFakeNeedsRelease("2.0.0-rc1", true),
		fake.MakeFakeNeedsRelease("1.2.3", false),
		fake.MakeFakeNeedsRelease("1.0.0", false),
	)

	router := mux.NewRouter()
	router.HandleFunc("/repos/cf-platform-eng/needs/releases", func(w http.ResponseWriter, r *http.Request) {
		var err error
		page := 1
		if r.URL.Query().Get("page") != "" {
			page, err = strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				http.Error(w, "strconv.Atoi failure", http.StatusInternalServerError)
			}
		}
		w.Header().Set("Content-Type", "application/json")

		firstLink := makePaginatedUrl(r.URL, 1)
		prevLink := makePaginatedUrl(r.URL, math.Max(1, float64(page-1)))

		lastPage := math.Ceil(float64(len(fake.Releases)) / float64(fake.ReleasesPerPage))
		nextLink := makePaginatedUrl(r.URL, math.Min(float64(page+1), lastPage))
		lastLink := makePaginatedUrl(r.URL, lastPage)

		w.Header().Set("Link", fmt.Sprintf("<%s>; rel=\"first\", <%s>; rel=\"prev\", <%s>; rel=\"next\", <%s>; rel=\"last\"",
			firstLink.String(),
			prevLink.String(),
			nextLink.String(),
			lastLink.String(),
		))

		startIndex := (page - 1) * fake.ReleasesPerPage
		endIndex := int(math.Min(float64(page*fake.ReleasesPerPage), float64(len(fake.Releases))))
		data, err := json.Marshal(fake.Releases[startIndex:endIndex])
		if err != nil {
			http.Error(w, "marshal failure", http.StatusInternalServerError)
		}

		_, err = w.Write(data)
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/repos/cf-platform-eng/needs/releases/assets/{asset_id}", func(w http.ResponseWriter, r *http.Request) {
		assetID, err := strconv.Atoi(mux.Vars(r)["asset_id"])
		if err != nil {
			http.Error(w, "strconv.Atoi failure", http.StatusInternalServerError)
		}
		asset := fake.ReleaseAssets[assetID-1]

		w.Header().Set("Content-Type", "application/octet-stream")

		_, err = fmt.Fprintf(w, "needs %s content", *asset.Name)
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	fake.Host = "http://localhost:9877/"
	fake.Server = http_server.New("localhost:9877", router)

	return fake
}
