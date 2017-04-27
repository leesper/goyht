package goyht

// AddUserParams represents paramters for /userInfo/addUser.
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
type UserTokenParams struct {
	AppUserID string `param:"appUserId"`
}

// UserTokenResponse represents the reponse returned.
type UserTokenResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
	Value   struct {
		Token string `json:"token"`
	} `json:"value"`
}

// CreateTemplateContractParams represents paramters for /contract/templateContract?
type CreateTemplateContractParams struct {
	Title         string `param:"title"`
	DefContractNo string `param:"defContractNo"`
	TemplateID    string `param:"templateId"`
	UseCer        string `param:"useCer"`
	Param         string `param:"param"`
}

// CreateTemplateContractResponse represents the reponse returned.
type CreateTemplateContractResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
	Value   struct {
		ContractID string `json:"contractId"`
	} `json:"value"`
}

// CreateFileContractParams represents paramters for /contract/fileContract
type CreateFileContractParams struct {
	Title         string `param:"title"`
	DefContractNo string `param:"defContractNo"`
	TemplateID    string `param:"templateId"`
	UseCer        string `param:"useCer"`
	File          []byte `param:"file"`
}

// CreateFileContractResponse represents the reponse returned.
type CreateFileContractResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
	Value   struct {
		ContractID string `json:"contractId"`
	} `json:"value"`
}

// AddPartnerParams represents paramters for /contract/addPartner
type AddPartnerParams struct {
	ContractID string `param:"contractId"`
	Partners   string `param:"partners"`
}

// AddPartnerResponse represents the reponse returned.
type AddPartnerResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
}

// SignContractParams represents paramters for /contract/signContract
type SignContractParams struct {
	ContractID string `param:"contractId"`
	Signer     string `param:"signer"`
}

// SignContractResponse represents the reponse returned.
type SignContractResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
}

// InvalidateContractParams represents paramters for /contract/invalid
type InvalidateContractParams struct {
	ContractID string `param:"contractId"`
}

// InvalidateContractResponse represents the reponse returned.
type InvalidateContractResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
}

// ListContractsParams represents paramters for /contract/list
type ListContractsParams struct {
	PageNum  string `param:"pageNum"`
	PageSize string `param:"pageSize"`
}

// ListContractsResponse represents the reponse returned.
type ListContractsResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
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

// LookupContractDetailParams represents paramters for /contract/detail
type LookupContractDetailParams struct {
	ContractID string `param:"contractId"`
}

// LookupContractDetailResponse represents the reponse returned.
type LookupContractDetailResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
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

// DownloadContractParams represents paramters for /contract/download
type DownloadContractParams struct {
	ContractID string `param:"contractId"`
}

// DownloadContractResponse represents the reponse returned.
type DownloadContractResponse struct {
	File []byte
}
