package goyht

// AddUserParam represents paramters for /userInfo/addUser.
type AddUserParams struct {
	AppUserID       string `param:"appUserId"`
	CellNum         string `param:"cellNum"`
	UserType        string `param:"userType"`
	UserName        string `param:"userName"`
	CertifyType     string `param:"certifyType"`
	CertifyNumber   string `param:"certifyNumber"`
	CreateSignature string `param:"createSignature"`
}

// URI returns the URL of API.
func (p AddUserParams) URI() string {
	return "/userInfo/addUser"
}

// AddUserResponse represents the reponse returned.
type AddUserResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
}

// ModifyPhoneNumberParams represents paramters for /userInfo/modifyCellNum.
type ModifyPhoneNumberParams struct {
	CellNum string `param:"cellNum"`
}

// URI returns the URL of API.
func (p ModifyPhoneNumberParams) URI() string {
	return "/userInfo/modifyCellNum"
}

// ModifyPhoneNumberResponse represents the reponse returned.
type ModifyPhoneNumberResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
}

// ModifyUserNameParams represents paramters for /userInfo/modifyUserName.
type ModifyUserNameParams struct {
	CellNum string `param:"cellNum"`
}

// URI returns the URL of API.
func (p ModifyUserNameParams) URI() string {
	return "/userInfo/modifyUserName"
}

// ModifyUserNameResponse represents the reponse returned.
type ModifyUserNameResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
}

// UserTokenParams represents paramters for /token/getToken.
type UserTokenParams struct{}

// UserTokenResponse represents the reponse returned.
type UserTokenResponse struct{}

// CreateTemplateContractParams represents paramters for /contract/templateContract?
type CreateTemplateContractParams struct{}

// CreateTemplateContractResponse represents the reponse returned.
type CreateTemplateContractResponse struct{}

// CreateFileContractParams represents paramters for /contract/fileContract
type CreateFileContractParams struct{}

// CreateFileContractResponse represents the reponse returned.
type CreateFileContractResponse struct{}

// AddPartnerParams represents paramters for /contract/addPartner
type AddPartnerParams struct{}

// AddPartnerResponse represents the reponse returned.
type AddPartnerResponse struct{}

// SignContractParams represents paramters for /contract/signContract
type SignContractParams struct{}

// SignContractResponse represents the reponse returned.
type SignContractResponse struct{}

// InvalidateContractParams represents paramters for /contract/invalid
type InvalidateContractParams struct{}

// InvalidateContractResponse represents the reponse returned.
type InvalidateContractResponse struct{}

// ListContractsParams represents paramters for /contract/list
type ListContractsParams struct{}

// ListContractsResponse represents the reponse returned.
type ListContractsResponse struct{}

// LookupContractDetailParams represents paramters for /contract/detail
type LookupContractDetailParams struct{}

// LookupContractDetailResponse represents the reponse returned.
type LookupContractDetailResponse struct{}

// DownloadContractParams represents paramters for /contract/download
type DownloadContractParams struct{}

// DownloadContractResponse represents the reponse returned.
type DownloadContractResponse struct{}
