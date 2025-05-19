package hash

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/corona10/goimagehash"
)

func Generate(filepath string) (uint64, error) {
	fmt.Println("Opening file:", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}
	defer file.Close()

	fmt.Println("Decoding image...")
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return 0, err
	}

	fmt.Println("Generating hash...")
	hash, err := goimagehash.AverageHash(img)
	if err != nil {
		fmt.Println("Error generating hash:", err)
		return 0, err
	}

	fmt.Println("Hash generated successfully:", hash.GetHash())
	return hash.GetHash(), nil
}

func CompareHashes(hash1, hash2 uint64) (int, error) {
	h1 := goimagehash.NewImageHash(hash1, goimagehash.AHash)
	h2 := goimagehash.NewImageHash(hash2, goimagehash.AHash)
	return h1.Distance(h2)
}
