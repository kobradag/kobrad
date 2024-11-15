package pow

import (
	"github.com/kobradag/kobrad/domain/consensus/model/externalapi"
	"github.com/kobradag/kobrad/domain/consensus/utils/consensushashing"
	"github.com/kobradag/kobrad/domain/consensus/utils/hashes"
	"github.com/kobradag/kobrad/domain/consensus/utils/serialization"
	"github.com/kobradag/kobrad/util/difficulty"

	"math/big"

	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"
	"github.com/aead/skein"
	"golang.org/x/crypto/sha3"
)

// State is an intermediate data structure with pre-computed values to speed up mining.
type State struct {
	mat        matrix
	Timestamp  int64
	Nonce      uint64
	Target     big.Int
	prePowHash externalapi.DomainHash
	blockVersion uint16
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
	if header.Version() == 2 {
		return &State{
			Target:       *target,
			prePowHash:   *prePowHash,
			mat:          *generateKodaMatrix(prePowHash),
			Timestamp:    timestamp,
			Nonce:        nonce,
			blockVersion: header.Version(),
		}
	}
	return &State{
		Target:       *target,
		prePowHash:   *prePowHash,
		mat:          *generateMatrix(prePowHash),
		Timestamp:    timestamp,
		Nonce:        nonce,
		blockVersion: header.Version(),
	}
}

func (state *State) CalculateProofOfWorkValue() *big.Int {
	if state.blockVersion == 1 {
		return state.CalculateProofOfWorkValuePyrinhash()
	} else if state.blockVersion == 2 {
		return state.CalculateProofOfWorkValueKodahash()
	} else {
		return state.CalculateProofOfWorkValuePyrinhash() // default to the oldest version.
	}
}

// CalculateProofOfWorkValue hashes the internal header and returns its big.Int value
func (state *State) CalculateProofOfWorkValuePyrinhash() *big.Int {
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
	heavyHash := state.mat.HeavyHash(powHash)
	return toBig(heavyHash)
}

// CalculateProofOfWorkValueKodahash implements the hash chain (BLAKE2 -> Skein -> SHA3-256) for PoW
func (state *State) CalculateProofOfWorkValueKodahash() *big.Int {
	// PRE_POW_HASH || TIME || 32 zero byte padding || NONCE
	writer := hashes.HeavyHashWriter()
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

	// 1. BLAKE2 hashing
	blake2Hash := blake2b.Sum256(powHash.ByteSlice())

	// 2. Skein hashing
	skeinHasher := skein.New256(nil)
	skeinHasher.Write(blake2Hash[:])
	skeinHash := skeinHasher.Sum(nil)

	// 3. SHA3-256 hashing
	sha3Hasher := sha3.New256()
	sha3Hasher.Write(skeinHash)
	finalHash := sha3Hasher.Sum(nil)

	// Convert the final SHA3-256 hash to *externalapi.DomainHash
	var fixedHash [32]byte
	copy(fixedHash[:], finalHash[:32])
	cnHashDomain := externalapi.NewDomainHashFromByteArray(&fixedHash)

	// Pass the domain hash to HeavyKodaHash
	multiplied := state.mat.HeavyKodaHash(cnHashDomain)
	return toBig(multiplied)
}

// IncrementNonce increments the nonce in State by 1
func (state *State) IncrementNonce() {
	state.Nonce++
}

// CheckProofOfWork verifies if the block has a valid PoW according to the provided target
func (state *State) CheckProofOfWork() bool {
	// The block pow must be less than the claimed target
	powNum := state.CalculateProofOfWorkValue()

	// The block hash must be less or equal than the claimed target.
	return powNum.Cmp(&state.Target) <= 0
}

// CheckProofOfWorkByBits verifies if the block has a valid PoW according to its Bits field
func CheckProofOfWorkByBits(header externalapi.MutableBlockHeader) bool {
	return NewState(header).CheckProofOfWork()
}

// ToBig converts a externalapi.DomainHash into a big.Int treated as a little endian string
func toBig(hash *externalapi.DomainHash) *big.Int {
	// We treat the Hash as little-endian for PoW purposes, but the big package wants the bytes in big-endian, so reverse them.
	buf := hash.ByteSlice()
	blen := len(buf)
	for i := 0; i < blen/2; i++ {
		buf[i], buf[blen-1-i] = buf[blen-1-i], buf[i]
	}

	return new(big.Int).SetBytes(buf)
}

// BlockLevel returns the block level of the given header
func BlockLevel(header externalapi.BlockHeader, maxBlockLevel int) int {
	// Genesis is defined to be the root of all blocks at all levels, so we define it to be the maximal
	// block level.
	if len(header.DirectParents()) == 0 {
		return maxBlockLevel
	}

	proofOfWorkValue := NewState(header.ToMutable()).CalculateProofOfWorkValue()
	level := maxBlockLevel - proofOfWorkValue.BitLen()
	// If the block has a level lower than genesis, make it zero.
	if level < 0 {
		level = 0
	}
	return level
}
