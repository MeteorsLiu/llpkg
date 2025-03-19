package main

import (
	"fmt"

	zip "github.com/goplus/llpkg/bzip3"
)

//export PKG_CONFIG_PATH="/Users/heulucklu/code/gop/llpkg/bzip3:$PKG_CONFIG_PATH"
func main() {
	fmt.Println(*zip.Bz3Version())

	// 测试用例：压缩和解压缩
	input := []byte("Hello, bzip3 compression!")
	output := make([]byte, zip.Bz3Bound(uintptr(len(input))))
	outputSize := uintptr(len(output))

	// 压缩
	errCode := zip.Bz3Compress(1024*1024, &input[0], &output[0], uintptr(len(input)), &outputSize)
	if errCode != zip.BZ3_OK {
		fmt.Println("Compression failed with error code:", errCode)
		return
	}
	fmt.Println("Compression successful. Compressed size:", outputSize)

	// 解压缩
	decompressed := make([]byte, len(input))
	decompressedSize := uintptr(len(decompressed))
	errCode = zip.Bz3Decompress(&output[0], &decompressed[0], outputSize, &decompressedSize)
	if errCode != zip.BZ3_OK {
		fmt.Println("Decompression failed with error code:", errCode)
		return
	}
	fmt.Println("Decompression successful. Decompressed data:", string(decompressed))
}
