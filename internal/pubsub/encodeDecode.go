package pubsub

import(
        "bytes"
        "encoding/gob"
)

func Encode[T any](val T) ([]byte, error) {
        var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	if err := enc.Encode(val); err != nil {
	        return nil, err
	}
	return network.Bytes(), nil
}

func Decode[T any](data []byte) (T, error) {
        network := bytes.NewBuffer(data)
	dec := gob.NewDecoder(network)
	
	var v T
	if err := dec.Decode(&v); err != nil {
	        var noop T
		return noop, err
	}
	return v, nil
}