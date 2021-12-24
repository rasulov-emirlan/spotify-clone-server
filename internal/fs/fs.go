package fs

import "io"

type FileSystem interface {
	UploadFile(name string, mimeType string, content io.Reader, folderName string) error
	DeleteFile(filename string) error
}
