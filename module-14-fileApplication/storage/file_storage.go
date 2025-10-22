package storage

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// This will store metadata info which is mapped by file name
var MetadataMap = map[string]FileMetadata{}

func SaveFile(filename string, data []byte) error {
	// to extract the file name
	fName := filepath.Base(filename)

	err := os.WriteFile("storage/"+fName, data, 0777)
	if err != nil {
		log.Fatalf("Unable to save the file: ", err)
	}
	info, err := os.Stat("storage/" + fName)
	if err != nil {
		return err
	}

	MetadataMap[fName] = FileMetadata{
		Filename:   fName,
		Size:       info.Size(),
		UploadTime: time.Now().Format("2006-01-02 15:02:01"),
	}

	return nil
}

func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile("storage/" + filename)
}

func GetMetadata(filename string) ([]FileMetadata, error) {
	if filename == "" {
		values := []FileMetadata{}
		for _, v := range MetadataMap {
			values = append(values, v)
		}
		return values, nil
	}
	if md, ok := MetadataMap[filename]; ok {
		return []FileMetadata{md}, nil
	}
	return nil, os.ErrNotExist
}

type FileMetadata struct {
	Filename   string
	Size       int64
	UploadTime string
}
