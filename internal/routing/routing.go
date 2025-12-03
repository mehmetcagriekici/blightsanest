package routing

// exchange names
const (
	CryptoExchange  = "crypto_topic"
	ClientExchange  = "client_topic"
)

// routing keys
const (
        BlightCrypto       = "blight_crypto"
	BlightClientCrypto = "blight_client_crypto"
)

// exchange types
const (
        BlightTopic  = "topic"
	BlightDirect = "direct"
)

// queue names
const (
        CryptoGet       = "crypto_get"
	CryptoClientGet = "crypto_client_get"
)

// queue types
type QueueType string
const (
        BlightDurable   QueueType = "durable"
	BlightTransient QueueType = "transient"
)