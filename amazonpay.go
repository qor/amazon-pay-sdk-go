package amazonpay

// MWSServiceURLs mwsservice urls
var MWSServiceURLs = map[string]string{
	"eu": "mws-eu.amazonservices.com",
	"na": "mws.amazonservices.com",
	"jp": "mws.amazonservices.jp",
}

var (
	// OAuthEndpoint oauth endpoint
	OAuthEndpoint = "https://api.amazon.com"
	// OAuthSandboxEndpoint oauth endpoint for sandbox env
	OAuthSandboxEndpoint = "https://api.sandbox.amazon.com"
)

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
	MerchantID    string
	AccessKey     string
	SecretKey     string
	Sandbox       bool
	Region        string
	CurrencyCode  string
	Endpoint      string
	OAuthEndpoint string
	ModePath      string
	APIVersion    string
}

// New initialize amazon pay
func New(config *Config) *AmazonPay {
	if config == nil {
		config = &Config{}
	}

	if region, ok := RegionMappings[config.Region]; ok {
		config.Region = region
	}

	if config.Region == "" {
		config.Region = "na"
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

	if config.OAuthEndpoint == "" {
		if config.Sandbox {
			config.OAuthEndpoint = OAuthSandboxEndpoint
		} else {
			config.OAuthEndpoint = OAuthEndpoint
		}
	}

	return &AmazonPay{
		Config: config,
	}
}
