package datagenerator

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func RandomAge() int {
	return seededRand.Intn(90) + 10
}

func RandomInt() int {
	return seededRand.Intn(100)
}

func RandomBool() bool {
	return seededRand.Int()%2 == 0
}

func RandomTime() time.Time {
	start := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now().AddDate(1, 5, 15)

	duration := end.Sub(start)

	randomDuration := time.Duration(rand.Int63n(int64(duration)))

	return start.Add(randomDuration)
}
