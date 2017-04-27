package goyht

// AddUserParam represents paramters for /userInfo/addUser.
type AddUserParam struct {
	AppUserID       string `param:"appUserId"`
	CellNum         string `param:"cellNum"`
	UserType        string `param:"userType"`
	UserName        string `param:"userName"`
	CertifyType     string `param:"certifyType"`
	CertifyNumber   string `param:"certifyNumber"`
	CreateSignature string `param:"createSignature"`
}

// URI returns the URL of API.
func (p AddUserParam) URI() string {
	return "/userInfo/addUser"
}

// AddUserResponse represents the reponse returned by /userInfo/addUser.
type AddUserResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"subCode"`
	Message string `json:"message"`
}
