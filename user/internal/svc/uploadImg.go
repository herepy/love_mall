/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/20 9:38
 */

package svc

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"time"
)

func UploadImg(r *http.Request, accessKey, secret, endpoint, bucketName, key string) (string, error) {
	client, err := oss.New(endpoint, accessKey, secret)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	objectName := fmt.Sprintf("image/%d.png", time.Now().Unix())
	err = r.ParseMultipartForm(5242880)
	if err != nil {
		return "", err
	}

	fileList, ok := r.MultipartForm.File[key]
	if !ok {
		return "", errors.New("upload file not found with key:" + key)
	}

	file, err := fileList[0].Open()
	if err != nil {
		return "", err
	}

	defer file.Close()
	err = bucket.PutObject(objectName, file)

	return fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, objectName), err
}
