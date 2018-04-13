package service

type StateRepository interface {
	CallGetSyncBlockCount() (uint64, error)
	CallTokenAging(int64) error
}
