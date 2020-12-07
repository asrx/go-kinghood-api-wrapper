package Entry

type ShipReply struct {
	EcOrder				interface{} `json:"EcOrder"`
	FreightCalss		interface{} `json:"FreightCalss"`
	ShipmentId			interface{} `json:"ShipmentId"`
	LabelReady			bool `json:"LabelReady"`
	LabelUrl			string `json:"LabelUrl"`
	ShippingWeight		float64 `json:"ShippingWeight"`
	FactWeight			interface{} `json:"FactWeight"`
	ShippingPrice		float64 `json:"ShippingPrice"`
	TrackingNumbers		[]string `json:"TrackingNumbers"`
	TN2					interface{} `json:"TN2"`
	Carrier				string `json:"Carrier"`
	CarrierServiceLevel	string `json:"CarrierServiceLevel"`
	ShipFrom			*Address `json:"ShipFrom"`
	ShipTo				*Address `json:"ShipTo"`
	OrderNum			interface{} `json:"OrderNum"`
	Parcels				[]*Parcels `json:"Parcels"`
	Signature			interface{} `json:"Signature"`
	LabelSize			interface{} `json:"LabelSize"`
	Remark				interface{} `json:"Remark"`
	Remark1				interface{} `json:"Remark1"`
	Remark2				interface{} `json:"Remark2"`
	Zone				interface{} `json:"Zone"`
	WeightUnitType		interface{} `json:"WeightUnitType"`
	AutoChangeService	interface{} `json:"AutoChangeService"`
	ShippingDate		interface{} `json:"ShippingDate"`
	ItemList			interface{} `json:"ItemList"`
}