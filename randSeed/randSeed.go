package main

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"math/rand"
	"net"
	"time"
)

/* 加密 encryption */
// net
func GetMacAddrs() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		return macAddr, nil
	}
	return "", err
}

// hash output uint32
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// add mac and time.now() as seed
func Hashseed() int64 {
	mac_adr, _ := GetMacAddrs()
	t := time.Now().UnixNano() // int64
	//test := fmt.Sprintf("%d %s", t, mac_adr)
	fmt.Println(rand.Uint32())
	return 0
	return int64(Hash(fmt.Sprintf("%d %s", t, mac_adr)))
}

func FindSmallestPrimeLargerThan(n int) int {

Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				fmt.Print(1)
				break Outer
			case n%i == 0:
				fmt.Print(2)
				break Outer
			}
		}
	}
	return n

}


func main() {

	rand.Seed(Hashseed())
	rnd := make([]byte, 4)
	binary.LittleEndian.PutUint32(rnd, rand.Uint32())
	fmt.Println(rnd)

}
