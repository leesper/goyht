## Constants
constants about user types.
```go
const (
    UserTypePersonal   = "1" // 个人
    UserTypeEnterprise = "2" // 企业
    UserTypePlatform   = "4" // 平台
)
constants about certification types.

const (
    CertTypeIDCard   = "1" // 身份证
    CertTypePassport = "2" // 护照
    CertTypeOfficer  = "3" // 军官证
    CertTypeLicence  = "4" // 营业执照
    CertTypeOrgan    = "5" // 组织机构代码
    CertTypeSocial   = "6" // 社会代码
)
constants for url and keys

const (
    YHTAPIGateway  = "https://sdk.yunhetong.com/sdk"
    YHTAuthGateway = "https://authentic.yunhetong.com"
    AppIDKey       = "appId"
    PasswordKey    = "password"
)
```

## type AddPartnerResponse

AddPartnerResponse represents the reponse returned.

```go
type AddPartnerResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
}
```

## type AddUserResponse

AddUserResponse represents the reponse returned.

```go
type AddUserResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
}
```

## type AsyncNotifyResult

AsyncNotifyResult represents the result returned from YunHeTong service.
```go
type AsyncNotifyResult struct {
    Content      string                 `json:"content"`
    NoticeType   int                    `json:"noticeType"`
    NoticeParams string                 `json:"noticeParams"`
    InfoMap      map[string]interface{} `json:"map"`
}
```

## type AuthResponse

AuthResponse represents the response returned.

```go
type AuthResponse struct {
    Code    int    `json:"code"`
    Msg     string `json:"msg"`
    Success bool   `json:"success"`
    Data    string `json:"data"`
    Message string
    Status  string
}
```

## type Client

Client handles all APIs for YunHeTong service.

```go
type Client struct {
    // contains filtered or unexported fields
}
```

## func NewClient

```go
func NewClient(cfg Config) *Client
```
NewClient returns a \*Client.

## func (\*Client) AddPartner

```go
func (c *Client) AddPartner(contractID int64, token string, partners ...Partner) (*AddPartnerResponse, error)
```
AddPartner adds partners of contract.

## func (\*Client) AddUser
```go
func (c *Client) AddUser(userID, phone, name, certNum string, userType string, certType string, autoSign bool) (*AddUserResponse, error)
```
AddUser imports user into YunHeTong service.

## func (\*Client) AnswerAsyncNotify
```go
func (c *Client) AnswerAsyncNotify(rsp bool, msg string) string
```
AnswerAsyncNotify returns a json string answering async notification.

## func (\*Client) AsyncNotify
```go
func (c *Client) AsyncNotify(req *http.Request) (*AsyncNotifyResult, error)
```
AsyncNotify returns asynchronous notification from YunHeTong service.

## func (\*Client) AuthRealName
```go
func (c *Client) AuthRealName(idNum, idName string, portrait bool) (*AuthResponse, error)
```
AuthRealName authenticates ID number and name via YunHeTong service.

## func (\*Client) AuthRealNameBank
```go
func (c *Client) AuthRealNameBank(idNum, idName, bankCard, mobile string) (*AuthResponse, error)
```

## func (\*Client) CreateFileContract
```go
func (c *Client) CreateFileContract(title, contractNo, token string, useCer bool, data []byte) (*CreateFileContractResponse, error)
```

CreateFileContract creates contract by uploading file.

## func (\*Client) CreateTemplateContract
```go
func (c *Client) CreateTemplateContract(title, contractNo, templateID, token string, useCer bool, placeHolders M) (*CreateTemplateContractResponse, error)
```
CreateTemplateContract creates contract based on template.

## func (\*Client) DownloadContract
```go
func (c *Client) DownloadContract(contractID, token string) ([]byte, error)
```
DownloadContract downloads a contract.

## func (\*Client) InvalidateContract
```go
func (c *Client) InvalidateContract(contractID, token string) (*InvalidateContractResponse, error)
```
InvalidateContract invalidates contract.

