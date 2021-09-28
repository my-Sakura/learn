package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Tar
func Tar(src, dst string) error {
	if isDir(src) {
		return tarDir(src, dst)
	}

	return tarFile(src, dst)
}

// UnTar
func UnTar(src string) error {
	oldFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if err = oldFile.Close(); err != nil {
			panic(err)
		}
	}()

	gr, err := gzip.NewReader(oldFile)
	if err != nil {
		if err == gzip.ErrHeader {
			// Untar file
			fw, err := os.Open(src)
			if err != nil {
				return err
			}
			defer func() {
				if err = fw.Close(); err != nil {
					panic(err)
				}
			}()
			tr := tar.NewReader(fw)
			for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
				if err != nil {
					return err
				}

				switch hdr.Typeflag {
				case tar.TypeDir:
					if err := os.MkdirAll(hdr.Name, hdr.FileInfo().Mode()); err != nil {
						return err
					}

				case tar.TypeReg:
					file, err := os.OpenFile(hdr.Name, os.O_CREATE|os.O_RDWR, hdr.FileInfo().Mode())
					if err != nil {
						return err
					}
					_, err = io.Copy(file, tr)
					if err != nil {
						return err
					}

					file.Close()
				}
			}
			return nil
		} else {
			return err
		}
	}
	defer func() {
		if err := gr.Close(); err != nil {
			panic(err)
		}
	}()
	tr := tar.NewReader(gr)

	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			return err
		}

		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(hdr.Name, hdr.FileInfo().Mode()); err != nil {
				return err
			}

		case tar.TypeReg:
			file, err := os.OpenFile(hdr.Name, os.O_CREATE|os.O_RDWR, hdr.FileInfo().Mode())
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}

			file.Close()
		}
	}

	return nil
}

func tarFile(src, dst string) error {
	// Get oldFile header
	oldFileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	header, err := tar.FileInfoHeader(oldFileInfo, "")
	if err != nil {
		return err
	}

	// New tar writer
	newFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if err = newFile.Close(); err != nil {
			panic(err)
		}
	}()
	tw := tar.NewWriter(newFile)
	defer func() {
		if err = tw.Close(); err != nil {
			panic(err)
		}
	}()

	// write header
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// write body
	oldFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if err = oldFile.Close(); err != nil {
			panic(err)
		}
	}()
	_, err = io.Copy(tw, oldFile)
	if err != nil {
		return err
	}

	return nil
}

func tarDir(src, dst string) error {
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}

		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))

		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fr.Close()

		_, err = io.Copy(tw, fr)
		if err != nil {
			return err
		}

		log.Printf("succeed tar %s \n", fileName)

		return nil
	})
}
