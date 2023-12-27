package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: wget <URL>")
		os.Exit(1)
	}

	url := flag.Arg(0)
	err := downloadFile(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading file: %v\n", err)
		os.Exit(1)
	}
}

func downloadFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Получение имени файла из URL
	fileName := path.Base(url)
	if fileName == "/" || fileName == "." {
		fileName = "index.html"
	}

	// Создание файла для записи
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Копирование данных из ответа в файл
	_, err = io.Copy(file, resp.Body)
	return err
}
