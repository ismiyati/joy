package subtlecrypto

import (
	"github.com/matthewmueller/golly/dom/cryptokey"
	"github.com/matthewmueller/golly/js"
)

// SubtleCrypto struct
// js:"SubtleCrypto,omit"
type SubtleCrypto struct {
}

// Decrypt fn
func (*SubtleCrypto) Decrypt(algorithm interface{}, key *cryptokey.CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.decrypt($1, $2, $3)", algorithm, key, data)
	return i
}

// DeriveBits fn
func (*SubtleCrypto) DeriveBits(algorithm interface{}, baseKey *cryptokey.CryptoKey, length uint) (i interface{}) {
	js.Rewrite("await $<.deriveBits($1, $2, $3)", algorithm, baseKey, length)
	return i
}

// DeriveKey fn
func (*SubtleCrypto) DeriveKey(algorithm interface{}, baseKey *cryptokey.CryptoKey, derivedKeyType interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.deriveKey($1, $2, $3, $4, $5)", algorithm, baseKey, derivedKeyType, extractable, keyUsages)
	return i
}

// Digest fn
func (*SubtleCrypto) Digest(algorithm interface{}, data []byte) (i interface{}) {
	js.Rewrite("await $<.digest($1, $2)", algorithm, data)
	return i
}

// Encrypt fn
func (*SubtleCrypto) Encrypt(algorithm interface{}, key *cryptokey.CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.encrypt($1, $2, $3)", algorithm, key, data)
	return i
}

// ExportKey fn
func (*SubtleCrypto) ExportKey(format string, key *cryptokey.CryptoKey) (i interface{}) {
	js.Rewrite("await $<.exportKey($1, $2)", format, key)
	return i
}

// GenerateKey fn
func (*SubtleCrypto) GenerateKey(algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.generateKey($1, $2, $3)", algorithm, extractable, keyUsages)
	return i
}

// ImportKey fn
func (*SubtleCrypto) ImportKey(format string, keyData []byte, algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.importKey($1, $2, $3, $4, $5)", format, keyData, algorithm, extractable, keyUsages)
	return i
}

// Sign fn
func (*SubtleCrypto) Sign(algorithm interface{}, key *cryptokey.CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.sign($1, $2, $3)", algorithm, key, data)
	return i
}

// UnwrapKey fn
func (*SubtleCrypto) UnwrapKey(format string, wrappedKey []byte, unwrappingKey *cryptokey.CryptoKey, unwrapAlgorithm interface{}, unwrappedKeyAlgorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.unwrapKey($1, $2, $3, $4, $5, $6, $7)", format, wrappedKey, unwrappingKey, unwrapAlgorithm, unwrappedKeyAlgorithm, extractable, keyUsages)
	return i
}

// Verify fn
func (*SubtleCrypto) Verify(algorithm interface{}, key *cryptokey.CryptoKey, signature []byte, data []byte) (i interface{}) {
	js.Rewrite("await $<.verify($1, $2, $3, $4)", algorithm, key, signature, data)
	return i
}

// WrapKey fn
func (*SubtleCrypto) WrapKey(format string, key *cryptokey.CryptoKey, wrappingKey *cryptokey.CryptoKey, wrapAlgorithm interface{}) (i interface{}) {
	js.Rewrite("await $<.wrapKey($1, $2, $3, $4)", format, key, wrappingKey, wrapAlgorithm)
	return i
}
