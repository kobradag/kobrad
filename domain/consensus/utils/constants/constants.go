package constants

import "math"

var (
	// BlockVersion represents the current block
	// 1 Pyrinhash
	// 2 Kodahash
	BlockVersion uint16 = 1
)

const (
	DevFee        = 2
	DevFeeMin     = 1
	DevFeeAddress = "kobra:qp8eyg0rs0jvggygh8rk7y8ux5e3wyq03w36cewcrz7542xdgumh7sr5x77ma"
	// MaxTransactionVersion is the current latest supported transaction version.
	MaxTransactionVersion uint16 = 0

	// MaxScriptPublicKeyVersion is the current latest supported public key script version.
	MaxScriptPublicKeyVersion uint16 = 0

	// LeorPerKobra is the number of leor in one kobra (1 KODA).
	LeorPerKobra = 100_000_000

	// MaxLeor is the maximum transaction amount allowed in leor.
	MaxLeor = uint64(1_000_000_000 * LeorPerKobra)

	// MaxTxInSequenceNum is the maximum sequence number the sequence field
	// of a transaction input can be.
	MaxTxInSequenceNum uint64 = math.MaxUint64

	// SequenceLockTimeDisabled is a flag that if set on a transaction
	// input's sequence number, the sequence number will not be interpreted
	// as a relative locktime.
	SequenceLockTimeDisabled uint64 = 1 << 63

	// SequenceLockTimeMask is a mask that extracts the relative locktime
	// when masked against the transaction input sequence number.
	SequenceLockTimeMask uint64 = 0x00000000ffffffff

	// LockTimeThreshold is the number below which a lock time is
	// interpreted to be a DAA score.
	LockTimeThreshold = 5e11 // Tue Nov 5 00:53:20 1985 UTC

	// UnacceptedDAAScore is used to for UTXOEntries that were created by transactions in the mempool, or otherwise
	// not-yet-accepted transactions.
	UnacceptedDAAScore = math.MaxUint64
)
