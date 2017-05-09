package goyht

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateContractFromTemplate(t *testing.T) {
	// t.SkipNow()
	cli := NewClient(Config{
		AppID:      "2017042710231700001",
		Password:   "qiaoyu001",
		APIGateway: YHTAPIGateway,
	})
	cli.AddUser("testUserID1", "15928009057", "香樟有限公司", "915201903470159141", UserTypeEnterprise, CertTypeLicence, true)

	_, err := cli.AddUser("testUserID2", "15928009058", "王二", "520103198801011430", UserTypePersonal, CertTypeIDCard, false)
	if err != nil {
		t.Fatal(err)
	}

	tokRsp, err := cli.UserToken("testUserID2")
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	after := now.AddDate(0, 2, 0)
	beginY, beginM, beginD := now.Date()
	endY, endM, endD := after.Date()
	holder := M{
		"${lessor}":           "香樟有限公司",
		"${lessorID}":         "915201903470159141",
		"${lessorName}":       "香樟",
		"${lessorAddress}":    "香樟Address",
		"${lessorPhone}":      "15928009057",
		"${lessee}":           "王二",
		"${lesseeID}":         "520103198801011430",
		"${lesseeAddress}":    "王二Address",
		"${lesseePhone}":      "15928009058",
		"${apartmentAddress}": "apartment Address",
		"${roomid}":           "room Name",
		"${staryear}":         fmt.Sprintf("%d", beginY),
		"${starmonth}":        fmt.Sprintf("%02d", beginM),
		"${starday}":          fmt.Sprintf("%02d", beginD),
		"${endyear}":          fmt.Sprintf("%d", endY),
		"${endmonth}":         fmt.Sprintf("%02d", endM),
		"${endday}":           fmt.Sprintf("%02d", endD),
		"${leasemonths}":      fmt.Sprintf("%d", 2),
		"${monthlyRent}":      fmt.Sprintf("%.2f", 100.0),
		"${aggregaterents}":   fmt.Sprintf("%.2f", 200.0),
		"${deposit}":          fmt.Sprintf("%.2f", 100.0),
		"${paystay}":          1,
		"${date}":             fmt.Sprintf("%d-%02d-%02d", beginY, beginM, beginD),
	}
	tRsp, err := cli.CreateTemplateContract("合同", "testContract123", "82288", tokRsp.Value.Token, false, holder)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("contract", tRsp.Value.ContractID)
	partnerA := Partner{AppUserID: "testUserID1", LocationName: "${lessor}"}
	partnerB := Partner{AppUserID: "testUserID2", LocationName: "${lessee}"}
	_, err = cli.AddPartner(tRsp.Value.ContractID, tokRsp.Value.Token, partnerA, partnerB)
	if err != nil {
		t.Fatal(err)
	}
}

// func TestAddUser(t *testing.T) {}
//
// func TestModifyPhoneNumber(t *testing.T) {}
//
// func TestModifyUserName(t *testing.T) {}
//
// func TestUserToken(t *testing.T) {}
//
// func TestCreateTemplateContract(t *testing.T) {}
//
// func TestCreateFileContract(t *testing.T) {}
//
// func TestAddPartner(t *testing.T) {}
//
// func TestSignContract(t *testing.T) {}
//
// func TestInvalidateContract(t *testing.T) {}
//
// func TestListContracts(t *testing.T) {}
//
// func TestLookupContractDetail(t *testing.T) {}
//
// func TestDownloadContract(t *testing.T) {}
