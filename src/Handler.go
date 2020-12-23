package src

import (
	"errors"
	"fmt"
	"github.com/asrx/go-kinghood-api-wrapper/src/Conf"
	"github.com/asrx/go-kinghood-api-wrapper/src/Entry"
	"github.com/idoubi/goz"
)

type KinghoodService struct {
	key string
}

func NewKinghoodService(key string) *KinghoodService {
	return &KinghoodService{
		key: key,
	}
}

func (s KinghoodService)mergeKey(mapParam map[string]interface{}) {
	mapParam["Key"] = s.key
}

func (s KinghoodService)postRequest(url string, mapParam map[string]interface{}) (content string, err error) {

	s.mergeKey(mapParam)

	client := goz.NewClient()
	resp, err := client.Post(url, goz.Options{
		Headers:    map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		FormParams: mapParam,
	})
	if err != nil {
		return
	}

	responseBody, err1 := resp.GetBody()
	if err1 != nil {
		err = err1
		return
	}
	content = responseBody.String()
	return
}

// 获取余额
func (s KinghoodService)GetBalance() (money float64, err error){
	var mapParam = make(map[string]interface{})
	content, err1 := s.postRequest(Conf.GetApiBalance(), mapParam)

	if err1 != nil {
		err = err1
		fmt.Println("Client Post Request Error", err1)
		return
	}

	// Result
	balance := Entry.Balance{}
	resp, err := Entry.ResponseResult(content, &balance)

	if resp.Code != "200" {
		err = errors.New("获取额度失败，" + resp.Message)
		return
	}

	if err != nil {
		fmt.Println("Response to Struct Error", err)
		return
	}
	money = balance.Money
	return
}

// 获取服务
func (s KinghoodService)GetCarriers() (content string, err error){
	var mapParam = make(map[string]interface{})
	content, err1 := s.postRequest(Conf.GetApiCarriers(), mapParam)
	if err1 != nil {
		err = err1
		fmt.Println("Client Post Request Error", err1)
		return
	}
	return
}

// Rate
func (s KinghoodService)GetRate(params string) (rateReply *Entry.RateReply, err error) {
	//content := `{"Code":"200","Message":"","Data":[{"RateId":"0","Carrier":"UPS Ground 6","CarrierServiceLevel":"UPS Ground","CarrierServiceName":"UPS Ground","ChargeWeight":40,"TotalCharge":32.40,"PackageType":null,"Message":null,"Zone":null,"IsResidential":false,"PublishCharge":0,"IsPublishChargeSpecified":false}]}`
	var mapParam = make(map[string]interface{})
	mapParam["Data"] = params

	content, err1 := s.postRequest(Conf.GetApiRate(), mapParam)
	if err1 != nil {
		fmt.Println("Client Post Request Error", err1)
		return
	}
	// Result
	list := []*Entry.RateReply{}
	//reply := Entry.Reply{}

	reply, err1 := Entry.ResponseResult(content, &list)

	if err1 != nil {
		err = err1
		fmt.Println("Response to Struct Error", err1)
		return
	}

	if len(list) == 0 {
		err = errors.New(reply.Message)
		return
	}
	rateReply =  list[0]
	return
}

//Ship
func (s KinghoodService)GetShip(params string) (shipReply *Entry.ShipReply,err error) {
	//content := `{"Code":"200","Message":null,"Data":{"EcOrder":null,"FreightCalss":null,"ShipmentId":null,"LabelReady":true,"LabelUrl":"https://www.ex17usa.com/UploadFiles/UD/2020-12-02/90c7e355-67c5-4223-b077-3d6bf1d408cb.pdf","ShippingWeight":8,"FactWeight":0,"ShippingPrice":6.48,"TrackingNumbers":["1Z7A03740311821113"],"TN2":null,"Carrier":"UPS Ground 6","CarrierServiceLevel":"UPS Ground","ShipFrom":{"PostalCode":"92879","Name":"micheal","Address1":"2565 SAMPSON AVENUE","Address2":"","State":"CA","City":"CORONA","Country":"US","CountryCode":null,"PhoneNumber":"09512680000","AddrType":1,"Company":""},"ShipTo":{"PostalCode":"92551","Name":"Donovan","Address1":"24300 Nandina Ave","Address2":null,"State":"CA","City":"Moreno Valley","Country":"US","CountryCode":null,"PhoneNumber":"13474483190","AddrType":1,"Company":""},"OrderNum":"EX145719630","Parcels":[{"TrackingNumber":"1Z7A03740311821113","Weight":7.83,"Length":39.37,"Width":14.57,"Height":9.45,"Insure":0,"Picking":null,"FbaRef":null,"Quantity":0,"LabelUrl":"https://www.ex17usa.com/UploadFiles/UD/2020-12-02/90c7e355-67c5-4223-b077-3d6bf1d408cb.pdf","AdditionalFee":0}],"Signature":false,"LabelSize":0,"Remark":null,"Remark1":null,"Remark2":null,"Zone":null,"WeightUnitType":0,"AutoChangeService":false,"ShippingDate":null,"ItemList":null,"WarehouseCode":null,"CarrierType":""}}`
	var mapParam = make(map[string]interface{})
	mapParam["Data"] = params

	content, err1 := s.postRequest(Conf.GetApiShip(), mapParam)
	if err1 != nil {
		fmt.Println("Client Post Request Error", err1)
		return
	}

	// Result
	//shipReply = new(Entry.ShipReply)
	_, err = Entry.ResponseResult(content, &shipReply)

	if err != nil {
		fmt.Println("Response to Struct Error", err)
		return
	}

	return
}

func (s KinghoodService)GetCancel(orderNum string) bool {
	mOrderNum := make(map[string]interface{})
	mOrderNum["OrderNum"] = orderNum
	content, err1 := s.postRequest(Conf.GetApiCancel(), mOrderNum)
	if err1 != nil {
		fmt.Println("Client Post Request Error", err1)
		return false
	}
	//content := `{"Code":"200","Message":null,"Data":null}`
	reply, err := Entry.ResponseResult(content, nil)
	if err != nil {
		fmt.Println("Response Error", err)
	}
	if reply.Code == "200" {
		return true
	}
	return false
}