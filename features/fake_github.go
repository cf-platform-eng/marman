package features

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syscall"

	"github.com/tedsuo/ifrit/http_server"

	"github.com/google/go-github/v25/github"
	"github.com/gorilla/mux"
	"github.com/tedsuo/ifrit"
)

type FakeGitHub struct {
	Host          string
	Server        ifrit.Runner
	serverProcess ifrit.Process
	Requests      []*http.Request
	RequestBodies []interface{}

	Releases []github.RepositoryRelease
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

func NewFakeGitHub() *FakeGitHub {
	fake := &FakeGitHub{}

	createString := func(x string) *string {
		return &x
	}

	createBool := func(x bool) *bool {
		return &x
	}

	createInt := func(x int64) *int64 {
		return &x
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/repos/cf-platform-eng/needs/releases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var firstAssets = []github.ReleaseAsset{
			{
				ID:                 createInt(1),
				URL:                createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/assets/1.2.3"),
				Name:               createString("1.2.3-linux"),
				BrowserDownloadURL: createString("https://localhost:9877/cf-platform-eng/needs/releases/download/1.2.3/needs-1.2.3-linux"),
			},
			{
				ID:                 createInt(2),
				URL:                createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/assets/1.2.3"),
				Name:               createString("1.2.3-darwin"),
				BrowserDownloadURL: createString("https://localhost:9877/cf-platform-eng/needs/releases/download/1.2.3/needs-1.2.3-darwin"),
			},
		}

		var secondAssets = []github.ReleaseAsset{
			{
				ID:                 createInt(3),
				URL:                createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/assets/2.0.0"),
				Name:               createString("2.0.0-linux"),
				BrowserDownloadURL: createString("https://localhost:9877/cf-platform-eng/needs/releases/download/2.0.0/needs-2.0.0-linux"),
			},
			{
				ID:                 createInt(4),
				URL:                createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/assets/2.0.0"),
				Name:               createString("2.0.0-darwin"),
				BrowserDownloadURL: createString("https://localhost:9877/cf-platform-eng/needs/releases/download/2.0.0/needs-2.0.0-darwin"),
			},
		}

		var response = [...]github.RepositoryRelease{
			{
				Prerelease: createBool(false),
				URL:        createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/2.0.0"),
				TagName:    createString("2.0.0"),
				AssetsURL:  createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/2.0.0/assets"),
				Assets:     secondAssets,
				Name:       createString("2.0.0"),
			},
			{
				Prerelease: createBool(false),
				URL:        createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/1.2.3"),
				TagName:    createString("1.2.3"),
				AssetsURL:  createString("https://localhost:9877/repos/cf-platform-eng/needs/releases/1.2.3/assets"),
				Assets:     firstAssets,
				Name:       createString("1.2.3"),
			},
		}

		data, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "marshal failure", http.StatusInternalServerError)
		}

		_, err = w.Write(data)
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/repos/cf-platform-eng/needs/releases/assets/1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")

		var response = "needs linux 1.2.3 asset content"

		_, err := w.Write([]byte(response))
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/repos/cf-platform-eng/needs/releases/assets/3", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")

		var response = "needs linux 2.0.0 asset content"

		_, err := w.Write([]byte(response))
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	fake.Host = "http://localhost:9877/"
	fake.Server = http_server.New("localhost:9877", mux)

	return fake
}
