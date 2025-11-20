package routing

// exchange names
const (
	CryptoExchange  = "crypto_topic"
)

// exchange keys
const (
        BlightCrypto = "crypto"
)

// exchange types
const (
        BlightTopic  = "topic"
	BlightDirect = "direct"
)

// queue names
const (
        CryptoGet = "crypto_get"
)

// queue types
type QueueType string
const (
        BlightDurable   QueueType = "durable"
	BlightTransient QueueType = "transient"
)