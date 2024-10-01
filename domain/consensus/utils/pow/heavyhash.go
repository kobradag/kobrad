package pow

import (
	"math"

	"github.com/kobradag/kobrad/domain/consensus/model/externalapi"
	"github.com/kobradag/kobrad/domain/consensus/utils/hashes"
)

const eps float64 = 1e-9

type matrix [64][64]uint16

// Define the DAA Score threshold for switching to the new method
// The value of 10,860,000 marks the transition point to the new HeavyHash method.
const transitionDAAScore uint64 = 100

// Function that chooses the appropriate hashing algorithm based on DAA Score
// If the DAA Score is above the transition threshold, it uses the new HeavyHash method.
// Otherwise, it defaults to the old HeavyHash method.
func HeavyHashWithTransition(mat *matrix, hash *externalapi.DomainHash, blockDAAScore uint64) *externalapi.DomainHash {
	// If the DAA Score is above or equal to the threshold, use the new method
	if blockDAAScore >= transitionDAAScore {
		// Use the new hashing method
		return mat.newHeavyHash(hash)
	}
	// Use the old hashing method if the DAA Score is below the threshold
	return mat.oldHeavyHash(hash)
}

// Old method (the one you provided)
// This function implements the original version of HeavyHash, used for blocks below the DAA threshold.
func (mat *matrix) oldHeavyHash(hash *externalapi.DomainHash) *externalapi.DomainHash {
	hashBytes := hash.ByteArray()
	var vector [64]uint16
	var product [64]uint16
	for i := 0; i < 32; i++ {
		vector[2*i] = uint16(hashBytes[i] >> 4)
		vector[2*i+1] = uint16(hashBytes[i] & 0x0F)
	}
	// Perform matrix-vector multiplication and convert to 4 bits
	for i := 0; i < 64; i++ {
		var sum uint16
		for j := 0; j < 64; j++ {
			sum += mat[i][j] * vector[j]
		}
		product[i] = sum >> 10
	}

	// Concatenate the 4 least significant bits back to 8-bit XOR with sum1
	var res [32]byte
	for i := range res {
		res[i] = hashBytes[i] ^ (byte(product[2*i]<<4) | byte(product[2*i+1]))
	}
	// Hash again
	writer := hashes.HeavyHashWriter()
	writer.InfallibleWrite(res[:])
	return writer.Finalize()
}

// New method (your updated logic)
// This function implements the new version of HeavyHash, which is used after the DAA Score threshold is reached.
func (mat *matrix) newHeavyHash(hash *externalapi.DomainHash) *externalapi.DomainHash {
    hashBytes := hash.ByteArray()

    var vector [64]uint16
    var product [64]uint16

    for i := 0; i < 32; i++ {
        vector[2*i] = uint16(hashBytes[i] >> 4)
        vector[2*i+1] = uint16(hashBytes[i] & 0x0F)
    }

    for i := 0; i < 64; i++ {
        var sum uint16
        for j := 0; j < 64; j++ {
            dynamicShift := (vector[j] % 7) + 1  // Ð¡Ð´Ð²Ð¸Ð³ Ð¾Ñ‚ 1 Ð´Ð¾ 7 Ð±Ð¸Ñ‚
            condition := (vector[j] + mat[i][j]) % 2 == 0

            if condition {
                sum += (mat[i][j] * vector[j]) << dynamicShift
            } else {
                sum += (mat[i][j] * vector[j]) >> dynamicShift
            }
        }
        product[i] = uint16(math.Sin(float64(sum)) * 0xFFFF) ^
                     uint16(math.Tan(float64(sum)) * 0xFFFF) ^
                     ((sum >> 4) & 0xF) ^
                     ((sum >> 8) & 0xF)
    }


    var res [32]byte
    for i := range res {
        dynamicShift := (product[2*i] % 5) + 3  
        res[i] = hashBytes[i] ^ (byte(product[2*i]<<dynamicShift) | byte(product[2*i+1]>>dynamicShift))
    }

    rounds := (uint32(hashBytes[0]) % 5) + 4  
    for r := 0; r < int(rounds); r++ {
        writer := hashes.HeavyHashWriter()
        writer.InfallibleWrite(res[:])
        finalizedHash := writer.Finalize()

        res = *finalizedHash.ByteArray()
    }

    writer := hashes.HeavyHashWriter()
    writer.InfallibleWrite(res[:])
    return writer.Finalize()
}




// Function to generate the matrix (no changes here)
// This function generates the matrix required for the HeavyHash computation.
func generateMatrix(hash *externalapi.DomainHash) *matrix {
	var mat matrix
	generator := newxoShiRo256PlusPlus(hash)
	for {
		for i := range mat {
			for j := 0; j < 64; j += 16 {
				val := generator.Uint64()
				for shift := 0; shift < 16; shift++ {
					mat[i][j+shift] = uint16(val >> (4 * shift) & 0x0F)
				}
			}
		}
		if mat.computeRank() == 64 {
			return &mat
		}
	}
}

// Compute matrix rank (no changes here)
// This function computes the rank of the matrix to ensure its suitability for use in HeavyHash.
func (mat *matrix) computeRank() int {
	var B [64][64]float64
	for i := range B {
		for j := range B[0] {
			B[i][j] = float64(mat[i][j])
		}
	}
	var rank int
	var rowSelected [64]bool
	for i := 0; i < 64; i++ {
		var j int
		for j = 0; j < 64; j++ {
			if !rowSelected[j] && math.Abs(B[j][i]) > eps {
				break
			}
		}
		if j != 64 {
			rank++
			rowSelected[j] = true
			for p := i + 1; p < 64; p++ {
				B[j][p] /= B[j][i]
			}
			for k := 0; k < 64; k++ {
				if k != j && math.Abs(B[k][i]) > eps {
					for p := i + 1; p < 64; p++ {
						B[k][p] -= B[j][p] * B[k][i]
					}
				}
			}
		}
	}
	return rank
}

func (mat *matrix) HeavyHash(hash *externalapi.DomainHash) *externalapi.DomainHash {

	hashBytes := hash.ByteArray()

	var vector [64]uint16

	var product [64]uint16

	for i := 0; i < 32; i++ {

		vector[2*i] = uint16(hashBytes[i] >> 4)

		vector[2*i+1] = uint16(hashBytes[i] & 0x0F)

	}

	// Matrix-vector multiplication, and convert to 4 bits.

	for i := 0; i < 64; i++ {

		var sum uint16

		for j := 0; j < 64; j++ {

			sum += mat[i][j] * vector[j]

		}

		product[i] = (sum & 0xF) ^ ((sum >> 4) & 0xF) ^ ((sum >> 8) & 0xF)

	}



	// Concatenate 4 LSBs back to 8 bit xor with sum1

	var res [32]byte

	for i := range res {

		res[i] = hashBytes[i] ^ (byte(product[2*i]<<4) | byte(product[2*i+1]))

	}

	// Hash again

	writer := hashes.HeavyHashWriter()

	writer.InfallibleWrite(res[:])

	return writer.Finalize()

}
