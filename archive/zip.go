package archive

import (
	"archive/zip"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Zip(src, dst string) error {
	if isDir(src) {
		return zipDir(src, dst)
	}

	return zipFile(src, dst)
}

func UnZip(src, dst string) error {
	zr, err := zip.OpenReader(src)
	if err != nil {
		if err == zip.ErrFormat {
			return errors.New("Unsupport unzip dir now")
		}

		return err
	}
	defer func() {
		if err = zr.Close(); err != nil {
			panic(err)
		}
	}()

	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}

		fr, err := file.Open()
		if err != nil {

			return err
		}
		defer func() {
			if err = fr.Close(); err != nil {
				panic(err)
			}
		}()
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err = fw.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
	}

	return nil
}

func zipFile(src, dst string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}
	hdr, err := zip.FileInfoHeader(fi)
	if err != nil {
		return err
	}

	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if err = fw.Close(); err != nil {
			panic(err)
		}
	}()

	zw := zip.NewWriter(fw)
	defer func() {
		if err = zw.Close(); err != nil {
			panic(err)
		}
	}()

	w, err := zw.CreateHeader(hdr)
	if err != nil {
		return err
	}

	fr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if err = fr.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err = io.Copy(w, fr); err != nil {
		return err
	}

	return nil
}

func zipDir(src, dst string) error {
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if err = fw.Close(); err != nil {
			panic(err)
		}
	}()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := zip.NewWriter(gw)
	defer tw.Close()

	return filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		if fi.IsDir() {
			fh.Name += "/"
		}

		w, err := tw.CreateHeader(fh)
		if err != nil {
			return err
		}

		if !fh.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			if err = fr.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}

		return nil
	})
}
