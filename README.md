# Text Compression Command Line Tool
This is an implementation of a huffman encoder/decoder based text compression tool, implemented in Go.

This project is part of the [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-huffman).

## How to use
1. To build the project, use `go build .`
    A `compressText` binary will be generated.
2. To use the tool, use `./compressText` command.
    Flags to specify:

    a. `-d`: decompress operation

    b. `-in`: input file path(relative path)

    c. `-out`: output file path(relative path)
3. Example usage:

    a. To compress the `test.txt` in `test_files` directory,
    ```
   ./compressText -in=test_files/test.txt -out=test_files/test_compressed.txt
   ```
   You could see that the file size is compressed to 1.8MB.
   ![example](https://github.com/OkabeRintarouBeta/TextCompressionTool/blob/main/imgs/example-size.jpg?raw=true)    

    b. To decompress back to the original file, use
    ```
    ./compressText -d -in=test_files/test_compressed.txt -out=test_files/test_recovered.txt
   ```
   
    It could be seen that the decompressed file is the same as the original `test.txt`.
   ![example-2](https://github.com/OkabeRintarouBeta/TextCompressionTool/blob/main/imgs/example-size-2.jpg?raw=true)
## Next steps
Compress multiple files concurrently

## References
1. Coding Challenge: https://codingchallenges.fyi/challenges/challenge-huffman
2. Huffman Tree: https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Huffman.html
3. Priority queue in Go: https://pkg.go.dev/container/heap