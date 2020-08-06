package aramex

// Scan is the object representing a parcel scan event
type Scan struct {
	Type              string
	Courier           string
	Description       string
	Date              string
	RealDateTime      string
	Name              string
	Franchise         string
	Status            string
	StatusDescription string
	CompanyInfo       CompanyInfo
	UploadDate        string
	Signature         string
	// ParcelConnectAgent: Ommitted due to insufficient information during early tests
}
