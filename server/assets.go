package server

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type assetsFileSystem struct {
	box *rice.Box
}

func newAssetsFileSystem() *assetsFileSystem {
	return &assetsFileSystem{
		box: rice.MustFindBox("../ui/build"),
	}
}

func (fs *assetsFileSystem) Open(name string) (http.File, error) {
	f, err := fs.box.Open(name)

	if os.IsNotExist(err) {
		return fs.box.Open("index.html")
	}
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (s *Server) assets(c echo.Context) error {
	return echo.WrapHandler(http.FileServer(newAssetsFileSystem()))(c)
}
