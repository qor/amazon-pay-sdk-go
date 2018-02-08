package amazonpay

import (
	"fmt"
	"testing"
)

var Client = New(&Config{
	MerchantID: "test",
	AccessKey:  "test",
	SecretKey:  "test",
	Sandbox:    true,
})

func TestSetOrderReferenceDetails(t *testing.T) {
	data := Client.SetOrderReferenceDetails("orderReferenceID", OrderReferenceAttributes{
		OrderTotal: OrderTotal{
			CurrencyCode: "usd",
			Amount:       "100",
		},
		PlatformID: "test",
		SellerNote: "test",
		SellerOrderAttributes: SellerOrderAttributes{
			SellerOrderID:     "orderid",
			StoreName:         "test",
			CustomInformation: "test",
		},
	})

	fmt.Printf("%#v\n", data)
}
