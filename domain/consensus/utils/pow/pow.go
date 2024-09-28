package pow

import (
	"github.com/kobradag/kobrad/domain/consensus/model/externalapi"
	"github.com/kobradag/kobrad/domain/consensus/utils/consensushashing"
	"github.com/kobradag/kobrad/domain/consensus/utils/hashes"
	"github.com/kobradag/kobrad/domain/consensus/utils/serialization"
	"github.com/kobradag/kobrad/util/difficulty"
	"math/big"
	"github.com/pkg/errors"
)

// State is an intermediate data structure with pre-computed values to speed up mining.
type State struct {
	mat        matrix
	Timestamp  int64
	Nonce      uint64
	Target     big.Int
	prePowHash externalapi.DomainHash
}

// NewState creates a new state with pre-computed values to speed up mining
// It takes the target from the Bits field
func NewState(header externalapi.MutableBlockHeader) *State {
	target := difficulty.CompactToBig(header.Bits())
	// Zero out the time and nonce.
	timestamp, nonce := header.TimeInMilliseconds(), header.Nonce()
	header.SetTimeInMilliseconds(0)
	header.SetNonce(0)
	prePowHash := consensushashing.HeaderHash(header)
	header.SetTimeInMilliseconds(timestamp)
	header.SetNonce(nonce)

	return &State{
		Target:     *target,
		prePowHash: *prePowHash,
		mat:        *generateMatrix(prePowHash),
		Timestamp:  timestamp,
		Nonce:      nonce,
	}
}

// Add HeavyHashWithTransition function to the matrix type
func (mat *matrix) HeavyHashWithTransition(hash *externalapi.DomainHash, blockDAAScore uint64) *externalapi.DomainHash {
    // If blockDAAScore is greater than or equal to the transition threshold, use the new method
    if blockDAAScore >= transitionDAAScore {
        return mat.newHeavyHash(hash)
    }
    // Otherwise, use the old method
    return mat.oldHeavyHash(hash)
}

// CalculateProofOfWorkValue hashes the internal header and returns its big.Int value
// This function now checks the DAA score to determine whether to use the old or new HeavyHash method
func (state *State) CalculateProofOfWorkValue() *big.Int {
	// PRE_POW_HASH || TIME || 32 zero byte padding || NONCE
	writer := hashes.PoWHashWriter()
	writer.InfallibleWrite(state.prePowHash.ByteSlice())
	err := serialization.WriteElement(writer, state.Timestamp)
	if err != nil {
		panic(errors.Wrap(err, "this should never happen. Hash digest should never return an error"))
	}
	zeroes := [32]byte{}
	writer.InfallibleWrite(zeroes[:])
	err = serialization.WriteElement(writer, state.Nonce)
	if err != nil {
		panic(errors.Wrap(err, "this should never happen. Hash digest should never return an error"))
	}
	powHash := writer.Finalize()

	// Set DAA Score using getDAAScore method
	blockDAAScore := state.getDAAScore()

	// Call the transition-based HeavyHash function.
	heavyHash := state.mat.HeavyHashWithTransition(powHash, blockDAAScore)

	return toBig(heavyHash)
}

// IncrementNonce increments the nonce in State by 1
func (state *State) IncrementNonce() {
	state.Nonce++
}

// CheckProofOfWork checks if the block has a valid PoW according to the provided target
// It does not check if the difficulty itself is valid or less than the maximum for the appropriate network
func (state *State) CheckProofOfWork() bool {
	// The block pow must be less than the claimed target
	powNum := state.CalculateProofOfWorkValue()

	// The block hash must be less or equal than the claimed target.
	return powNum.Cmp(&state.Target) <= 0
}

// CheckProofOfWorkByBits checks if the block has a valid PoW according to its Bits field
// It does not check if the difficulty itself is valid or less than the maximum for the appropriate network
func CheckProofOfWorkByBits(header externalapi.MutableBlockHeader) bool {
	return NewState(header).CheckProofOfWork()
}

// ToBig converts a externalapi.DomainHash into a big.Int treated as a little endian string.
func toBig(hash *externalapi.DomainHash) *big.Int {
	// We treat the Hash as little-endian for PoW purposes, but the big package wants the bytes in big-endian, so reverse them.
	buf := hash.ByteSlice()
	blen := len(buf)
	for i := 0; i < blen/2; i++ {
		buf[i], buf[blen-1-i] = buf[blen-1-i], buf[i]
	}

	return new(big.Int).SetBytes(buf)
}

// BlockLevel returns the block level of the given header.
func BlockLevel(header externalapi.BlockHeader, maxBlockLevel int) int {
	// Genesis is defined to be the root of all blocks at all levels, so we define it to be the maximal
	// block level.
	if len(header.DirectParents()) == 0 {
		return maxBlockLevel
	}

	proofOfWorkValue := NewState(header.ToMutable()).CalculateProofOfWorkValue()
	level := maxBlockLevel - proofOfWorkValue.BitLen()
	// If the block has a level lower than genesis make it zero.
	if level < 0 {
		level = 0
	}
	return level
}

// Manually set DAA Score for now, replace this with real logic
func (state *State) getDAAScore() uint64 {
	// Placeholder: Replace this with actual logic to retrieve DAA score from the header.
	return transitionDAAScore - 1 // Example: Set below the threshold for now
}
