package Entry

type OrderReply struct {
	EcOrder 			interface{}
	FreightCalss		interface{}
	ShipmentId			interface{}
	LabelReady			bool
	LabelUrl			string
	ShippingWeight		float64
	FactWeight			interface{}
	ShippingPrice		float64
	TrackingNumbers		[]string
	TN2					interface{}
	Carrier				string
	CarrierServiceLevel string
	ShipFrom			interface{}
	ShipTo				interface{}
	OrderNum			string
	Parcels 			[]interface{}
	Signature			interface{}
	LabelSize			interface{}
	Remark				interface{}
	Remark1				interface{}
	Remark2				interface{}
	Zone				interface{}
	WeightUnitType		interface{}
	AutoChangeService	interface{}
	ShippingDate		interface{}
	ItemList			interface{}
}