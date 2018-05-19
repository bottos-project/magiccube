package minio

import (
	"errors"
	"fmt"
	"log"
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

	fmt.Println("get cache")
	useSSL := false
	fmt.Println(r.minioEndpoint)
	fmt.Println(r.minioAccessKey)
	fmt.Println(r.minioSecretKey)

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		result := "failed"
		fmt.Println(err)
		return result, nil
	}
	if objectName == "" || username == "" {
		result := "invalid para"
		log.Println("invalid para", objectName, username)
		return result, errors.New("invalid para")
	}
	location := "us-east-1"
	err = minioClient.MakeBucket(username, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(username)
		if err == nil && exists {
			log.Printf("We already own %s\n", username)
		} else {
			log.Fatalln(err)
		}
		log.Println(err)
	}
	fmt.Println("GetCacheURL %s", objectName)

	CacheURL, err := minioClient.PresignedPutObject(username, objectName, 1000*time.Second)
	if err != nil {
		result := "get presigned put url failed"
		log.Println(err)
		return result, errors.New("get presigned url failed")
	}
	fmt.Println("get signed")
	fmt.Println(CacheURL)
	url := CacheURL.String()
	return url, nil
}
func (r *MinioRepository) GetCacheFile(username string, objectName string) (*minio.Object, error) {

	fmt.Println("get cache file")
	useSSL := false
	fmt.Println(r.minioEndpoint)
	fmt.Println(r.minioAccessKey)
	fmt.Println(r.minioSecretKey)

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)

	if objectName == "" || username == "" {
		log.Println("invalid para", objectName, username)
		fmt.Println("invalid para")
	}
	file, err := minioClient.GetObject(username, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	return file, nil
}
func (r *MinioRepository) PutFile(username string, objectName string, reader io.Reader, objectSize int64) (int64, error) {

	fmt.Println("put file")
	useSSL := false
	fmt.Println(r.minioEndpoint)
	fmt.Println(r.minioAccessKey)
	fmt.Println(r.minioSecretKey)

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)

	if objectName == "" || username == "" {
		log.Println("invalid para", objectName, username)
		fmt.Println("invalid para")
	}
	n, err := minioClient.PutObject(username, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}
	return n, nil
}
func (r *MinioRepository) ComposeFile(dst minio.DestinationInfo, srcs []minio.SourceInfo) error {

	fmt.Println("ComposeFile")
	useSSL := false
	fmt.Println(r.minioEndpoint)
	fmt.Println(r.minioAccessKey)
	fmt.Println(r.minioSecretKey)

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		fmt.Println(err)
	}
	err = minioClient.ComposeObject(dst, srcs)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (r *MinioRepository) GetPutState(username string, objectName string) (int64, error) {

	useSSL := false

	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, useSSL)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	reader, err := minioClient.GetObject(username, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println("getobject")
		log.Println(err)
		return 0, err
	}
	defer reader.Close()

	objectInfo, err := reader.Stat()
	if err != nil {
		fmt.Println("stat")
		log.Println(err)
		return 0, err
	}
	fmt.Println("size", objectInfo.Size)
	fmt.Println("gooo", objectInfo)
	return objectInfo.Size, nil
}

func (r *MinioRepository) GetFileDownloadURL(username string, objectName string) (string, error) {

	//useSSL := true
	minioClient, err := minio.New(r.minioEndpoint, r.minioAccessKey, r.minioSecretKey, false)
	if err != nil {
		result := "failed"
		return result, nil
	}
	fmt.Println(objectName)
	fmt.Println(username)
	if objectName == "" || username == "" {
		result := "invalid para"
		return result, errors.New("invalid para")
	}
	//reqParams := make(url.Values)
	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
	location, _ := minioClient.GetBucketLocation(username)
	fmt.Println("location", location)
	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject(username, objectName, time.Second*24*60*60, nil)

	if err != nil {
		result := "get url failed"
		fmt.Println(err)
		return result, errors.New("get presigned url failed")
	}
	url := presignedURL.String()
	fmt.Println(url)
	return url, nil
}
