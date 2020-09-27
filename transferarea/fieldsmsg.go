package transferarea

//Campos estrutura de campos da mensagem ISO

type Campos struct {
	Pan            string
	ProcessingCode string
	TransAmount    string
	MerchantType   int
	PointOfService string
}
