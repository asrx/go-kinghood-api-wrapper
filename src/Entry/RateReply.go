package Entry

type RateReply struct {
	RateId				string 		`json:"RateId"`
	Carrier				string		`json:"Carrier"`
	CarrierServiceLevel string		`json:"CarrierServiceLevel"`
	CarrierServiceName	string		`json:"CarrierServiceName"`
	ChargeWeight		float64 	`json:"ChargeWeight"`
	TotalCharge			float64 	`json:"TotalCharge"`
	PackageType			interface{} `json:"PackageType"`
	Message				interface{} `json:"Message"`
	Zone				interface{} `json:"Zone"`
	IsResidential		bool		`json:"IsResidential"`
	PublishCharge		int			`json:"PublishCharge"`
	IsPublishChargeSpecified		bool		`json:"IsPublishChargeSpecified"`
}