package amazonpay

// AmazonPay amazon pay
type AmazonPay struct {
	*Config
}

// Config Amazon Pay Config
type Config struct {
	MerchantID   string
	AccessKey    string
	SecretKey    string
	Sandbox      bool
	Region       string
	CurrencyCode string
}

// New initialize amazon pay
func New(config *Config) *AmazonPay {
	return &AmazonPay{
		Config: config,
	}
}
