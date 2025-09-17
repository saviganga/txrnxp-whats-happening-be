package media

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	configs "txrnxp-whats-happening/config"
)

type ImageKit struct {
	PublicKey  string
	PrivateKey string
	Endpoint   string
}

func NewImageKit(env *configs.Config) *ImageKit {
	return &ImageKit{
		PublicKey:  env.ImageKitPublicKey,
		PrivateKey: env.ImageKitPrivateKey,
		Endpoint:   env.ImageKitURL,
	}
}

// UploadFile implements MediaStorageProvider
// func (ik *ImageKit) UploadFile(fileName string, fileData []byte) (string, error) {
// 	var b bytes.Buffer
// 	writer := multipart.NewWriter(&b)

// 	base64Content := base64.StdEncoding.EncodeToString(fileData)
// 	_ = writer.WriteField("file", "data:image/jpeg;base64,"+base64Content)
// 	_ = writer.WriteField("fileName", fileName)
// 	_ = writer.WriteField("folder", "/golang_uploads")

// 	if err := writer.Close(); err != nil {
// 		return "", err
// 	}

// 	req, err := http.NewRequest("POST", "https://upload.imagekit.io/api/v1/files/upload", &b)
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	req.SetBasicAuth(ik.PrivateKey, "")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)
// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("upload failed: %s", string(body))
// 	}

// 	return string(body), nil
// }

// func (ik *ImageKit) UploadFile(fileName string, fileData []byte) (string, error) {
// 	var b bytes.Buffer
// 	writer := multipart.NewWriter(&b)

// 	base64Content := base64.StdEncoding.EncodeToString(fileData)
// 	_ = writer.WriteField("file", "data:image/jpeg;base64,"+base64Content)
// 	_ = writer.WriteField("fileName", fileName)
// 	_ = writer.WriteField("folder", "/golang_uploads")

// 	if err := writer.Close(); err != nil {
// 		return "", err
// 	}

// 	req, err := http.NewRequest("POST", "https://upload.imagekit.io/api/v1/files/upload", &b)
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	req.SetBasicAuth(ik.PrivateKey, "")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)
// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("upload failed: %s", string(body))
// 	}

// 	// Extract the "url" field from the JSON response
// 	var jsonResp map[string]interface{}
// 	if err := json.Unmarshal(body, &jsonResp); err != nil {
// 		return "", fmt.Errorf("failed to parse response: %v", err)
// 	}

// 	url, ok := jsonResp["url"].(string)
// 	if !ok {
// 		return "", fmt.Errorf("failed to get url from response")
// 	}

// 	return url, nil
// }

func (ik *ImageKit) UploadFile(fileName string, fileData []byte) (string, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	base64Content := base64.StdEncoding.EncodeToString(fileData)
	_ = writer.WriteField("file", "data:image/jpeg;base64,"+base64Content)
	_ = writer.WriteField("fileName", fileName)
	_ = writer.WriteField("folder", "/golang_uploads")

	if err := writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://upload.imagekit.io/api/v1/files/upload", &b)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(ik.PrivateKey, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("upload failed: %s", string(body))
	}

	// Parse JSON and extract "filePath"
	var jsonResp map[string]interface{}
	if err := json.Unmarshal(body, &jsonResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	path, ok := jsonResp["filePath"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get filePath from response")
	}

	return path, nil // <-- store this in your DB
}

// RetrieveFile builds full URL
// func (ik *ImageKit) RetrieveFile(path string) string {
// 	return fmt.Sprintf("%s/%s", ik.Endpoint, path)
// }

func (ik *ImageKit) RetrieveFile(path string) string {
	url := fmt.Sprintf("%s/%s", ik.Endpoint, path)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to fetch image: %v\n", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("image not found or access denied: status %d\n", resp.StatusCode)
		return ""
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to read image body: %v\n", err)
		return ""
	}

	mimeType := resp.Header.Get("Content-Type")
	base64Content := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, base64Content)
}
