package main

import (
	"compressText/Util"
	"flag"
)

func main() {
	toDecompress := flag.Bool("d", false, "Enable  to decompress the target file")
	filePath := flag.String("in", "", "Input file path of the file to compress/decompress")
	outputFilePath := flag.String("out", "output.txt", "Output file path of the compressed/decompressed file")
	flag.Parse()

	if *toDecompress {
		Util.Decompress(*filePath, *outputFilePath)
	} else {
		Util.Compress(*filePath, *outputFilePath)
	}
}
