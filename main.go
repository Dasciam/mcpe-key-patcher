package main

import (
	"bytes"
	"fmt"
	"os"
)

var (
	keyLength = 160
	oldPubKey = []byte("MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE8ELkixyLcwlZryUQcu1TvPOmI2B7vX83ndnWRUaXm74wFfa5f/lwQNTfrLVHa2PmenpGI6JhIMUJaWZrjmMj90NoKNFSNBuKdm8rYiXsfaz3K36x/1U26HpG0ZxK/V1V")
	newPubKey = []byte("MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAECRXueJeTDqNRRgJi/vlRufByu/2G0i2Ebt6YMar5QX/R0DIIyrJMcUpruK4QveTfJSTp3Shlq4Gk34cD/4GUWwkv0DVuzeuB+tXija7HBxii03NHDbPAD0AKnLr2wdAp")
)

// patch patches the file to the specified path
func patch(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	var buf bytes.Buffer

	lastIdx := 0
	for i := 0; i < len(data)-keyLength+1; i++ {
		if bytes.Equal(data[i:i+keyLength], oldPubKey) {
			buf.Write(data[lastIdx:i])
			buf.Write(newPubKey)
			lastIdx = i + keyLength
			fmt.Printf("Patched at %x\n", i)
		}
	}
	buf.Write(data[lastIdx:])

	err = os.WriteFile(path, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: ./patcher <file>")
		return
	}
	err := patch(args[0])
	if err != nil {
		fmt.Printf("Error patching: %v\n", err)
	}
}
