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
		return 
	}
	defer src.Close()

	targetDir := "./"
	targetFileName := filepath.Join(targetDir, req.Filename)

	dst, err := os.Create(targetFileName)
	if err != nil {
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return 
	}

	files, err := os.Open(targetFileName)
	if err != nil {
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

	headers := getRequestHeaders(ctx)
	body, status, err := w.client.Post(ctx, path, headers, requestBody)
	if err != nil {
		err = fmt.Errorf("[Core] UploadImage error: %v", err.Error())
		return
	}

	os.Remove(targetFileName)

	if status != http.StatusOK {
		err = fmt.Errorf("[Core] UploadImage return non 200 http status code. got %d", status)
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("[Core] Unmarshal Response Error %v", err.Error())
	}

	fmt.Println(ctx, "[UploadImage Response]", address+path, resp)

	return
}
