package Entry

type Address struct {
	PostalCode  string `json:"PostalCode"`
	Name        string `json:"Name"`
	Company     string `json:"Company"`
	Address1    string `json:"Address1"`
	Address2    string `json:"Address2"`
	State       string `json:"State"`
	City        string `json:"City"`
	Country     string `json:"Country"`
	PhoneNumber string `json:"PhoneNumber"`
	// 0:自动判断，1：商业，2：住宅
	AddrType uint `json:"AddrType"`
}
