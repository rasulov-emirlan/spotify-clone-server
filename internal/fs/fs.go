package fs

import "io"

type FileSystem interface {
	UploadFile(name string, mimeType string, content io.Reader, folderName string) (string, error)
	DeleteFile(filename string) error
}
