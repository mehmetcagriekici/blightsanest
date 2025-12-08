package routing

// exchange names
const (
	CryptoExchange = "crypto_topic"
	ClientExchange = "client_topic"
	CryptoDLX      = "crypto_dlx"
	TestExchange   = "exchange_test"
)

// routing keys
const (
        BlightCrypto       = "blight_crypto"
	BlightClientCrypto = "blight_client_crypto"
	BlightTesting      = "blight_test"
)

// exchange types
const (
        BlightTopic  = "topic"
	BlightDirect = "direct"
	BlightFanout = "fanout"
)

// queue names
const (
        CryptoGet       = "crypto_get"
	CryptoClientGet = "crypto_client_get"
	CryptoDLQ       = "crypto_dlq"
	TestQueue       = "queue_test"
)

// queue types
type QueueType string
const (
        BlightDurable   QueueType = "durable"
	BlightTransient QueueType = "transient"
)

// acknowledgement types
type AckType string
const (
        ACK          AckType = "ack"
	NACK_REQUEUE AckType = "nack_requeue"
	NACK_DISCARD AckType = "nack_discard"
)