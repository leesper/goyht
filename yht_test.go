package goyht

import (
	"fmt"
	"testing"
	"time"
)

func TestAuthIDCard(t *testing.T) {
	cli := NewClient(Config{
		AppID:       "XXX",
		Password:    "YYY",
		APIGateway:  YHTAPIGateway,
		AuthID:      "4502d72d02604fbdbcc41f488a760e98",
		AuthPWD:     "d72a39c0e8ca4cb1a831a05e9c699b9a",
		AuthGateway: YHTAuthGateway,
	})
	rsp, err := cli.AuthIDCard("520103198712312831", "李科君", false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rsp.Data)
}

func TestCreateContractFromTemplate(t *testing.T) {
	t.SkipNow()
	cli := NewClient(Config{
		AppID:       "XXX",
		Password:    "YYY",
		APIGateway:  YHTAPIGateway,
		AuthID:      "4502d72d02604fbdbcc41f488a760e98",
		AuthPWD:     "d72a39c0e8ca4cb1a831a05e9c699b9a",
		AuthGateway: YHTAuthGateway,
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
	tRsp, err := cli.CreateTemplateContract("合同", "testContract123", "92130", tokRsp.Value.Token, false, holder)
	if err != nil {
		t.Fatal(err)
	}

	partnerA := Partner{AppUserID: "testUserID1", LocationName: "56006"}
	partnerB := Partner{AppUserID: "testUserID2", LocationName: "02289"}
	_, err = cli.AddPartner(tRsp.Value.ContractID, tokRsp.Value.Token, partnerA, partnerB)
	if err != nil {
		_, err = cli.AddPartner(tRsp.Value.ContractID, tokRsp.Value.Token, partnerA, partnerB)
		if err != nil {
			t.Fatal(err)
		}
	}
}
