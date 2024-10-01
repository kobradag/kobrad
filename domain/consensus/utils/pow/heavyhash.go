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
const transitionDAAScore uint64 = 10_860_000

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
	// Convert the hash bytes into a vector with 64 elements
	hashBytes := hash.ByteArray()
	var vector [64]uint16
	var product [64]uint16
	for i := 0; i < 32; i++ {
		// Each byte is split into two 4-bit values to create 64 elements
		vector[2*i] = uint16(hashBytes[i] >> 4)
		vector[2*i+1] = uint16(hashBytes[i] & 0x0F)
	}
	// Perform matrix-vector multiplication and convert to 4 bits
	for i := 0; i < 64; i++ {
		var sum uint16
		for j := 0; j < 64; j++ {
			// Multiply matrix element with the corresponding vector element and sum it up
			sum += mat[i][j] * vector[j]
		}
		// Keep only the top 4 bits of the sum for each product element
		product[i] = sum >> 10
	}

	// Concatenate the 4 least significant bits back to 8-bit XOR with sum1
	var res [32]byte
	for i := range res {
		// XOR the result with the original hash byte using the matrix product
		res[i] = hashBytes[i] ^ (byte(product[2*i]<<4) | byte(product[2*i+1]))
	}
	// Hash again using the HeavyHash writer
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

    // Convert the hash bytes into a vector of 64 elements
    for i := 0; i < 32; i++ {
        vector[2*i] = uint16(hashBytes[i] >> 4)
        vector[2*i+1] = uint16(hashBytes[i] & 0x0F)
    }

    for i := 0; i < 64; i++ {
        var sum uint16
        for j := 0; j < 64; j++ {
            // Dynamic bit shift based on the vector element
            dynamicShift := (vector[j] % 7) + 1  // Shift from 1 to 7 bits
            // Condition determines whether to shift left or right
            condition := (vector[j] + mat[i][j]) % 2 == 0

            if condition {
                sum += (mat[i][j] * vector[j]) << dynamicShift
            } else {
                sum += (mat[i][j] * vector[j]) >> dynamicShift
            }
        }
        // Use trigonometric functions and bit manipulation to further modify the sum
        product[i] = uint16(math.Sin(float64(sum)) * 0xFFFF) ^
                     uint16(math.Tan(float64(sum)) * 0xFFFF) ^
                     ((sum >> 4) & 0xF) ^
                     ((sum >> 8) & 0xF)
    }

    // Combine the product into the result array by shifting its bits
    var res [32]byte
    for i := range res {
        dynamicShift := (product[2*i] % 5) + 3  // Shift from 3 to 7 bits
        res[i] = hashBytes[i] ^ (byte(product[2*i]<<dynamicShift) | byte(product[2*i+1]>>dynamicShift))
    }

    // Repeat the hashing process for a dynamic number of rounds
    rounds := (uint32(hashBytes[0]) % 5) + 4  // 4 to 8 rounds
    for r := 0; r < int(rounds); r++ {
        writer := hashes.HeavyHashWriter()
        writer.InfallibleWrite(res[:])
        finalizedHash := writer.Finalize()

        // Use the result of each round as the input for the next
        res = *finalizedHash.ByteArray()
    }

    // Final hash round to get the final result
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
				// Generate random values for the matrix using the generator
				val := generator.Uint64()
				for shift := 0; shift < 16; shift++ {
					// Fill the matrix with 16 values from the generator
					mat[i][j+shift] = uint16(val >> (4 * shift) & 0x0F)
				}
			}
		}
		// Ensure the matrix is full rank before returning it
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
			// Convert the matrix to floating point for rank calculation
			B[i][j] = float64(mat[i][j])
		}
	}
	var rank int
	var rowSelected [64]bool
	for i := 0; i < 64; i++ {
		var j int
		// Look for the first unselected row with a non-zero value
		for j = 0; j < 64; j++ {
			if !rowSelected[j] && math.Abs(B[j][i]) > eps {
				break
			}
		}
		if j != 64 {
			rank++
			rowSelected[j] = true
			// Normalize the row by its pivot element
			for p := i + 1; p < 64; p++ {
				B[j][p] /= B[j][i]
			}
			// Subtract multiples of the pivot row from all other rows
			for k := 0; k < 64; k++ {
				if k != j && math.Abs(B[k][i]) > eps {
					for p := i + 1; p < 64; p++ {
						B[k][p] -= B[j][p] * B[k][i]
					}
				}
			}
		}
	}
	// Return the rank of the matrix
	return rank
}

// Original HeavyHash method (without transition logic)
// This method multiplies the matrix by a vector derived from the hash, 
// applies bitwise XOR, and hashes the result again.
func (mat *matrix) HeavyHash(hash *externalapi.DomainHash) *externalapi.DomainHash {
	hashBytes := hash.ByteArray()

	var vector [64]uint16
	var product [64]uint16

	// Convert hash bytes to a vector of 64 elements
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
		// Manipulate the sum using bitwise operations
		product[i] = (sum & 0xF) ^ ((sum >> 4) & 0xF) ^ ((sum >> 8) & 0xF)
	}

	// Concatenate 4 LSBs back to 8-bit xor with sum1
	var res [32]byte
	for i := range res {
		res[i] = hashBytes[i] ^ (byte(product[2*i]<<4) | byte(product[2*i+1]))
	}

	// Hash again
	writer := hashes.HeavyHashWriter()
	writer.InfallibleWrite(res[:])
	return writer.Finalize()
}
