package minio

import (
	"errors"
	"github.com/minio/minio-go"
	"time"
)

type StorageRepository struct {
	minioEndpoint  string
	minioAccessKey string
	minioSecretKey string
}

// NewMinioRepository creates a new StorageRepository
func NewMinioRepository(endpoint string, accessKey string, secretKey string) *StorageRepository {
	return &StorageRepository{minioEndpoint: endpoint, minioAccessKey: accessKey, minioSecretKey: secretKey}
}
func (r *StorageRepository) GetPutURL(username string, objectName string) (string, error) {

	useSSL := true
	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		result := "failed"
		return result, nil
	}
	if objectName == "" || username == "" {
		result := "invalid para"
		return result, errors.New("invalid para")
	}
	presignedURL, err := minioClient.PresignedPutObject(username, objectName, 1000*time.Second)
	if err != nil {
		result := "get presigned put url failed"
		return result, errors.New("get presigned url failed")
	}
	url := presignedURL.String()
	return url, nil
}
func (r *StorageRepository) GetFileDownloadURL(username string, objectName string) (string, error) {

	useSSL := true
	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		result := "failed"
		return result, nil
	}
	if objectName == "" || username == "" {
		result := "invalid para"
		return result, errors.New("invalid para")
	}
	//reqParams := make(url.Values)
	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject("mybucket", "myobject", time.Second * 24 * 60 * 60, nil)
	//presignedURL, err := minioClient.PresignedGetObject("mybucket", "myobject", time.Second * 24 * 60 * 60, reqParams)
	if err != nil {
		result := "get presigned get url failed"
		return result, errors.New("get presigned url failed")
	}
	url := presignedURL.String()
	return url, nil
}

//
//
//// connet
//func (r *StorageRepository) connect(access MinioRepository) (result string, err error) {
//	useSSL := true
//
//	// Initialize minio client object.
//	minioClient, err := minio.New(access.minioHostString, access.accessKeyID, access.secretAccessKey, useSSL)
//	if err != nil {
//		//	log.Fatalln(err)
//		result := "failed"
//	}
//	result := "OK"
//	return result, error
//
//}
//func (api mediaHandlers) GetPresignedURLHandler(w http.ResponseWriter, r *http.Request) {
//	// The object for which the presigned URL has to be generated is sent as a query
//	// parameter from the client.
//	objectName := r.URL.Query().Get("objName")
//	if objectName == "" {
//		http.Error(w, "No object name set, invalid request.", http.StatusBadRequest)
//		return
//	}
//	presignedURL, err := api.storageClient.PresignedGetObject(*bucketName, objectName, 1000*time.Second, nil)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Write([]byte(presignedURL))
//}
