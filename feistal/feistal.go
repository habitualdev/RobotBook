package feistal

import (
	"crypto/hmac"
	"crypto/sha256"
)

const BLOCK_SIZE = 64

type FeistelCipher struct {
	keys [][]byte
}

// New tries to read the keys from a local file (.feistelkeys). If
// if fails, it generates new keys and creates that file
func New(keys [][]byte) *FeistelCipher {

	f := FeistelCipher{keys: keys}
	return &f
}

func encryptBlock(block []byte, keys [][]byte) []byte {
	rounds := len(keys)
	l := block[:len(block)/2]
	r := block[(len(block) / 2):]

	for i := 1; i < rounds; i++ {
		l, r = r, roundFunc(l, r, keys[i])
	}

	return append(l, r...)
}

func decryptBlock(block []byte, keys [][]byte) []byte {
	rounds := len(keys)
	l := block[:len(block)/2]
	r := block[(len(block) / 2):]

	for i := rounds - 1; i > 0; i-- {
		r, l = l, roundFunc(r, l, keys[i])
	}
	return append(l, r...)
}

func roundFunc(l []byte, r []byte, key []byte) []byte {
	hmacObj := hmac.New(sha256.New, key)
	hmacObj.Write(r)
	hmacSum := hmacObj.Sum(nil)
	data := bytesXor(l, hmacSum)

	return data
}

func pad(msg []byte) []byte {
	msgLength := len(msg)
	max := BLOCK_SIZE
	for msgLength%BLOCK_SIZE != 0 {
		msg = append(msg, byte(max))
		msgLength = len(msg)
		max--
	}
	return msg
}

func stripPadding(msg []byte) []byte {
	paddingLength := BLOCK_SIZE + 1 - msg[len(msg)-1]
	messageLength := len(msg) - int(paddingLength)
	return msg[:messageLength]
}

func (f *FeistelCipher) Encrypt(msg []byte) []byte {
	paddedMsg := pad(msg)
	out := make([]byte, len(paddedMsg))
	for i := 0; i < len(paddedMsg); i += BLOCK_SIZE {
		msgBlock := paddedMsg[i : i+BLOCK_SIZE]
		encryptedBlock := encryptBlock(msgBlock, f.keys)
		for j := 0; j < BLOCK_SIZE; j++ {
			out[i+j] = encryptedBlock[j]
		}
	}
	return out
}

func (f *FeistelCipher) Decrypt(msg []byte) []byte {
	out := make([]byte, len(msg))
	for i := 0; i < len(msg); i += BLOCK_SIZE {
		msgBlock := msg[i : i+BLOCK_SIZE]
		decryptedBlock := decryptBlock(msgBlock, f.keys)
		for j := 0; j < BLOCK_SIZE; j++ {
			out[i+j] = decryptedBlock[j]
		}
	}
	return stripPadding(out)
}
