package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/gif"
	"image/png"
	"os"
	"strings"
)

func (p *HttpResponse) createStringBody(filename string) string {
	str := ""
	file, err := os.Open(filename)
	if err != nil {
		// Openエラー
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		str = str + sc.Text() + "\n"
	}

	return str

}

func (p *HttpResponse) createBinaryBody(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		// Openエラー
		fmt.Println("ファイルが開けません")
	}

	if strings.Contains(filename, ".png") {
		img, err := png.Decode(file)
		if err != nil {
			fmt.Println(err)
		}

		buffer := new(bytes.Buffer)
		if err := png.Encode(buffer, img); err != nil {
			fmt.Println(err)
		}
		imageBytes := buffer.Bytes()
		return string(imageBytes)
	} else if strings.Contains(filename, ".gif") {
		img, err := gif.Decode(file)
		if err != nil {
			fmt.Println(err)
		}

		buffer := new(bytes.Buffer)
		if err := gif.Encode(buffer, img, nil); err != nil {
			fmt.Println(err)
		}

		imageBytes := buffer.Bytes()
		return string(imageBytes)
	}

	return "hoge"

}
