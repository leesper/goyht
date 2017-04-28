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

func httpRequest(c *Client, p Params, paramMap map[string]string, factory func() interface{}) (interface{}, error) {
	data, err := c.doHTTPRequest(p.URI(), paramMap)
	if err != nil {
		return nil, err
	}

	rsp := factory()
	if err = json.NewDecoder(bytes.NewReader(data)).Decode(rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// AddUser imports user into YunHeTong service.
func (c *Client) AddUser(p AddUserParams) (*AddUserResponse, error) {
	paramMap, err := toMap(p, map[string]string{
		AppIDKey:    c.config.AppID,
		PasswordKey: c.config.Password,
	})
	if err != nil {
		return nil, err
	}

	ret, err := httpRequest(c, p, paramMap, func() interface{} {
		return &AddUserResponse{}
	})

	if err != nil {
		return nil, err
	}

	rsp := ret.(*AddUserResponse)

	if err = checkErr(rsp.Code, rsp.SubCode, rsp.Message); err != nil {
		return nil, err
	}

	return rsp, nil
}

// ModifyPhoneNumber modifies user's cell phone number.
func (c *Client) ModifyPhoneNumber(p ModifyPhoneNumberParams) (*ModifyPhoneNumberResponse, error) {
	paramMap, err := toMap(p, nil)
	if err != nil {
		return nil, err
	}

	ret, err := httpRequest(c, p, paramMap, func() interface{} {
		return &ModifyPhoneNumberResponse{}
	})

	if err != nil {
		return nil, err
	}

	rsp := ret.(*ModifyPhoneNumberResponse)

	if err = checkErr(rsp.Code, rsp.SubCode, rsp.Message); err != nil {
		return nil, err
	}

	return rsp, nil
}

// ModifyUserName modifies user's name.
func (c *Client) ModifyUserName(p ModifyUserNameParams) (*ModifyUserNameResponse, error) {
	paramMap, err := toMap(p, nil)
	if err != nil {
		return nil, err
	}

	ret, err := httpRequest(c, p, paramMap, func() interface{} {
		return &ModifyUserNameResponse{}
	})

	if err != nil {
		return nil, err
	}

	rsp := ret.(*ModifyUserNameResponse)

	if err = checkErr(rsp.Code, rsp.SubCode, rsp.Message); err != nil {
		return nil, err
	}

	return rsp, nil
}

// UserToken gets user's token string.
func (c *Client) UserToken(p UserTokenParams) (*UserTokenResponse, error) {
	paramMap, err := toMap(p, map[string]string{
		AppIDKey:    c.config.AppID,
		PasswordKey: c.config.Password,
	})
	if err != nil {
		return nil, err
	}

	ret, err := httpRequest(c, p, paramMap, func() interface{} {
		return &UserTokenResponse{}
	})

	if err != nil {
		return nil, err
	}

	rsp := ret.(*UserTokenResponse)

	if err = checkErr(rsp.Code, rsp.SubCode, rsp.Message); err != nil {
		return nil, err
	}

	return rsp, nil
}

// CreateTemplateContract creates contract based on template.
func (c *Client) CreateTemplateContract() {}

// CreateFileContract creates contract by uploading file.
func (c *Client) CreateFileContract() {}

// AddPartner adds partners of contract.
func (c *Client) AddPartner() {}

// SignContract signs contract automatically.
func (c *Client) SignContract() {}

// InvalidateContract invalidates contract.
func (c *Client) InvalidateContract() {}

// ListContracts returns a list of contracts finished or invalidated.
func (c *Client) ListContracts() {}

// LookupContractDetail returns the detail of a contract.
func (c *Client) LookupContractDetail() {}

// DownloadContract downloads a contract.
func (c *Client) DownloadContract() {}

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
