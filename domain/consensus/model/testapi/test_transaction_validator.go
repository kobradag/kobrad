package testapi

import (
	"github.com/kobradag/kobrad/domain/consensus/model"
	"github.com/kobradag/kobrad/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
