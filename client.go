package goyht

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// constants about user types.
const (
	UserTypePersonal   = 1 // 个人
	UserTypeEnterprise = 2 // 企业
	UserTypePlatform   = 4 // 平台
)

// constants about certification types.
const (
	CertTypeIDCard   = 1 // 身份证
	CertTypePassport = 2 // 护照
	CertTypeOfficer  = 3 // 军官证
	CertTypeLicence  = 4 // 营业执照
	CertTypeOrgan    = 5 // 组织机构代码
	CertTypeSocial   = 6 // 社会代码
)

// constants for url and keys
const (
	YHTAPIGateway = "https://sdk.yunhetong.com/sdk"
	AppIDKey      = "appId"
	PasswordKey   = "passWord"
)

// Config contains configurations about YunHeTong service.
type Config struct {
	AppID      string
	Password   string
	APIGateway string
}

// Client handles all APIs for YunHeTong service.
type Client struct {
	config    Config
	tlsClient http.Client
}

// NewClient returns a *Client.
func NewClient(cfg Config) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := http.Client{Transport: tr}
	return &Client{
		config:    cfg,
		tlsClient: client,
	}
}

// AddUser imports user into YunHeTong service.
func (c *Client) AddUser(p AddUserParams) (*AddUserResponse, error) {
	paramMap, err := toMap(p)
	if err != nil {
		return nil, err
	}

	data, err := c.doHTTPRequest(p.URI(), paramMap)
	if err != nil {
		return nil, err
	}

	rsp := &AddUserResponse{}
	if err = json.NewDecoder(bytes.NewReader(data)).Decode(rsp); err != nil {
		return nil, err
	}

	if rsp.Code != "200" || rsp.SubCode != "200" {
		return nil, fmt.Errorf("code %s subcode %s msg %s", rsp.Code, rsp.SubCode, rsp.Message)
	}

	return rsp, nil
}

// AsyncNotifyResult represents the result returned from YunHeTong service.
type AsyncNotifyResult struct{}

// AsyncNotify returns asynchronous notification from YunHeTong service.
func (c *Client) AsyncNotify(req *http.Request) (*AsyncNotifyResult, error) {
	return nil, errors.New("not defined")
}

func (c *Client) doHTTPRequest(uri string, paramMap map[string]string) ([]byte, error) {
	if token, ok := paramMap["token"]; ok {
		delete(paramMap, "token")
		uri = fmt.Sprintf("%s?token=%s", uri, token)
	}
	paramMap[AppIDKey] = c.config.AppID
	paramMap[PasswordKey] = c.config.Password

	formData := url.Values{}
	for k, v := range paramMap {
		formData.Add(k, v)
	}

	rsp, err := c.tlsClient.PostForm(fmt.Sprintf("%s%s", c.config.APIGateway, uri), formData)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
