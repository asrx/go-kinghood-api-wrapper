package Entry

type Parcels struct {
	TrackingNumber		string `json:"TrackingNumber,omitempty"`

	Weight		float64 `json:"Weight"`
	Length		float64 `json:"Length"`
	Width		float64 `json:"Width"`
	Height		float64 `json:"Height"`
	Insure		float64 `json:"Insure,omitempty"`
	//Picking	string `json:"Picking,omitempty"`
	//FbaRef	string `json:"FbaRef,omitempty"`
	Quantity	float64 `json:"Quantity,omitempty"`
	LabelUrl	string `json:"LabelUrl,omitempty"`
	AdditionalFee	float64 `json:"AdditionalFee,omitempty"`
}
