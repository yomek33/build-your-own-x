package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	mode := os.Args[1]
	fileName := os.Args[2]

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}

	switch mode {
	case "1":
		catOsReadFile(fileName)
	}
}


func catOsReadFile(fileName string){
	content, err := ReadFile(fileName)
	if err != nil {
		fmt.Println("Cannnot read file", err)
	}
	println(string(content))
}

// dive into the os/file
// ReadFile
// https://github.com/golang/go/blob/master/src/os/file.go#L788
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot open file", err)
		return nil, err
	}
	defer f.Close()

	var size int
	if info, err := f.Stat(); err == nil {
		size64 := info.Size()
		if int64(int(size64)) == size64 {
			size = int(size64)
		}
	}

	size ++ // 1 byte for final read at EOF

	if size <512 {
		size = 512
	}

	data := make([]byte, 0, size) 
	for {
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}

		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
	}
}


