package src

import (
	"errors"
	"fmt"

	"github.com/idoubi/goz"

	"github.com/asrx/go-kinghood-api-wrapper/src/Entry"
)

type KinghoodService struct {
	key     string
	BaseURI string
	// 余额
	_API_BALANCE string

	// 查可用服务
	_API_CARRIERS string

	// 查费率
	_API_RATE string

	// 运单下单
	_API_SHIP string

	// 标签明细
	_API_LABEL_INFO string

	// 取消标签
	_API_LABEL_CANCEL string
}

func NewKinghoodService(key, BaseURI string) *KinghoodService {
	return &KinghoodService{
		key:               key,
		BaseURI:           BaseURI,
		_API_BALANCE:      "/Api/UserInfo.ashx",
		_API_CARRIERS:     "/Api/Express/GetCarrierList.ashx",
		_API_RATE:         "/Api/Express/Rates.ashx",
		_API_SHIP:         "/Api/Express/CreateLabel2.ashx",
		_API_LABEL_INFO:   "/Api/Express/LabelsInfo.ashx",
		_API_LABEL_CANCEL: "/Api/Express/LabelsCancel.ashx",
	}
}

func (s KinghoodService) mergeKey(mapParam map[string]interface{}) {
	mapParam["Key"] = s.key
}

