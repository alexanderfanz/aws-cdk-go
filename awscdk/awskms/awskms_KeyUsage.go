package awskms


// The key usage, represents the cryptographic operations of keys.
//
// Example:
//   key := kms.NewKey(this, jsii.String("MyKey"), &keyProps{
//   	keySpec: kms.keySpec_ECC_SECG_P256K1,
//   	 // Default to SYMMETRIC_DEFAULT
//   	keyUsage: kms.keyUsage_SIGN_VERIFY,
//   })
//
// Experimental.
type KeyUsage string

const (
	// Encryption and decryption.
	// Experimental.
	KeyUsage_ENCRYPT_DECRYPT KeyUsage = "ENCRYPT_DECRYPT"
	// Signing and verification.
	// Experimental.
	KeyUsage_SIGN_VERIFY KeyUsage = "SIGN_VERIFY"
)

