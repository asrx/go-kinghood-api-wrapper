package Entry

type Shipment struct {
	Carrier 		string 			`json:"Carrier"`

	ShipFrom 		*Address	 	`json:"ShipFrom"`
	ShipTo 			*Address	 	`json:"ShipTo"`

	Parcels 		[]*Parcels	 	`json:"Parcels"`

	FreightCalss 	interface{}		`json:"FreightCalss,omitempty"`
	Signature 		interface{}		`json:"Signature,omitempty"`
	WeightUnitType 	interface{}  	`json:"WeightUnitType,omitempty"`
	CarrierServiceLevel string  	`json:"CarrierServiceLevel"`

	LabelSize		interface{}		`json:"LabelSize,omitempty"`
	Remark1			interface{} 	`json:"Remark1,omitempty"`
	Remark2			interface{} 	`json:"Remark2,omitempty"`
	EcOrder			interface{}		`json:"EcOrder,omitempty"`
}

