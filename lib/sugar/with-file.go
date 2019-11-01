package sugar

import (
	"log"
	"os"
)


func closes(files ...*os.File) {
	for i := range files {
		err := files[i].Close()
		if err != nil {
			log.Printf("close file %v error, err: %v\n", i, err)
		}
	}
}

func Files(bindFiles ...*os.File) []*os.File {
	return bindFiles
}

func WithFiles(f func(), bindFiles []*os.File, filePaths... string) {
	var files []*os.File

	for i := range filePaths {
		file, err := os.Create(filePaths[i])
		if err != nil {
			closes(files...)
			log.Printf("open file %v error, err: %v\n", filePaths[i], err)
		}
		files = append(files, file)
		*bindFiles[i] = *file
	}
	defer func() {
		if err := recover(); err != nil {
			PrintStack()
			log.Println("panic", "error", err)
		}
		closes(files...)
	}()

	f()
}

func WithFile(f func(file *os.File), filePath string) {
	WithOpenFile(f, filePath, os.O_RDWR|os.O_CREATE, 0755)
}

func WithWriteFile(f func(file *os.File), filePath string) {
	WithOpenFile(f, filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
}

func WithOpenFile(f func(file *os.File), filePath string, fileMode int, permMode os.FileMode) {
	file, err := os.OpenFile(filePath, fileMode, permMode)
	if err != nil {
		log.Printf("open file %v error, err: %v\n", filePath, err)
	}
	defer func() {
		if err := recover(); err != nil {
			PrintStack()
			log.Println("panic", "error", err)
		}
		closes(file)
	}()

	f(file)
}
