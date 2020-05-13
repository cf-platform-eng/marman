package features

import (
	"encoding/json"
	"fmt"
	"net/http"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/pivotal-cf/go-pivnet"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/http_server"
)

type FakeTanzuNetwork struct {
	Host          string
	Server        ifrit.Runner
	serverProcess ifrit.Process
	Requests      []*http.Request
	RequestBodies []interface{}

	Releases     []pivnet.Release
	ProductFiles []pivnet.ProductFile
}

func (m *FakeTanzuNetwork) Start() {
	m.serverProcess = ifrit.Invoke(m.Server)
}

func (m *FakeTanzuNetwork) Stop() {
	m.serverProcess.Signal(syscall.SIGKILL)
}

func NewFakeTanzuNetwork() *FakeTanzuNetwork {
	fake := &FakeTanzuNetwork{}

	mux := mux.NewRouter()
	mux.HandleFunc("/api/v2/products/{slug}/releases", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := pivnet.ReleasesResponse{
			Releases: fake.Releases,
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

	mux.HandleFunc("/api/v2/products/{slug}/releases/{release_id}/product_files", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := pivnet.ProductFilesResponse{
			ProductFiles: fake.ProductFiles,
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

	mux.HandleFunc("/api/v2/products/{slug}/releases/{release_id}/pivnet_resource_eula_acceptance", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/api/v2/products/{slug}/releases/{release_id}/product_files/{file_id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := pivnet.ProductFileResponse{
			ProductFile: fake.ProductFiles[0],
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

	mux.HandleFunc("/api/v2/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", fmt.Sprintf("%s%s", fake.Host, "/actual/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz"))
		w.WriteHeader(http.StatusFound)
	})

	mux.HandleFunc("/actual/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusPartialContent)
		_, err := fmt.Fprint(w, "stemcell file contents")
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("Request to an unknown endpoint: %s %s", r.Method, r.URL.Path), http.StatusNotFound)
	})

	fake.Host = "http://localhost:9876"
	fake.Server = http_server.New("localhost:9876", mux)
	return fake
}
