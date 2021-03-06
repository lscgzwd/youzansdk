package youzansdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lscgzwd/youzansdk/request"
	"github.com/lscgzwd/youzansdk/utils"
)

const aesIV = "0102030405060708"

// YouzanClient ..
type YouzanClient struct {
	Token           string
	URL             string
	ClientSecretKey string
	ClientID        string
}

func (c YouzanClient) DecryptMessage(message string) (string, error) {
	return utils.AesCBCDecrypt(message, c.ClientSecretKey[0:16], aesIV)
}

// Execute 执行请求
func (c YouzanClient) Execute(request request.BaseRequest, v interface{}) (string, error) {
	if err := request.CheckParam(); err != nil {
		return "", err
	}
	// apiName := request.GetApiName()
	// idx := strings.LastIndex(apiName, ".")
	// service := apiName[:idx]
	// action := apiName[idx+1 : len(apiName)]
	url := request.GetRequestUrl(c.Token)
	var response *http.Response
	var err error
	if request.GetMethod() != "POSTFORM" {
		param := request.GetParam()
		if nil == param {
			param = map[string]string{}
		}
		param["access_token"] = c.Token
		// url = fmt.Sprintf("%s?access_token=%s", url, c.Token)
		response, err = utils.HTTP(url, request.GetMethod(), request.GetBodyParam(), param, nil)
	} else {
		url = fmt.Sprintf("%s?access_token=%s", url, c.Token)
		response, err = utils.HttpPostForm(url, request.GetUrlValues())
	}
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("read response body fail")
	}
	println(fmt.Sprintf("youzansdk response:%s", string(responseBody)))
	err = json.Unmarshal(responseBody, v)

	return string(responseBody), err
}
