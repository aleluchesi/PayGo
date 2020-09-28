package transferarea

// Iso8585 fields msg ISO

type Iso8583 struct {
	Pan            string
	ProcessingCode string
	TransAmount    string
	MerchantType   int
	PointOfService string
}
