package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/jgttech/repo/src/assert"
	"github.com/jgttech/repo/src/errors"
	fsys "github.com/jgttech/repo/src/fs"
)

type archiveOption func(*Archive)

func Create(options ...archiveOption) (err error) {
	archive := &Archive{}

	for _, fn := range options {
		fn(archive)
	}

	from := fsys.NewNode(archive.From)
	to := fsys.NewNode(archive.To)

	if !from.Exists() {
		err = errors.New(fmt.Sprintf("Failed to create archive: %s", from.Path))
		return
	}

	if to.Exists() {
		err = errors.New(fmt.Sprintf("Archive already exists: %s", from.Path))
		return
	}

	out := assert.Must(os.Create(archive.To))
	defer out.Close()

	_gz := gzip.NewWriter(out)
	defer _gz.Close()

	_tar := tar.NewWriter(_gz)
	defer _tar.Close()

	base := archive.From

	filepath.Walk(base, func(file string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file != base {
			info, _ := os.Stat(file)

			if !info.IsDir() {
				target := assert.Must(os.Open(file))
				info := assert.Must(target.Stat())
				header := assert.Must(tar.FileInfoHeader(info, info.Name()))

				// Use full path as name (FileInfoHeader only takes the basename)
				// If we don't do this the directory structure would
				// not be preserved.
				//
				// https://golang.org/src/archive/tar/common.go?#L626
				header.Name = assert.Must(filepath.Rel(archive.From, file))

				assert.Throws(_tar.WriteHeader(header))
				assert.Must(io.Copy(_tar, target))
				target.Close()
			}
		}

		return nil
	})

	return
}
