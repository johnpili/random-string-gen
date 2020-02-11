package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	keyset := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890~!@#$%^&*()_+|}{[];':,./<>?fghijklmnopqrstuvwxyz1234567890~!@#$%^&*()_+|}{[];':,./<>?")
	var size int
	flag.IntVar(&size, "s", 32, "Set the byte size of the generated string")
	flag.Parse()
	gen := make([]byte, size)
	for i := 0; i <= size; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)                
		x := keyset[r1.Intn(len(keyset))]                
		gen = append(gen, x)
	}
	fmt.Println(string(gen))
}
