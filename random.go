package tool

import (
	"math/rand"
	"time"
)

/*RandomKind RandomKind */
type RandomKind int

/*random kinds */
const (
	RandomNum      RandomKind = iota // 纯数字
	RandomLower                      // 小写字母
	RandomUpper                      // 大写字母
	RandomLowerNum                   // 数字、小写字母
	RandomUpperNum                   // 数字、大写字母
	RandomAll                        // 数字、大小写字母
)

/*RandomString defines */
var (
	RandomString = []string{
		RandomNum:      "0123456789",
		RandomLower:    "abcdefghijklmnopqrstuvwxyz",
		RandomUpper:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomLowerNum: "0123456789abcdefghijklmnopqrstuvwxyz",
		RandomUpperNum: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomAll:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
)

//GenerateRandomString 随机字符串
func GenerateRandomString(size int, kind ...RandomKind) string {
	s := RandomString[RandomAll]
	if len(kind) > 0 {
		s = RandomString[kind[0]]
	}

	return string(randomByte(s, size))
}

func randomByte(src string, size int) (r []byte) {
	rand.Seed(time.Now().UnixNano())
	ls := len(src)
	for ; size > 0; size-- {
		r = append(r, src[rand.Intn(ls)])
	}
	return
}

//GenerateRandom random characters with []byte
func GenerateRandom(size int, kind ...RandomKind) []byte {
	s := RandomString[RandomAll]
	if len(kind) > 0 {
		s = RandomString[kind[0]]
	}
	return randomByte(s, size)
}
