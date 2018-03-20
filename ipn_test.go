package amazonpay

import (
	"fmt"
	"testing"
)

func TestIPN(t *testing.T) {
	msg := `{
    Type : Notification,
    MessageId : xxxxxx-xxxx-xxxx-xxxx-xxxxxEXAMPLE,
    TopicArn : arn:aws:sns:EXAMPLE:11111EXAMPLE:TestTopic,
    Message :
    {
        NotificationReferenceId: 1111111-1111-11111-1111-11111EXAMPLE,
        NotificationType: PaymentRefund,
        SellerId: A3G30QK8ZESA0D,
        ReleaseEnvironment: Sandbox,
        Version: 2013-01-01,
        NotificationData:<?xml version="1.0" encoding="UTF-8"?>
        <RefundNotification xmlns="https://mws.amazonservices.com/ipn/OffAmazonPayments/2013-01-01">
            <RefundDetails>
                <AmazonRefundId>P01-0000000-0000000-000000</AmazonRefundId>
                <RefundReferenceId>P01-0000000-0000000-000000</RefundReferenceId>
                <RefundType>SellerInitiated</RefundType>
                <RefundAmount>
                    <Amount>0.0</Amount>
                    <CurrencyCode>USD</CurrencyCode>
                </RefundAmount>
                <FeeRefunded>
                    <Amount>0.0</Amount>
                    <CurrencyCode>USD</CurrencyCode>
                </FeeRefunded>
                <CreationTimestamp>2013-01-01T01:01:01.001Z</CreationTimestamp>
                <RefundStatus>
                    <State>Completed</State>
                    <LastUpdateTimestamp>2013-01-01T01:01:01.001Z</LastUpdateTimestamp>
                        <ReasonCode>None</ReasonCode>
                </RefundStatus>
                <SoftDescriptor>AMZ*softDescriptor</SoftDescriptor>
            </RefundDetails>
        </RefundNotification>,
        Timestamp:2013-01-01T01:01:01Z
    }
    Timestamp : 2013-01-01T01:01:001Z,
    SignatureVersion : 1,
    Signature : rkne..9=kOUhF,
    SigningCertURL : https://sns.EXAMPLE.amazonaws.com/SimpleNotificationService-aaaaaabbbbbbccccccEXAMPLE.pem,
    UnsubscribeURL : https://sns.EXAMPLE.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:EXAMPLE:11111EXAMPLE:TestTopic:GUID
	}`

	fmt.Println(msg)
}

func TestPem(t *testing.T) {
	cert := getCert(&IPN{SigningCertURL: "https://sns.us-west-2.amazonaws.com/SimpleNotificationService-433026a4050d206028891664da859041.pem"})

	if cert == nil {
		t.Error("No Error should happen when download cert")
	} else if !verifyCertSubject(cert) {
		t.Error("Cert's subject should be correct")
	}
}
