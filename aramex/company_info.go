package aramex

// CompanyInfo is embedded in the scan event payload
type CompanyInfo struct {
	ContactName string `json:"contactName"`
	Company     string `json:"company"`

	// At some point, you must ask yourself if there's an easier way to do this.
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	Address4 string `json:"address4"`
	Address5 string `json:"address5"`
	Address6 string `json:"address6"`
	Address7 string `json:"address7"`
	Address8 string `json:"address8"`

	Comment string `json:"comment"`
}
