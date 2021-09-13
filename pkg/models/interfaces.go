package models

type Block interface {
	GetTimestamp() int64
	GetPrevHash() []byte
	GetData() []byte
	GetHash() []byte
}
