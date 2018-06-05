/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
 */
 
package minio

import (
	"errors"
	"fmt"
	"log"
	//"net/url"
	"time"

	"github.com/minio/minio-go"
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
func (r *MinioRepository) GetPutURL(username string, objectName string) (string, error) {

	fmt.Println("get put")
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
	fmt.Println("GetPutURL %s", objectName)

	presignedURL, err := minioClient.PresignedPutObject(username, objectName, 1000*time.Second)
	if err != nil {
		result := "get presigned put url failed"
		log.Println(err)
		return result, errors.New("get presigned url failed")
	}
	fmt.Println("get signed")
	fmt.Println(presignedURL)
	url := presignedURL.String()
	return url, nil
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
