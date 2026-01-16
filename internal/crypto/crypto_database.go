package crypto

import(
	"log"
	"time"
	"context"
	"encoding/json"

	"github.com/google/uuid"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

// create crypto row to the database
func CreateCryptoRow(ctx context.Context,
	qr *database.Queries,
	list []MarketData,
	cryptoKey string) error {
	
	// encode the crypto list
	encoded, err := json.Marshal(list)
	if err != nil {
		return err
	}

	params := database.CreateCryptoListParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CryptoKey: cryptoKey,
		CryptoList: json.RawMessage(encoded),
	}

	created, err := qr.CreateCryptoList(ctx, params)
	if err != nil {
		return err
	}

	log.Printf("Crypto List %s successfully saved to the database at %v with the key %s\n", cryptoKey, created.UpdatedAt, created.CryptoKey)

	return nil
}

// read crypto data from the database
func ReadCryptoRow(ctx context.Context, qr *database.Queries, cryptoKey string) ([]MarketData, error) {
	// get the row from the database
	data, err := qr.GetCryptoList(ctx, cryptoKey)
	if err != nil {
		log.Println("Couldn't fetch the crypto row from the database:")
		return nil, err
	}

	// marshal the json.RawMessage into bytes to be decoded
	m, err := data.CryptoList.MarshalJSON()
	if err != nil {
		log.Println("Couldn't marshal the json raw message:")
		return nil, err
	}

	// unmarshal the bytes
	var b []MarketData
	if err := json.Unmarshal(m, &b); err != nil {
		log.Println("Couldn't decode the marshalled json raw message:")
		return nil, err
	}

	log.Printf("Crypto List %s successfully fetched from the database.\n", data.CryptoKey)
	return b, nil
}

// update crypto row at the database
func UpdateCryptoRow(ctx context.Context,
	list []MarketData,
	qr *database.Queries,
	cryptoKey,
	newCryptoKey string) ([]MarketData, error) {
	// encode the crypto list
	encoded, err := json.Marshal(list)
	if err != nil {
		log.Println("Couldn't decode the crypto list:")
		return nil, err
	}

	// update crypto params
	params := database.UpdateCryptoListParams{
		CryptoKey: cryptoKey, // to check the match
		UpdatedAt: time.Now(),
		CryptoKey_2: newCryptoKey, // to save
		CryptoList: json.RawMessage(encoded),
	}

	// update the row on the database
	updated, err := qr.UpdateCryptoList(ctx, params)
	if err != nil {
		log.Println("Couldn't update the crypto list on the database:")
		return nil, err
	}

	log.Printf("Successfully updated the database %s with the id %s at %v\n", cryptoKey, updated.CryptoKey, updated.UpdatedAt)

	return list, nil
}

// delete crypto row from the database
func DeleteCryptoRow(ctx context.Context, qr *database.Queries, cryptoKey string) error {
	data, err := qr.DeleteCryptoList(ctx, cryptoKey)
	if err != nil {
		return err
	}

	for _, v := range data {
		log.Printf("List %s is deleted from the database.\n", v.CryptoKey)
	}

	return nil
}
