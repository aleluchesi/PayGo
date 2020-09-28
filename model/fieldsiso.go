package model

//FieldsIso Contem campos da mensagem ISO 8583
type FieldsIso struct {
	Pan           string // Bit 2
	MerchanteTyoe int    // Bit 19
}

//SetPan atribui conteúdo do PAN (cartão)
func (f *FieldsIso) SetPan(pan string) {
	f.Pan = pan
}

//GetPan retorna conteúdo do Pan (cartão)
func (f *FieldsIso) GetPan() string {
	return f.Pan
}
