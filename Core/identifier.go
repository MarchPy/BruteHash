package core

import (
	"encoding/hex"
	"fmt"
)

func Indentifier(target_hash string) {
	hash_bytes, err := hex.DecodeString(target_hash)
	if err != nil {
		fmt.Println("Erro ao decodificar hash:", err)
		return
	}

	switch len(hash_bytes) {
	case 16:
		fmt.Printf("\nO hash é MD5\n\n")
	case 20:
		fmt.Printf("\nO hash é SHA-1\n\n")
	case 32:
		fmt.Printf("\nO hash é SHA-256\n\n")
	case 64:
		fmt.Printf("\nO hash é SHA-512\n\n")
	default:
		fmt.Println("O tamanho do hash não é compatível com nenhum hash conhecido.")
	}
}