func (s KinghoodService) postRequest(url string, mapParam map[string]interface{}) (content string, err error) {

	s.mergeKey(mapParam)

	client := goz.NewClient()
	resp, err := client.Post(url, goz.Options{
		Headers: map[string]interface{}{
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
func (s KinghoodService) GetBalance() (money float64, err error) {
	var mapParam = make(map[string]interface{})
	content, err1 := s.postRequest(
		s.BaseURI+s._API_BALANCE, mapParam)

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
func (s KinghoodService) GetCarriers() (content string, err error) {
	var mapParam = make(map[string]interface{})
	content, err1 := s.postRequest(s.BaseURI+s._API_CARRIERS, mapParam)
	if err1 != nil {
		err = err1
		fmt.Println("Client Post Request Error", err1)
		return
	}
	return
}

// Rate
func (s KinghoodService) GetRate(params string) (rateReply *Entry.RateReply, err error) {
	// content := `{"Code":"200","Message":"","Data":[{"RateId":"0","Carrier":"UPS Ground 6","CarrierServiceLevel":"UPS Ground","CarrierServiceName":"UPS Ground","ChargeWeight":40,"TotalCharge":32.40,"PackageType":null,"Message":null,"Zone":null,"IsResidential":false,"PublishCharge":0,"IsPublishChargeSpecified":false}]}`
	var mapParam = make(map[string]interface{})
	mapParam["Data"] = params

	content, err1 := s.postRequest(s.BaseURI+s._API_RATE, mapParam)
	if err1 != nil {
		fmt.Println("Client Post Request Error", err1)
		return
	}
	// Result
	list := []*Entry.RateReply{}
	// reply := Entry.Reply{}

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
	rateReply = list[0]
	return
}

// Ship
func (s KinghoodService) GetShip(params string) (shipReply *Entry.ShipReply, err error) {
	// content := `{"Code":"200","Message":null,"Data":{"EcOrder":null,"FreightCalss":null,"ShipmentId":null,"LabelReady":true,"LabelUrl":"https://www.ex17usa.com/UploadFiles/UD/2020-12-02/90c7e355-67c5-4223-b077-3d6bf1d408cb.pdf","ShippingWeight":8,"FactWeight":0,"ShippingPrice":6.48,"TrackingNumbers":["1Z7A03740311821113"],"TN2":null,"Carrier":"UPS Ground 6","CarrierServiceLevel":"UPS Ground","ShipFrom":{"PostalCode":"92879","Name":"micheal","Address1":"2565 SAMPSON AVENUE","Address2":"","State":"CA","City":"CORONA","Country":"US","CountryCode":null,"PhoneNumber":"09512680000","AddrType":1,"Company":""},"ShipTo":{"PostalCode":"92551","Name":"Donovan","Address1":"24300 Nandina Ave","Address2":null,"State":"CA","City":"Moreno Valley","Country":"US","CountryCode":null,"PhoneNumber":"13474483190","AddrType":1,"Company":""},"OrderNum":"EX145719630","Parcels":[{"TrackingNumber":"1Z7A03740311821113","Weight":7.83,"Length":39.37,"Width":14.57,"Height":9.45,"Insure":0,"Picking":null,"FbaRef":null,"Quantity":0,"LabelUrl":"https://www.ex17usa.com/UploadFiles/UD/2020-12-02/90c7e355-67c5-4223-b077-3d6bf1d408cb.pdf","AdditionalFee":0}],"Signature":false,"LabelSize":0,"Remark":null,"Remark1":null,"Remark2":null,"Zone":null,"WeightUnitType":0,"AutoChangeService":false,"ShippingDate":null,"ItemList":null,"WarehouseCode":null,"CarrierType":""}}`
	var mapParam = make(map[string]interface{})
	mapParam["Data"] = params

	content, err1 := s.postRequest(s.BaseURI+s._API_SHIP, mapParam)
	if err1 != nil {
		fmt.Println("Client Post Request Error", err1)
		return
	}

	// Result
	// shipReply = new(Entry.ShipReply)
	resp, err := Entry.ResponseResult(content, &shipReply)

	if err != nil {
		fmt.Println("Response to Struct Error", err)
		return
	}

	if resp.Code != "200" {
		err = errors.New(resp.Message)
		return
	}

	return
}

// Cancel
func (s KinghoodService) GetCancel(orderNum string) bool {
	mOrderNum := make(map[string]interface{})
	mOrderNum["OrderNum"] = orderNum
	content, err1 := s.postRequest(s.BaseURI+s._API_LABEL_CANCEL, mOrderNum)
	if err1 != nil {
		fmt.Println("Client Post Request Error", err1)
		return false
	}
	// content := `{"Code":"200","Message":null,"Data":null}`
	reply, err := Entry.ResponseResult(content, nil)
	if err != nil {
		fmt.Println("Response Error", err)
	}
	if reply.Code == "200" {
		return true
	}
	return false
}

// 获取延迟的订单信息
func (s KinghoodService) GetLabelInfo(params string) (labelInfo *Entry.LabelData, err error) {
	var mapParam = make(map[string]interface{})
	mapParam["Data"] = params
	content, err := s.postRequest(s.BaseURI+s._API_LABEL_INFO, mapParam)
	// content := `{"Code":"200","Message":null,"Data":[{"ShipmentId":"","LabelReady":true,"LabelUrl":"https://www.goshipping.info/UploadFiles/EY/2024-08-05/f449a1d6-987b-4a3f-8ccf-2d0f33a3fc93.pdf","MasterTrackingNumber":"777806336591","TrackingNumbers":["777806336591"],"MasterTrackingNumbers":null,"Parcels":[{"TrackingNumber":"777806336591","Weight":5.00,"Width":2.00,"Length":3.00,"Height":1.00,"Quantity":0,"Insure":0.00,"Picking":null,"FbaRef":null,"LabelUrl":null,"AdditionalFee":0}],"OrderNum":"GS116259026","LabelStatus":30,"LabelStatusTxt":"等待发货","TN2":null,"Carrier":"美西优选比价渠道","CarrierServiceLevel":"11 FedEx Ground®","Weight":5,"ShippingPrice":8.8,"ChargeDetail":null}]}`
	// content := `{"Code":"200","Message":null,"Data":[{"ShipmentId":"","LabelReady":true,"LabelUrl":"http://www2.goshipping.info","MasterTrackingNumber":"777853363421","TrackingNumbers":["777853363421"],"MasterTrackingNumbers":null,"Parcels":[{"TrackingNumber":"777853363421","Weight":4.41,"Length":1.00,"Width":1.00,"Height":1.00,"Insure":0.00,"Picking":null,"FbaRef":null,"Quantity":0,"LabelUrl":null,"AdditionalFee":0}],"OrderNum":"GS116385122","LabelStatus":1810,"LabelStatusTxt":"退款审核中","TN2":null,"Carrier":"美西比价渠道","CarrierServiceLevel":"11 FedEx Home Delivery Adult","Weight":4.41,"ShippingPrice":25.33,"ChargeDetail":null}]}`
	if err != nil {
		fmt.Println("Client Post Request Error", err)
		return
	}
	labelsData := []*Entry.LabelData{}
	resp, err := Entry.ResponseResult(content, &labelsData)
	if err != nil {
		fmt.Println("Response to Struct Error", err)
		return
	}
	if resp.Code != "200" {
		err = errors.New(resp.Message)
		return
	}

	if labelsData == nil || len(labelsData) == 0 {
		return
	}
	if labelInfo.LabelStatus == 1810 {
		return
	}
	return labelsData[0], nil
}
