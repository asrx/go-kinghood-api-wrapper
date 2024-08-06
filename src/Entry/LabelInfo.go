package Entry

type LabelInfo struct {
	Code    string      `json:"Code"`
	Message interface{} `json:"Message"`
	Data    []struct {
		ShipmentId            string      `json:"ShipmentId"`
		LabelReady            bool        `json:"LabelReady"`
		LabelUrl              *string     `json:"LabelUrl"`
		MasterTrackingNumber  *string     `json:"MasterTrackingNumber"`
		TrackingNumbers       []string    `json:"TrackingNumbers"`
		MasterTrackingNumbers interface{} `json:"MasterTrackingNumbers"`
		Parcels               []struct {
			TrackingNumber string      `json:"TrackingNumber"`
			Weight         float64     `json:"Weight"`
			Length         float64     `json:"Length"`
			Width          float64     `json:"Width"`
			Height         float64     `json:"Height"`
			Insure         float64     `json:"Insure"`
			Picking        interface{} `json:"Picking"`
			FbaRef         interface{} `json:"FbaRef"`
			Quantity       int         `json:"Quantity"`
			LabelUrl       interface{} `json:"LabelUrl"`
			AdditionalFee  int         `json:"AdditionalFee"`
		} `json:"Parcels"`
		OrderNum            string      `json:"OrderNum"`
		LabelStatus         int         `json:"LabelStatus"`
		LabelStatusTxt      string      `json:"LabelStatusTxt"`
		TN2                 interface{} `json:"TN2"`
		Carrier             string      `json:"Carrier"`
		CarrierServiceLevel string      `json:"CarrierServiceLevel"`
		Weight              int         `json:"Weight"`
		ShippingPrice       float64     `json:"ShippingPrice"`
		ChargeDetail        interface{} `json:"ChargeDetail"`
	} `json:"Data"`
}
