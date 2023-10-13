/**
 * Created by GOLAND.
 * User: pengyu
 * Time: 2023/10/13 15:37
 */

package svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const OpenidUrl = "https://api.weixin.qq.com/sns/jscode2session"

type GetOpenidResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func GetOpenid(code, appid, secret string) (string, error) {
	query := url.Values{}
	query.Add("appid", appid)
	query.Add("secret", secret)
	query.Add("js_code", code)
	query.Add("grant_type", "authorization_code")

	response, err := http.Get(fmt.Sprintf("%s?%s", OpenidUrl, query.Encode()))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var data GetOpenidResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data.Openid == "" || data.Errcode != 0 {
		return "", errors.New(data.Errmsg)
	}

	return data.Openid, nil
}
