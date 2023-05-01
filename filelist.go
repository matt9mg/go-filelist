package filelist

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

type FileList struct {
}

func NewFileList() *FileList {
	return &FileList{}
}

// ListFromLocation reads the named directory recursively and returns a list of files and their locations in an
// unsorted order.
// If an error occurs during the list generation, all files before the error are returned including the error itself
func (f *FileList) ListFromLocation(dir string) ([]string, error) {
	entry, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	var output []string

	fl, err := f.buildOS(entry, output, dir)

	if err != nil {
		return nil, err
	}

	return fl, nil
}

// ListFromFS reads the embed.FS directory recursively and returns a list of files and their locations in an
// unsorted order.
// If an error occurs during the list generation, all files before the error are returned including the error itself
func (f *FileList) ListFromFS(fileSystem embed.FS, startingDIR string) ([]string, error) {
	entry, err := fileSystem.ReadDir(startingDIR)

	if err != nil {
		return nil, err
	}

	var output []string

	fl, err := f.buildFS(entry, output, startingDIR, fileSystem)

	if err != nil {
		return nil, err
	}

	return fl, err
}

func (f *FileList) buildFS(entries []fs.DirEntry, out []string, path string, fileSystem embed.FS) ([]string, error) {
	for _, entry := range entries {
		subPath := filepath.Join(path, entry.Name())

		if entry.IsDir() == true {
			subEntry, err := fs.ReadDir(fileSystem, subPath)

			if err != nil {
				return out, err
			}

			out, err = f.buildFS(subEntry, out, subPath, fileSystem)

			if err != nil {
				return nil, err
			}
			continue
		}

		out = append(out, subPath)
	}

	return out, nil
}

func (f *FileList) buildOS(entries []fs.DirEntry, out []string, path string) ([]string, error) {
	for _, entry := range entries {
		subPath := filepath.Join(path, entry.Name())

		if entry.IsDir() == true {
			subEntry, err := os.ReadDir(subPath)

			if err != nil {
				return out, err
			}

			out, err = f.buildOS(subEntry, out, subPath)

			if err != nil {
				return out, err
			}
			continue
		}

		out = append(out, subPath)
	}

	return out, nil
}
