package keyvalue

type Key string
type Value []byte

type Store interface {
	Put(Key, Value)
	Get(Key) Value
	Delete(Key)
}
