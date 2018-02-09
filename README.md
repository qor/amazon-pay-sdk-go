# Amazon Pay Go SDK

Amazon Pay SDK for Golang (WIP)

## Usage

```go
func main() {
  var client = amazonpay.New(&Config{
    MerchantID: "my merchant id",
    AccessKey:  "my access key",
    SecretKey:  "my secret key",
    Sandbox:    true,
    Region:     "jp",
  })

  // Set order details such as order amount and explanation in Order Reference
  client.SetOrderReferenceDetails(orderReferenceID, amazonpay.OrderReferenceAttributes) (amazonpay.OrderReferenceAttributes, error)

  // confirm order details
  client.ConfirmOrderReference(orderReferenceID) error

  // Returns the details and current state of the Order Reference object
  client.GetOrderReferenceDetails(orderReferenceID, addressToken) (amazonpay.OrderReferenceDetails, error)

  // Process secures the funds specified for the payment method stored in the Order Reference
  client.Authorize(orderReferenceID, transactionID, amount, amazonpay.AuthorizeInput) (amazonpay.AuthorizationDetails, error)

  // Returns the total authorized amount for authorization status and authorization
  client.GetAuthorizationDetails(authorizationID) (amazonpay.AuthorizationDetails, error)

  // Request funds from the authorized payment method
  client.Capture(authorizationID, transactionID, captureAmount, amazonpay.CaptureInput) (amazonpay.CaptureDetails, error)

  // Returns the detailed sales request status and the total amount refunded by sales request
  client.GetCaptureDetails(captureID) (amazonpay.CaptureDetails, error)

  // Complete order reference and will not be able to generate a new authorization from this Order Reference
  client.CloseOrderReference(orderReferenceID, closureReason) error

  // Refund refund the funds requested
  client.Refund(captureID, transactionID, refundAmount, amazonpay.RefundInput) (amazonpay.RefundDetails, error)

  // Get refund details
  client.GetRefundDetails(refundID) (amazonpay.RefundDetails, error)
}
```
