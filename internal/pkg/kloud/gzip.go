package kloud

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const GZIP_FILE_NAME = "kloud_arch.tar.gz"

func Gzip(source, target string) error {
	tarPath := filepath.Join(target, "kloud_arch.tar")
	if err := tarAct(source, tarPath); err != nil {
		return err
	}

	if err := gzipAct(tarPath, target); err != nil {
		return err
	}

	os.Remove(tarPath)
	return nil
}

func gzipFilePath(target string) string {
	return filepath.Join(target, GZIP_FILE_NAME)
}

func tarAct(source, target string) error {
	tarFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tarball := tar.NewWriter(tarFile)
	defer tarball.Close()

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if source == path || path == target {
				return nil
			}

			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}
			header.Name = strings.TrimPrefix(path, source)
			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

func gzipAct(tarPath, target string) error {
	reader, err := os.Open(tarPath)
	if err != nil {
		return err
	}

	targetFile := gzipFilePath(target)
	writer, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer writer.Close()

	archive := gzip.NewWriter(writer)
	archive.Name = GZIP_FILE_NAME
	defer archive.Close()

	_, err = io.Copy(archive, reader)
	return err
}
