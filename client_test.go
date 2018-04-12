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
	Region:     "jp",
})

func TestSetOrderReferenceDetails(t *testing.T) {
	data, err := Client.SetOrderReferenceDetails("orderReferenceID", OrderReferenceAttributes{
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
	fmt.Println(err)
}

func TestGetProfile(t *testing.T) {
	profile, err := Client.GetProfile("Atza|IwEBIHyBBHS8cgSZsYfoxWeU0gU7h3PCtfTzQeO2AU1QYAYz1vBDJwxBSJ89dVF3xffMwYUs8czy3uTGVCIU2QXxZIbD7j0ttSEAI3Dm2kqdRIRvuQYcQ-MkgrVmR4bZiFicZDYgJwql4jCX_T_A6y-UECFWFRgI7_uI5cw6iIj-iZWUDjucKgn_STtQTT8FQ04JrGZExDQquKMQIfkDu7HusxTKs23bIcsvLZLGW43QKwaK1QFdJ0qEu81QeMGCujsAunWBGpkS6OmwuSWqtBugA9HMOJ4XlitrGAZkayTAduSAx_KFyqm4swo-nZR0qIHuypENh6k8ZRJ5gwyNtmYNumyAVIVNxwpcBSPL8VQ665ctttYQg__YReWZSNBQiOqfOH881Z9YcOMnEsE9ObrczFZ1XpBBWNy4quDC_hHwBXC6ecYg5Kxx8OcB0i7zFTsAL5uOkRlCutTJyiWSjUH5UCry2Mu6QZJlKhCusjT74mIa04yaassw0Ih8n7yIaqt-omkxz18T_aX-ZEdpG8r0KvCTcE8p2RJ5gBkAD3NF6LvpHlegm3nj1ybGa4gCg--eUOqpytutPpajor1qXgYDU8anymLAacM5gxLz-JHiQlk9Rg")
	if err != nil {
		t.Errorf("No error should happen when get profile, but got %v", err)
	}

	if profile.Name == "" {
		t.Errorf("Profile should has value")
	}
}
