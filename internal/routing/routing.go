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

// queue types
type QueueType string
const (
        BlightDurable   QueueType = "durable"
	BlightTransient QueueType = "transient"
)