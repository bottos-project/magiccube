package minio

import (
	"errors"
	"fmt"
	//"log"
	log "github.com/cihub/seelog"
	//	"net/url"
	"github.com/minio/minio-go"
	"io"
	//"os"
	"time"
)

type MinioRepository struct {
	minioEndpoint  string
	minioAccessKey string
	minioSecretKey string
}

// NewMinioRepository creates a new MinioRepository
func NewMinioRepository(endpoint string, accessKey string, secretKey string) *MinioRepository {
	return &MinioRepository{minioEndpoint: endpoint, minioAccessKey: accessKey, minioSecretKey: secretKey}
}
func (r *MinioRepository) GetCacheURL(username string, objectName string) (string, error) {


	log.Info("get cache")
	useSSL := false
	log.Info("r.minioEndpoint")
	log.Info(r.minioEndpoint)
	log.Info("r.minioAccessKey")
	log.Info(r.minioAccessKey)
	log.Info("r.minioSecretKey")
	log.Info(r.minioSecretKey)
	

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		result := "failed"
		fmt.Println(err)
		return result, nil
	}
	if objectName == "" || username == "" {
		result := "invalid para"
		log.Info("invalid para")
		log.Info("objectName")
		log.Info(objectName)
		log.Info("username")
		log.Info(username)
		return result, errors.New("invalid para")
	}
	location := "us-east-1"
	err = minioClient.MakeBucket(username, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(username)
		if err == nil && exists {
			log.Info("We already own")
			log.Info(username)
		} else {
			log.Info(err)
		}
		log.Info(err)
	}
    log.Info("GetCacheURL")
	log.Info(objectName)
	CacheURL, err := minioClient.PresignedPutObject(username, objectName, 1000*time.Second)
	if err != nil {
		result := "get presigned put url failed"
		log.Info(err)
		return result, errors.New("get presigned url failed")
	}
	log.Info("get signed")
	log.Info(CacheURL)
	url := CacheURL.String()
	return url, nil
}
func (r *MinioRepository) GetCacheFile(username string, objectName string) (*minio.Object, error) {
    
	log.Info("get cache file")
	useSSL := false
	log.Info("r.minioEndpoint")
	log.Info(r.minioEndpoint)
	log.Info("r.minioAccessKey")
	log.Info(r.minioAccessKey)
	log.Info("r.minioSecretKey")
	log.Info(r.minioSecretKey)

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)

	if objectName == "" || username == "" {
		log.Info("invalid para")
		log.Info("objectName")
		log.Info(objectName)
		log.Info("username")
		log.Info(username)
	}
	file, err := minioClient.GetObject(username, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Info(err)
	}
	return file, nil
}
func (r *MinioRepository) PutFile(username string, objectName string, reader io.Reader, objectSize int64) (int64, error) {

	log.Info("put file")
	useSSL := false
	log.Info("r.minioEndpoint")
	log.Info(r.minioEndpoint)
	log.Info("r.minioAccessKey")
	log.Info(r.minioAccessKey)
	log.Info("r.minioSecretKey")
	log.Info(r.minioSecretKey)

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)

	if objectName == "" || username == "" {
		log.Info("invalid para")
		log.Info("objectName")
		log.Info(objectName)
		log.Info("username")
		log.Info(username)
	}
	n, err := minioClient.PutObject(username, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Info(err)
	}
	return n, nil
}
func (r *MinioRepository) ComposeFile(dst minio.DestinationInfo, srcs []minio.SourceInfo) error {

	log.Info("ComposeFile")
	useSSL := false
	log.Info("r.minioEndpoint")
	log.Info(r.minioEndpoint)
	log.Info("r.minioAccessKey")
	log.Info(r.minioAccessKey)
	log.Info("r.minioSecretKey")
	log.Info(r.minioSecretKey)

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		log.Info(err)
	}
	err = minioClient.ComposeObject(dst, srcs)
	if err != nil {
		log.Info(err)
	}
	return nil
}

func (r *MinioRepository) GetPutState(username string, objectName string) (int64, error) {
	
	log.Info("get put state")
	useSSL := false
    log.Info("r.minioEndpoint")
	log.Info(r.minioEndpoint)
	log.Info("r.minioAccessKey")
	log.Info(r.minioAccessKey)
	log.Info("r.minioSecretKey")
	log.Info(r.minioSecretKey)
	
	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		log.Info(err)
		return 0, err
	}

	reader, err := minioClient.GetObject(username, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Info(err)
		return 0, err
	}
	defer reader.Close()

	objectInfo, err := reader.Stat()
	if err != nil {
		log.Info(err)
		return 0, err
	}
	
	return objectInfo.Size, nil
}

func (r *MinioRepository) GetFileDownloadURL(username string, objectName string) (string, error) {
    
	log.Info("get file downloadURL")
	log.Info("r.minioEndpoint")
	log.Info(r.minioEndpoint)
	log.Info("r.minioAccessKey")
	log.Info(r.minioAccessKey)
	log.Info("r.minioSecretKey")
	log.Info(r.minioSecretKey)
	//useSSL := true
	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, false)
	if err != nil {
		result := "failed"
		return result, nil
	}

	if objectName == "" || username == "" {
		log.Info("invalid para")
		log.Info("objectName")
		log.Info(objectName)
		log.Info("username")
		log.Info(username)
	}
	//reqParams := make(url.Values)
	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	location, _ := minioClient.GetBucketLocation(username)
	log.Info(location)
	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject(username, objectName, time.Second*24*60*60, nil)

	if err != nil {
		result := "get url failed"
		return result, errors.New("get presigned url failed")
	}
	url := presignedURL.String()

	return url, nil
}
