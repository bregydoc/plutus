package plutus

// SaleRepresentation is a representation of one sale
type SaleRepresentation struct {
	Data        []byte
	Name        string
	Extension   string
	ContentType string
}

// DeliveryChannel a delivery channel is a way to represent and send a voucher of yourtransaction.
// examples of delivery channel are: SMTP, SMS, ThermalPrint, etc...
type DeliveryChannel interface {
	Name() string
	SendSaleReceipt(from *Company, sale *Sale, metadata ...map[string]interface{}) error
	GetSaleRepresentation(from *Company, sale *Sale, metadata ...map[string]interface{}) (*SaleRepresentation, error)
}
