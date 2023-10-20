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
	"path"
	"time"
)

func UploadImg(r *http.Request, accessKey, secret, endpoint, bucketName, key string) (string, error) {
	file, header, err := r.FormFile(key)
	if err != nil {
		return "", err
	}

	ext := path.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" && ext != ".bmp" {
		return "", errors.New("not image file")
	}

	client, err := oss.New(endpoint, accessKey, secret)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	objectName := fmt.Sprintf("image/%d%s", time.Now().Unix(), ext)

	defer file.Close()
	err = bucket.PutObject(objectName, file)

	return fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, objectName), err
}