## func (\*Client) ListContracts
```go
func (c *Client) ListContracts(pageNum, pageSize int, token string) (*ListContractsResponse, error)
```
ListContracts returns a list of contracts finished or invalidated.

## func (\*Client) LookupContractDetail
```go
func (c *Client) LookupContractDetail(contractID, token string) (*LookupContractDetailResponse, error)
```
LookupContractDetail returns the detail of a contract.

## func (\*Client) ModifyPhoneNumber
```go
func (c *Client) ModifyPhoneNumber(phone, token string) (*ModifyPhoneNumberResponse, error)
```
ModifyPhoneNumber modifies user's cell phone number.

## func (\*Client) ModifyUserName
```go
func (c *Client) ModifyUserName(name, token string, autoSign bool) (*ModifyUserNameResponse, error)
```
ModifyUserName modifies user's name.

## func (\*Client) SignContract
```go
func (c *Client) SignContract(contractID, token string, signers ...string) (*SignContractResponse, error)
```
SignContract signs contract automatically.

## func (\*Client) UserToken
```go
func (c *Client) UserToken(userID string) (*UserTokenResponse, error)
```
UserToken gets user's token string.

## type Config

Config contains configurations about YunHeTong service.
```go
type Config struct {
    AppID       string
    Password    string
    APIGateway  string
    AuthID      string
    AuthPWD     string
    AuthGateway string
}
```

## type CreateFileContractResponse
CreateFileContractResponse represents the reponse returned.
```go
type CreateFileContractResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
    Value   struct {
        ContractID string `json:"contractId"`
    } `json:"value"`
}
```

## type CreateTemplateContractResponse

CreateTemplateContractResponse represents the reponse returned.
```go
type CreateTemplateContractResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
    Value   struct {
        ContractID int64 `json:"contractId"`
    } `json:"value"`
}
```

## type DownloadContractResponse

DownloadContractResponse represents the reponse returned.
```go
type DownloadContractResponse struct {
    File []byte
}
```

## type InvalidateContractResponse

InvalidateContractResponse represents the reponse returned.
```go
type InvalidateContractResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
}
```

## type ListContractsResponse

ListContractsResponse represents the reponse returned.
```go
type ListContractsResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
    Value   struct {
        ContractList []struct {
            ID          string `json:"id"`
            Title       string `json:"title"`
            Status      string `json:"status"`
            AppName     string `json:"appName"`
            GmtModify   string `json:"gmtModify"`
            PartnerList string `json:"partnerList"`
        } `json:"contractList"`
    } `json:"value"`
}
```

## type LookupContractDetailResponse

LookupContractDetailResponse represents the reponse returned.
```go
type LookupContractDetailResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
    Value   struct {
        PartnerList []struct {
            SignStatus string `json:"signStatus"`
            UserID     string `json:"userId"`
        } `json:"partnerList"`
        Title  string `param:"title"`
        Status string `json:"status"`
    } `json:"value"`
}
```

## type M

M is a convenient alias for a map[string]interface{} map.
```go
type M map[string]interface{}
```

## type ModifyPhoneNumberResponse

ModifyPhoneNumberResponse represents the reponse returned.
```go
type ModifyPhoneNumberResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
}
```

## type ModifyUserNameResponse

ModifyUserNameResponse represents the reponse returned.
```go
type ModifyUserNameResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
}
```

## type Partner

Partner represents a participant in a contract.
```go
type Partner struct {
    AppUserID    string `json:"appUserId"`
    LocationName string `json:"locationName,omitempty"` // 模板签名占位符名称(与keyWord必填其一)
    Keyword      string `json:"keyWord,omitempty"`
}
```

## type SignContractResponse

SignContractResponse represents the reponse returned.
```go
type SignContractResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
}
```

## type UserTokenResponse

UserTokenResponse represents the reponse returned.
```go
type UserTokenResponse struct {
    Code    int    `json:"code"`
    SubCode int    `json:"subCode"`
    Message string `json:"message"`
    Value   struct {
        Token string `json:"token"`
    } `json:"value"`
}
```
