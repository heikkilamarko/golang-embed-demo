package utils

import (
	"io/fs"
	"net/http"
)

type SPAHandler struct {
	fsys       fs.FS
	indexHTML  []byte
	fileServer http.Handler
}

func NewSPAHandler(fsys fs.FS, dirPath string, indexPath string) (*SPAHandler, error) {
	if dirPath != "" {
		var err error
		if fsys, err = fs.Sub(fsys, dirPath); err != nil {
			return nil, err
		}
	}

	indexHTML, err := fs.ReadFile(fsys, indexPath)
	if err != nil {
		return nil, err
	}

	return &SPAHandler{
		fsys:       fsys,
		indexHTML:  indexHTML,
		fileServer: http.FileServer(http.FS(fsys)),
	}, nil
}

func (h *SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if file, err := fs.Stat(h.fsys, r.URL.Path); err != nil || file.IsDir() {
		w.WriteHeader(http.StatusOK)
		w.Write(h.indexHTML)
		return
	}
	h.fileServer.ServeHTTP(w, r)
}
