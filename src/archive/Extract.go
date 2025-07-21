package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/errors"
	fsys "github.com/jgttech/repo/src/fs"
)

func Extract(options ...archiveOption) (err error) {
	archive := &Archive{}

	for _, fn := range options {
		fn(archive)
	}

	if archive.From == "" {
		err = errors.New("Missing archive.From value.")
		return
	}

	if archive.To == "" {
		err = errors.New("Missing archive.To value.")
		return
	}

	cwd := assert.Must(os.Getwd())
	from := fsys.NewNode(archive.From)
	to := fsys.NewNode(archive.To)

	os.Chdir(to.Path)

	file := assert.Must(os.Open(from.Path))
	defer file.Close()

	_gz := assert.Must(gzip.NewReader(file))
	defer _gz.Close()

	_tar := tar.NewReader(_gz)

	for true {
		header, headerErr := _tar.Next()
		err = headerErr

		if err == io.EOF {
			break
		}

		if err != nil {
			err = errors.New("Failed to extract data.")
			os.Chdir(cwd)
			return
		}

		name := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(name, 0755); err != nil {
				err = errors.Errorf("Failed to extract directory:\n%s", name)
			}
		case tar.TypeReg:
			asset, err := os.Create(name)

			if err != nil {
				err = errors.Errorf("Archive extraction failed to extract file:\n%s", name)
			}

			if _, err := io.Copy(asset, _tar); err != nil {
				err = errors.Errorf("Archive extraction failed to copy file:\n%s", name)
			}

			asset.Close()
		default:
			err = errors.New("Unknown error, failed to extract data")
		}
	}

	os.Chdir(cwd)
	return
}
