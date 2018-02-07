package amazonpay

// SetOrderReferenceDetails set order details such as order amount and explanation in Order Reference
func (amazonPay *AmazonPay) SetOrderReferenceDetails(orderReferenceID string, attrs OrderReferenceAttributes) OrderReferenceAttributes {
	return OrderReferenceAttributes{}
}

// ConfirmOrderReference confirm order details
func (amazonPay *AmazonPay) ConfirmOrderReference(orderReferenceID string) error {
	return nil
}

// GetOrderReferenceDetails Returns the details and current state of the Order Reference object.
func (amazonPay *AmazonPay) GetOrderReferenceDetails(orderReferenceID string, addressToken string) OrderReferenceDetails {
	return OrderReferenceDetails{}
}

// AuthorizeInput authorize input struct
type AuthorizeInput struct {
	SellerAuthorizationNote string
	TransactionTimedOut     uint
	CaptureNow              bool
	SoftDecriptor           string
}

// Authorize process secures the funds specified for the payment method stored in the Order Reference.
func (amazonPay *AmazonPay) Authorize(orderReferenceID string, authReferenceID string, amount Price) AuthorizationDetails {
	return AuthorizationDetails{}
}
