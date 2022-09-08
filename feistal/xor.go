package feistal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

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
