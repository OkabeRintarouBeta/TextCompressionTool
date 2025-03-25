package Util

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Decompress(compressedFileName, outputFileName string) {
	file, err := os.Open(compressedFileName)
	if err != nil {
		log.Fatal("Failed to open compressed file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error while reading compressed file:", err)
	}
	parts := strings.SplitN(string(data), HeaderBodySeparator, 2)
	if len(parts) < 2 {
		log.Fatal("Invalid compressed file format: missing header-body separator")
	}
	header, body := parts[0], parts[1]

	codeToChar := make(map[string]string)
	for _, pair := range strings.Split(header, KeyValPairSeparator) {
		kvPair := strings.Split(pair, KeyValSeparator)
		if len(kvPair) != 2 {
			log.Fatal("Invalid compressed file header format: missing key-value separator")
		}
		codeToChar[kvPair[1]] = kvPair[0]
	}
	// Convert compressed body to bitstream
	compressedData := []byte(body)
	bitStream := bytesToBitStream(compressedData)

	var decoded strings.Builder
	var currCode strings.Builder

	for _, bit := range bitStream {
		currCode.WriteByte(byte(bit))
		if char, exists := codeToChar[currCode.String()]; exists {
			decoded.WriteString(char)
			currCode.Reset()
		}
	}
	err = os.WriteFile(outputFileName, []byte(decoded.String()), 0644)
	if err != nil {
		log.Fatal("Failed to write output:", err)
	}
}

func bytesToBitStream(data []byte) string {
	var result strings.Builder
	for _, b := range data {
		result.WriteString(fmt.Sprintf("%08b", b))
	}
	return result.String()
}
