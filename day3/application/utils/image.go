package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadImage(userId int, r *http.Request) string {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Unable to parse form")
		return ""
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Unable to retrive file", err)
		return ""
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)

	filename := fmt.Sprintf("user_%d%s", userId, ext)
	savePath := filepath.Join("images", filename)

	outFile, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Unable to create file")
		return ""
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		fmt.Println("Unable to copy file")
		return ""
	}

	return filename
}

func GetImage(filename string, r *http.Request) string {
	host := r.Host
	imagePath := fmt.Sprintf("http://%s/images/%s", host, filename)
	return imagePath
}

func DeleteImage(filename string) error {
	filepath := filepath.Join("images", filename)
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		fmt.Printf("file %s does not exist", filename)
		return nil
	}

	err = os.Remove(filepath)
	if err != nil {
		return fmt.Errorf("err delete file: %v", err)
	}

	return nil
}
