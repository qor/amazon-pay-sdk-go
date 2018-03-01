package amazonpay

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestOrderRerenceDetails(t *testing.T) {
	resp := `<GetOrderReferenceDetailsResponse xmlns="http://mws.amazonservices.com/schema/OffAmazonPayments/2013-01-01">
  <GetOrderReferenceDetailsResult>
    <OrderReferenceDetails>
      <OrderReferenceStatus>
        <State>Draft</State>
      </OrderReferenceStatus>
      <Destination>
        <DestinationType>Physical</DestinationType>
        <PhysicalDestination>
          <StateOrRegion>Tokyo</StateOrRegion>
          <City>Ota-ku</City>
          <CountryCode>JP</CountryCode>
          <PostalCode>144-8588</PostalCode>
        </PhysicalDestination>
      </Destination>
      <ExpirationTimestamp>2018-08-28T16:00:23.013Z</ExpirationTimestamp>
      <IdList/>
      <SellerOrderAttributes/>
      <OrderTotal>
        <CurrencyCode>JPY</CurrencyCode>
        <Amount>23.00</Amount>
      </OrderTotal>
      <ReleaseEnvironment>Sandbox</ReleaseEnvironment>
      <AmazonOrderReferenceId>S03-1594298-4362493</AmazonOrderReferenceId>
      <CreationTimestamp>2018-03-01T16:00:23.013Z</CreationTimestamp>
      <RequestPaymentAuthorization>false</RequestPaymentAuthorization>
    </OrderReferenceDetails>
  </GetOrderReferenceDetailsResult>
  <ResponseMetadata>
    <RequestId>ae47aef3-e030-4fdf-a0bb-87c4767e63d6</RequestId>
  </ResponseMetadata>
</GetOrderReferenceDetailsResponse>`

	detail := GetOrderReferenceDetailsResponse{}
	if err := xml.Unmarshal([]byte(resp), &detail); err != nil {
		t.Errorf("no error should get when unmarshal order reference detail, but got %v", err)
	}

	fmt.Printf("%#v", detail)
}
