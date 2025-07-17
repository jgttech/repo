package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/jgttech/repo/src/assert"
	// "github.com/jgttech/repo/src/pathlib"
)

func (self *Archive) Write() (err error) {
	self.Files = []string{
		"/home/user/docs/file1.txt",
		"/home/admin/user/file2.txt",
	}

	_, err = os.Stat(self.Path)

	if err == nil {
		err = fmt.Errorf("Archive already exists: %s\n", self.Path)
		return
	}

	file, err := os.Create(self.Path)

	if file == nil {
		err = fmt.Errorf("File not found: %s\n", self.Path)
		return
	}

	defer file.Close()

	gzWriter := gzip.NewWriter(file)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// base := pathlib.SharedDir(self.Files)

	for _, filename := range self.Files {
		target := assert.Must(os.Open(filename))
		info := assert.Must(target.Stat())

		header := assert.Must(tar.FileInfoHeader(info, info.Name()))

		// Use full path as name (FileInfoHeader only takes the basename)
		// If we don't do this the directory structure would
		// not be preserved.
		//
		// https://golang.org/src/archive/tar/common.go?#L626
		header.Name = filename
		fmt.Println(header.Name)

		assert.Throws(tarWriter.WriteHeader(header))
		assert.Must(io.Copy(tarWriter, target))
		target.Close()
	}

	return
}
