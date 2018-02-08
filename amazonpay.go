package amazonpay

// MWSServiceURLs mwsservice urls
var MWSServiceURLs = map[string]string{
	"eu": "mws-eu.amazonservices.com",
	"na": "mws.amazonservices.com",
	"jp": "mws.amazonservices.jp",
}

// RegionMappings region mapping
var RegionMappings = map[string]string{
	"de": "eu",
	"uk": "eu",
	"us": "na",
	"jp": "jp",
}

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
	Endpoint     string
	ModePath     string
	APIVersion   string
}

// New initialize amazon pay
func New(config *Config) *AmazonPay {
	if config == nil {
		config = &Config{}
	}

	if config.Region == "" {
		config.Region = "us"
	}

	if config.Endpoint == "" {
		config.Endpoint = MWSServiceURLs[config.Region]
	}

	if config.APIVersion == "" {
		config.APIVersion = "2013-01-01"
	}

	if config.ModePath == "" {
		if config.Sandbox {
			config.ModePath = "OffAmazonPayments_Sandbox"
		} else {
			config.ModePath = "OffAmazonPayments"
		}
	}

	return &AmazonPay{
		Config: config,
	}
}
