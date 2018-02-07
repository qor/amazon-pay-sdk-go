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
func (amazonPay *AmazonPay) Authorize(orderReferenceID string, transactionID string, amount Price, input AuthorizeInput) AuthorizationDetails {
	return AuthorizationDetails{}
}

// GetAuthorizationDetails returns the total authorized amount for authorization status and authorization.
func (amazonPay *AmazonPay) GetAuthorizationDetails(authorizationID string) AuthorizationDetails {
	return AuthorizationDetails{}
}

// CaptureInput capture input struct
type CaptureInput struct {
	SellerCaptureNote string
	SoftDecriptor     string
}

// Capture request funds from the authorized payment method.
func (amazonPay *AmazonPay) Capture(authorizationID string, transactionID string, captureAmount Price, input CaptureInput) CaptureDetails {
	return CaptureDetails{}
}

// GetCaptureDetails returns the detailed sales request status and the total amount refunded by sales request.
func (amazonPay *AmazonPay) GetCaptureDetails(captureID string) CaptureDetails {
	return CaptureDetails{}
}

// CloseOrderReference complete order reference and will not be able to generate a new authorization from this Order Reference.
func (amazonPay *AmazonPay) CloseOrderReference(orderReferenceID string, closeReason string) error {
	return nil
}

// RefundInput refund input struct
type RefundInput struct {
	SellerRefundNote string
	SoftDescriptor   string
}

// Refund refund the funds requested
func (amazonPay *AmazonPay) Refund(captureID string, transactionID string, refundAmount Price, input RefundInput) RefundDetails {
	return RefundDetails{}
}

// GetRefundDetails get refund details
func (amazonPay *AmazonPay) GetRefundDetails(refundID string) RefundDetails {
	return RefundDetails{}
}
