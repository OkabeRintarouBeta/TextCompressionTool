package Util

import (
	"bufio"
	"compressText/HuffmanTree"
	"io"
	"log"
	"os"
	"strings"
)

const (
	KeyValSeparator     = "-<>~"
	KeyValPairSeparator = "%->"
	HeaderBodySeparator = "\n~<=>-\n"
)

func Compress(fileName, compressedFileName string) {
	// Open the original file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Calculate the frequency of each character in the original file
	charFrequencyMap := make(map[string]int)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading rune:", err)
		}
		charFrequencyMap[string(r)]++
	}

	// Build Huffman encoding map
	charEncodingMap := HuffmanTree.BuildHuffmanEncoding(charFrequencyMap)

	// Build header
	var headerBuilder strings.Builder
	for k, v := range charEncodingMap {
		headerBuilder.WriteString(k)
		headerBuilder.WriteString(KeyValSeparator)
		headerBuilder.WriteString(v)
		headerBuilder.WriteString(KeyValPairSeparator)
	}
	headerStr := strings.TrimRight(headerBuilder.String(), KeyValPairSeparator)

	// Second pass: encode and count valid bits
	file.Seek(0, 0)

	reader.Reset(file)
	compressedData := encodeToPackedBytes(reader, charEncodingMap)

	// Final content
	finalContent := []byte(headerStr + HeaderBodySeparator)
	finalContent = append(finalContent, compressedData...)

	// Write to output file
	outputFile, err := os.OpenFile(compressedFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Failed to open output file:", err)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(finalContent)
	if err != nil {
		log.Fatal("Failed to write output:", err)
	}
}

func encodeToPackedBytes(reader *bufio.Reader, encodingMap map[string]string) []byte {
	var output []byte
	var currentByte byte
	bitCount := 0

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading byte:", err)
		}

		code := encodingMap[string(r)]
		for _, bit := range code {
			currentByte <<= 1
			if bit == '1' {
				currentByte |= 1
			}
			bitCount++
			if bitCount == 8 {
				output = append(output, currentByte)
				currentByte = 0
				bitCount = 0
			}
		}
	}
	if bitCount > 0 {
		currentByte <<= 8 - bitCount
		output = append(output, currentByte)
	}

	return output
}
