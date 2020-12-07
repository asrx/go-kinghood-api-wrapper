package Conf

const (
	// 接口地址
	_URL_PRODUCTION = "https://www.ex17usa.com/Api"

	// 余额
	_API_BALANCE = "/UserInfo.ashx"

	// 查可用服务
	_API_CARRIERS = "/Express/GetCarrierList.ashx"

	// 查费率
	_API_RATE = "/Express/Rates.ashx"

	// 运单下单
	_API_SHIP = "/Express/CreateLabel2.ashx"

	// 标签明细
	_API_LABEL_INFO = "/Express/LabelsInfo.ashx"

	// 取消标签
	_API_LABEL_CANCEL = "/Express/LabelsCancel.ashx"
)

func GetApiBalance() string {
	return _URL_PRODUCTION + _API_BALANCE
}
func GetApiCarriers() string {
	return _URL_PRODUCTION + _API_CARRIERS
}
func GetApiRate() string {
	return _URL_PRODUCTION + _API_RATE
}
func GetApiShip() string {
	return _URL_PRODUCTION + _API_SHIP
}
func GetApiLabelInfo() string {
	return _URL_PRODUCTION + _API_LABEL_INFO
}
func GetApiCancel() string {
	return _URL_PRODUCTION + _API_LABEL_CANCEL
}

const (
	WeightUnitType_LBS_IN = 0
	WeightUnitType_KGS_CN = 1
	WeightUnitType_OZ_IN = 2
)

const (
	CARRIER_UPS_GROUND6 = "UPS Ground 6"
	//CARRIER_UPS_GROUND3 = "UPS Ground3"
	SERVICE_UPS_GROUND = "UPS Ground"


	CARRIER_UPS_GFP = "UPS GFP"
	SERVICE_UPS_GFP = "ups_ground_freight"
)
