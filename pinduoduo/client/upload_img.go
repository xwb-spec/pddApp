package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type ImgUploadRequestParam struct {
	File bytes.Reader
	CommonRequestParam
}
type ImgUploadResponse struct {
	GoodsImgUploadResponse GoodsImgUploadResponse `json:"goods_img_upload_response"`
}

type GoodsImgUploadResponse struct {
	url string
}

func UploadImg(filename string, targetUrl string) (string, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return "", err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return "", err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "", err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return "", err
	}
	rspBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		var result ImgUploadResponse
		if err = json.Unmarshal(rspBody, &result); err != nil {
			log.Fatal("Unmarshal fail, err:%v", err)
			return "", err
		}
	} else {
		var result ErrorResult
		if err = json.Unmarshal(rspBody, &result); err != nil {
			log.Fatal("Unmarshal fail, err:%v", err)
			return "", err
		}
	}
	return "", nil
}
