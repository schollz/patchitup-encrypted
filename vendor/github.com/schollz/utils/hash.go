package utils

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// Filemd5Sum returns the md5 sum of a file and produces the same
// hash for both Windows and Unix systems.
func Filemd5Sum(pathToFile string) (result string, err error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hash.Write(scanner.Bytes())
	}
	result = hex.EncodeToString(hash.Sum(nil))
	return
}

func Md5Sum(s string) (result string) {
	hash := md5.New()
	hash.Write([]byte(s))
	result = hex.EncodeToString(hash.Sum(nil))
	return
}

func Hash(data string) string {
	return HashBytes([]byte(data))
}

func HashBytes(data []byte) string {
	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum)
}
