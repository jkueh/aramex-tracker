package aramex

// TrackingResult is the main result object
type TrackingResult struct {
	LabelNumber         string
	Scans               []Scan
	Signature           string
	DistributedTo       string
	DistributedDate     string
	Reference           string
	OriginalLabelNumber string

	// WHY!?
	IsOnforward       string
	IsNZPostOnforward bool

	OnforwardLabelNumber string
	CallingCard          string
	CountryCode          string

	// I'm assuming this (like most) of the payload, is a string
	CallingCardLabelNumbers []string

	HasDScan              bool
	LastScanDate          string
	WasScannedLast24Hours bool
}
