package amazonpay

// SetOrderReferenceDetails set order details such as order amount and explanation in Order Reference
func (amazonPay *AmazonPay) SetOrderReferenceDetails(orderReferenceID string, attrs OrderReferenceAttributes) (result OrderReferenceAttributes, err error) {
	var params = Params{
		"Action":                 "SetOrderReferenceDetails",
		"AmazonOrderReferenceId": orderReferenceID,
	}

	err = updateParams(&params, "OrderReferenceAttributes", attrs)

	if err == nil {
		err = amazonPay.Post(params)
	}

	return result, err
}

// ConfirmOrderReference confirm order details
func (amazonPay *AmazonPay) ConfirmOrderReference(orderReferenceID string) error {
	var params = Params{
		"Action":                 "ConfirmOrderReference",
		"AmazonOrderReferenceId": orderReferenceID,
	}

	return amazonPay.Post(params)
}

// GetOrderReferenceDetails Returns the details and current state of the Order Reference object.
func (amazonPay *AmazonPay) GetOrderReferenceDetails(orderReferenceID string, addressToken string) (result OrderReferenceDetails, err error) {
	var params = Params{
		"Action":                 "GetOrderReferenceDetails",
		"AmazonOrderReferenceId": orderReferenceID,
	}

	err = amazonPay.Post(params)

	return result, err
}

// AuthorizeInput authorize input struct
type AuthorizeInput struct {
	SellerAuthorizationNote string
	TransactionTimedOut     uint
	CaptureNow              bool
	SoftDecriptor           string
}

// Authorize process secures the funds specified for the payment method stored in the Order Reference.
func (amazonPay *AmazonPay) Authorize(orderReferenceID string, transactionID string, amount Price, input AuthorizeInput) (result AuthorizationDetails, err error) {
	var params = Params{
		"Action":                   "Authorize",
		"AmazonOrderReferenceId":   orderReferenceID,
		"AuthorizationReferenceId": transactionID,
	}

	updateParams(&params, "AuthorizationAmount", amount)
	updateParams(&params, "", input)

	err = amazonPay.Post(params)

	return result, err
}

// GetAuthorizationDetails returns the total authorized amount for authorization status and authorization.
func (amazonPay *AmazonPay) GetAuthorizationDetails(authorizationID string) (result AuthorizationDetails, err error) {
	var params = Params{
		"Action":                "GetAuthorizationDetails",
		"AmazonAuthorizationId": authorizationID,
	}

	err = amazonPay.Post(params)

	return result, err
}

// CaptureInput capture input struct
type CaptureInput struct {
	SellerCaptureNote string
	SoftDecriptor     string
}

// Capture request funds from the authorized payment method.
func (amazonPay *AmazonPay) Capture(authorizationID string, transactionID string, captureAmount Price, input CaptureInput) (result CaptureDetails, err error) {
	var params = Params{
		"Action":                "Capture",
		"AmazonAuthorizationId": authorizationID,
		"CaptureReferenceId":    transactionID,
	}

	updateParams(&params, "CaptureAmount", captureAmount)
	updateParams(&params, "", input)

	err = amazonPay.Post(params)

	return result, err
}

// GetCaptureDetails returns the detailed sales request status and the total amount refunded by sales request.
func (amazonPay *AmazonPay) GetCaptureDetails(captureID string) (result CaptureDetails, err error) {
	var params = Params{
		"Action":          "GetCaptureDetails",
		"AmazonCaptureId": captureID,
	}

	err = amazonPay.Post(params)
	return result, err
}

// CloseOrderReference complete order reference and will not be able to generate a new authorization from this Order Reference.
func (amazonPay *AmazonPay) CloseOrderReference(orderReferenceID string, closureReason string) error {
	var params = Params{
		"Action":                 "CloseOrderReference",
		"AmazonOrderReferenceId": orderReferenceID,
		"ClosureReason":          closureReason,
	}

	return amazonPay.Post(params)
}

// RefundInput refund input struct
type RefundInput struct {
	SellerRefundNote string
	SoftDescriptor   string
}

// Refund refund the funds requested
func (amazonPay *AmazonPay) Refund(captureID string, transactionID string, refundAmount Price, input RefundInput) (result RefundDetails, err error) {
	var params = Params{
		"Action":            "Refund",
		"AmazonCaptureId":   captureID,
		"RefundReferenceId": transactionID,
	}

	updateParams(&params, "RefundAmount", refundAmount)
	updateParams(&params, "", input)

	err = amazonPay.Post(params)

	return result, err
}

// GetRefundDetails get refund details
func (amazonPay *AmazonPay) GetRefundDetails(refundID string) (result RefundDetails, err error) {
	var params = Params{
		"Action":         "GetRefundDetails",
		"AmazonRefundId": refundID,
	}

	err = amazonPay.Post(params)
	return result, err
}
