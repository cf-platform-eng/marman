package features

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	ReleaseFiles map[string][]pivnet.ProductFile
	Releases     []pivnet.Release
	ProductFiles []pivnet.ProductFile
}

type Handler func(w http.ResponseWriter, r *http.Request)

func (m *FakeTanzuNetwork) logAndHandle(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("TanzuNetwork Request: [%s] %s\n", r.Method, r.URL.String())
		handler(w, r)
	}
}

func (m *FakeTanzuNetwork) Start() {
	m.serverProcess = ifrit.Invoke(m.Server)
}

func (m *FakeTanzuNetwork) Stop() {
	if m.serverProcess != nil {
		m.serverProcess.Signal(syscall.SIGKILL)
	}
	m.serverProcess = nil
}

func NewFakeTanzuNetwork() *FakeTanzuNetwork {
	fake := &FakeTanzuNetwork{}

	router := mux.NewRouter()
	router.HandleFunc("/api/v2/products/{slug}/releases", func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc("/api/v2/products/{slug}/releases/{release_id}/product_files", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		response := pivnet.ProductFilesResponse{
			ProductFiles: fake.ReleaseFiles[vars["release_id"]],
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

	router.HandleFunc("/api/v2/products/{slug}/releases/{release_id}/pivnet_resource_eula_acceptance", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/api/v2/products/{slug}/releases/{release_id}/product_files/{file_id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		var chosenFile pivnet.ProductFile
		for _, file := range fake.ReleaseFiles[vars["release_id"]] {
			if strconv.Itoa(file.ID) == vars["file_id"] {
				chosenFile = file
				break
			}
		}

		response := pivnet.ProductFileResponse{
			ProductFile: chosenFile,
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

	router.HandleFunc("/api/v2/path/to/light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", fmt.Sprintf("%s%s", fake.Host, "/actual/path/to/light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz"))
		w.WriteHeader(http.StatusFound)
	})

	router.HandleFunc("/api/v2/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", fmt.Sprintf("%s%s", fake.Host, "/actual/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz"))
		w.WriteHeader(http.StatusFound)
	})

	router.HandleFunc("/api/v2/path/to/bosh-stemcell-100.000-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", fmt.Sprintf("%s%s", fake.Host, "/actual/path/to/bosh-stemcell-100.000-google-super-duper-stemcell.tgz"))
		w.WriteHeader(http.StatusFound)
	})

	router.HandleFunc("/actual/path/to/light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusPartialContent)
		_, err := fmt.Fprint(w, "light stemcell file contents")
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/actual/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusPartialContent)
		_, err := fmt.Fprint(w, "stemcell file contents")
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/actual/path/to/bosh-stemcell-100.000-google-super-duper-stemcell.tgz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusPartialContent)
		_, err := fmt.Fprint(w, "stemcell file contents")
		if err != nil {
			http.Error(w, "printf failure", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("[\"Request to an unknown endpoint: %s %s\"]", r.Method, r.URL.Path), http.StatusNotFound)
	})

	fake.Host = "http://localhost:9876"
	fake.Server = http_server.New("localhost:9876", router)
	return fake
}
