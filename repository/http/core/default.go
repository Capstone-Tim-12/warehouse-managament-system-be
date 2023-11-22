package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fadilahonespot/library/rest"
)


type wrapper struct {
	client rest.RestClient
}

var address string

func NewWrapper() CoreWrapper {
	restOptions := rest.Options{
		Address: os.Getenv("CORE_HOST"),
		Timeout: time.Duration(3 * time.Minute),
		SkipTLS: false,
	}
	client := rest.New(restOptions)
	address = restOptions.Address

	return &wrapper{client: client}
}

func getRequestHeaders(ctx context.Context) (headers http.Header) {
	headers = http.Header{
		"Content-Type": []string{"application/json"},
	}

	return
}

func (w *wrapper) GetUtilityData(ctx context.Context, key string) (resp GetUtilityResponse, err error) {
	path := fmt.Sprintf("/utility/get?key=%v", key)

	fmt.Println("[Get Request]", address+path)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Get(ctx, path, headers)
	if err != nil {
		err = fmt.Errorf("[CORE] GetData error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[CORE] GetData return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[CORE] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[GetData Response]", address+path, resp)

	return
}

func (w *wrapper) SetUtility(ctx context.Context, req SetUtilityRequest) (resp SetUtilityResponse, err error) {
	path := "/utility/set"

	fmt.Println(ctx, "[SetUtility Request]", address+path, req)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Post(ctx, path, headers, req)
	if err != nil {
		err = fmt.Errorf("[Core] SetUtility error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Core] SetUtility return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Core] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[SetUtility Response]", address+path, resp)

	return
}

func (w *wrapper) SendEmail(ctx context.Context, req SendEmailRequest) (resp SendEmailResponse, err error) {
	path := "/utility/send-email"

	fmt.Println(ctx, "[SendEmail Request]", address+path, req)

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Post(ctx, path, headers, req)
	if err != nil {
		err = fmt.Errorf("[Core] SendEmail error: %v", err.Error())
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("[Core] SendEmail return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Core] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[SendEmail Response]", address+path, resp)

	return
}

func (w *wrapper) UploadImage(ctx context.Context, req *multipart.FileHeader) (resp UploadImageResponse, err error) {
	path := "/upload/image"

	src, err := req.Open()
	if err != nil {
		fmt.Println("[Core] UploadImage Error 1: ", err.Error())
		return 
	}
	defer src.Close()

	targetDir := "./"
	targetFileName := filepath.Join(targetDir, req.Filename)

	dst, err := os.Create(targetFileName)
	if err != nil {
		fmt.Println("[Core] UploadImage Error 2: ", err.Error())
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("[Core] UploadImage Error 3: ", err.Error())
		return 
	}

	files, err := os.Open(targetFileName)
	if err != nil {
		fmt.Println("[Core] UploadImage Error 4: ", err.Error())
		return 
	}
	defer files.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("images", targetFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(part, files)
	if err != nil {
		fmt.Println("error copy: ", err.Error())
		return
	}
	
	writer.Close()
	fmt.Println(ctx, "[UploadImage Request]", address+path, req)

	reqhttp, err := http.NewRequest("POST", address+path, &requestBody)
	if err != nil {
		fmt.Println("error http request declaration: ", err)
		return
	}

	reqhttp.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resphttp, err := client.Do(reqhttp)
	if err != nil {
		fmt.Println("error client do: ", err)
		return
	}
	defer resphttp.Body.Close()

	body, err := io.ReadAll(resphttp.Body)
	if err != nil {
		fmt.Println("error read all: ", err)
		return
	}

	os.Remove(targetFileName)
	if resphttp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[Core] UploadImage return non 200 http status code. got %d", resphttp.StatusCode)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Core] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[UploadImage Response]", address+path, resp)
	return
}
