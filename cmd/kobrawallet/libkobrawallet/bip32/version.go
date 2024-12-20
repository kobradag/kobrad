package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// KobraMainnetPrivate is the version that is used for
// kobra mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var KobraMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// KobraMainnetPublic is the version that is used for
// kobra mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var KobraMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// KobraTestnetPrivate is the version that is used for
// kobra testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var KobraTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// KobraTestnetPublic is the version that is used for
// kobra testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var KobraTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// KobradevnetPrivate is the version that is used for
// kobra devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var KobradevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// KobradevnetPublic is the version that is used for
// kobra devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var KobradevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// KobraSimnetPrivate is the version that is used for
// kobra simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var KobraSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// KobraSimnetPublic is the version that is used for
// kobra simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var KobraSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case KobraMainnetPrivate:
		return KobraMainnetPublic, nil
	case KobraTestnetPrivate:
		return KobraTestnetPublic, nil
	case KobradevnetPrivate:
		return KobradevnetPublic, nil
	case KobraSimnetPrivate:
		return KobraSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case KobraMainnetPrivate:
		return true
	case KobraTestnetPrivate:
		return true
	case KobradevnetPrivate:
		return true
	case KobraSimnetPrivate:
		return true
	}

	return false
}
