package hashRel

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/spaolacci/murmur3"
	"math/big"
	"strings"
	"testing"
)

func Test01(t *testing.T) {
	s := "xxx123"
	seed := uint32(127)
	hash := murmur3.New64WithSeed(seed)
	hash.Write([]byte(s))
	//v := hash.Sum64()
	//t.Logf("[Hash] key: '%s', seed: '%d', v: %d", s, seed, v)
	t.Log(hash.Sum64())

	hash.Reset()
	hash.Write([]byte(s))
	t.Log(hash.Sum64())

	fmt.Printf(strings.Repeat("-", 72))
	fmt.Println()

	hash.Write([]byte(s))
	t.Log(hash.Sum64())

	hash.Write([]byte(s))
	t.Log(hash.Sum64())

}

func Test02(t *testing.T) {
	s := "xxx123"
	seed := uint32(127)
	hash := md5.New()
	hash.Write([]byte(s))
	bi := big.NewInt(0)
	bi.SetString(hex.EncodeToString(hash.Sum(nil)), 16)
	v := bi
	t.Logf("[Hash] key: '%s', seed: '%d', v: %d", s, seed, v)
}

func Test03(t *testing.T) {
	s := "xxx123"
	//seed := uint32(127)
	t.Log(murmur32(s))
	t.Log(murmur64(s))
	t.Log(murmur32(s))
	t.Log(murmur64(s))
	t.Log(murmur64WithSeed(s, 123))
	t.Log(murmur64WithSeed(s, 127))
	t.Log(murmur32WithSeed(s, 97))
	t.Log(MurmurHash64A(s))
}

func Test031(t *testing.T) {
	lst:=[] string {"python","golang","scala","java","c++","c","123"}
	for _,s :=range lst{
		fmt.Printf("%v: %v\n", s, MurmurHash64A(s))
	}
	fmt.Println("--------------------------------------")
	for _,s :=range lst{
		fmt.Printf("%v: %v\n", s, NumHash(s))
	}

	//s := "xxx123"
	//t.Log(MurmurHash64A(s))
}

func Test04(t *testing.T) {
	var bucketSize = 10
	var N = 10000000
	var n = float64(N)
	var bucketMap = map[uint64]int{}
	var bucketMapFloat = map[uint64]float64{}
	for i := 15000000000; i < 15000000000+N; i++ {
		hashInt := murmur64WithSeed(fmt.Sprint(i), 127) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
	for k, v := range bucketMap {
		bucketMapFloat[k] = float64(v) / n
		//fmt.Printf("k=%d, v=%f\n", k, float64(v)/n)
	}
	fmt.Println(bucketMapFloat)
}

func Test05(t *testing.T) {
	uniformCheck2(129)
	fmt.Println(strings.Repeat("-", 32))
	uniformCheck2(131)
}

func uniformCheck(seed uint32) {
	var bucketSize = 10
	var N = 10000000
	var n = float64(N)
	var bucketMap = map[uint64]int{}
	var bucketMapFloat = map[uint64]float64{}
	for i := 15000000000; i < 15000000000+N; i++ {
		hashInt := murmur64WithSeed(fmt.Sprint(i), seed) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
	for k, v := range bucketMap {
		bucketMapFloat[k] = float64(v) / n
		//fmt.Printf("k=%d, v=%f\n", k, float64(v)/n)
	}
	fmt.Println(bucketMapFloat)
}

func uniformCheck2(seed uint32) {
	var bucketSize = 10
	var N = 10000000
	var n = float64(N)
	var bucketMap = map[uint64]int{}
	var bucketMapFloat = map[uint64]float64{}
	for i := 15000000000; i < 15000000000+N; i++ {
		hashInt := murmur64WithSeed2(fmt.Sprint(i), seed) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
	for k, v := range bucketMap {
		bucketMapFloat[k] = float64(v) / n
		//fmt.Printf("k=%d, v=%f\n", k, float64(v)/n)
	}
	fmt.Println(bucketMapFloat)
}

func uniformCheck3(seed uint32) {
	var bucketSize = 10
	var N = 10000000
	var n = float64(N)
	var bucketMap = map[uint64]int{}
	var bucketMapFloat = map[uint64]float64{}
	for i := 15000000000; i < 15000000000+N; i++ {
		hashInt := murmur64(fmt.Sprint(i)) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
	for k, v := range bucketMap {
		bucketMapFloat[k] = float64(v) / n
		//fmt.Printf("k=%d, v=%f\n", k, float64(v)/n)
	}
	fmt.Println(bucketMapFloat)
}

func md5Hash(str string) [16]byte {
	return md5.Sum([]byte(str))
}

func sha1Hash(str string) [20]byte {
	return sha1.Sum([]byte(str))
}

func murmur32(str string) uint32 {
	return murmur3.Sum32([]byte(str))
}

func murmur32WithSeed(str string, seed uint32) uint32 {
	return murmur3.Sum32WithSeed([]byte(str), seed)
}

func murmur64(str string) uint64 {
	return murmur3.Sum64([]byte(str))
}

func murmur64WithSeed2(str string, seed uint32) uint64 {
	return murmur3.Sum64WithSeed([]byte(str), seed)
}

func murmur64WithSeed(str string, seed uint32) uint64 {
	hash := murmur3.New64WithSeed(seed)
	hash.Write([]byte(str))

	return hash.Sum64()
}
