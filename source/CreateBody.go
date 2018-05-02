package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

// HttpResponse.bodyにコンテンツを追加する（ファイルを読み込む）
func (p *HttpResponse) addBodyPartsFromFile(path string) {
	filename := "../public/"

	switch path {
	case "/hello":
		filename = filename + "hello.html"
		p.body = p.createStringBody(filename)
	case "/":
		filename = filename + "index.html"
		p.body = p.createStringBody(filename)
	case "/eevee":
		filename = filename + "eevee.html"
		p.body = p.createStringBody(filename)
	case "/methodNotAllowed":
		filename = filename + "methodNotAllowed.html"
		p.body = p.createStringBody(filename)
	default:
		filename = filename + strings.Trim(path, "/")
		fmt.Println(filename)
		if _, err := os.Stat(filename); err != nil {
			// not found
			filename = "../public/notFound.html"
			p.body = p.createStringBody(filename)
			p.status = "404 Not Found"
		} else {
			if strings.Contains(filename, ".png") || strings.Contains(filename, ".gif") {
				p.body = p.createBinaryBody(filename)
			} else {
				p.body = p.createStringBody(filename)
			}
		}
	}

	if p.body == "" {
		filename = "../public/internalServerError.html"
		p.body = p.createStringBody(filename)
		p.status = "500 Internal Server Error"
	}

	// bodyからcontentlengthを計算してheaderに追加
	p.addHeaderParts("Content-Length", makeContentLength(p.body))

	// filenameからcontenttypeを設定してheaderに追加
	p.addHeaderParts("Content-Type", makeContentType(filename))
}

func (p *HttpResponse) createStringBody(filename string) string {
	str := ""
	file, err := os.Open(filename)
	if err != nil {
		// Openエラー
		defer func() {
			if err := recover(); err != nil {
				log.Printf("ファイルが開けません: %v", err)
			}
		}()
		panic("Occured panic!")
	} else {
		sc := bufio.NewScanner(file)
		for sc.Scan() {
			if err := sc.Err(); err != nil {
				break
			}
			str = str + sc.Text() + "\n"
		}
	}
	defer file.Close()
	return str

}

func (p *HttpResponse) createBinaryBody(filename string) string {
	str := ""
	file, err := os.Open(filename)
	if err != nil {
		// Openエラー
		defer func() {
			if err := recover(); err != nil {
				log.Printf("ファイルが開けません: %v", err)
			}
		}()
		panic("Occured panic!")
	}

	if strings.Contains(filename, ".png") {
		img, err := png.Decode(file)
		if err != nil {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("pngファイルをデコードできません: %v", err)
				}
			}()
			panic("Occured panic!")
		}

		buffer := new(bytes.Buffer)
		if err := png.Encode(buffer, img); err != nil {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("pngファイルをエンコードできません: %v", err)
				}
			}()
			panic("Occured panic!")
		}

		imageBytes := buffer.Bytes()
		str = string(imageBytes)

	} else if strings.Contains(filename, ".gif") {
		imgs, err := gif.DecodeAll(file)
		if err != nil {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("gifファイルをデコードできません: %v", err)
				}
			}()
			panic("Occured panic!")
		}

		buffer := new(bytes.Buffer)
		if err := gif.EncodeAll(buffer, imgs); err != nil {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("gifファイルをエンコードできません: %v", err)
				}
			}()
			panic("Occured panic!")
		}

		imageBytes := buffer.Bytes()
		str = string(imageBytes)
	} else if strings.Contains(filename, ".jpg") {
		img, err := jpeg.Decode(file)
		if err != nil {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("jpgファイルをデコードできません: %v", err)
				}
			}()
			panic("Occured panic!")
		}

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, img, nil); err != nil {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("jpgファイルをエンコードできません: %v", err)
				}
			}()
			panic("Occured panic!")
		}

		imageBytes := buffer.Bytes()
		str = string(imageBytes)
	}

	return str

}
