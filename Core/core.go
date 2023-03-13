package core

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func Bruteforce(wordlist_path, type_hash, hash_target string) {
	var found bool

	// Abrir arquivo com a wordlist
	file, err := os.Open(wordlist_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Preparar scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	// Processar cada linha
	for scanner.Scan() {
		word := scanner.Text()

		switch type_hash {

		case "sha256":
			hash := sha256.Sum256([]byte(word))
			hash_string := hex.EncodeToString(hash[:])

			if hash_string != hash_target {
				fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
			} else {
				fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]\n\n", hash_string, word)
				found = true
			}

		case "sha512":
			hash := sha512.Sum512([]byte(word))
			hash_string := hex.EncodeToString(hash[:])

			if hash_string != hash_target {
				fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
			} else {
				fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]\n\n", hash_string, word)
				found = true
			}

		case "md5":
			hash := md5.Sum([]byte(word))
			hash_string := hex.EncodeToString(hash[:])

			if hash_string != hash_target {
				fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
			} else {
				fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]\n\n", hash_string, word)
				found = true
			}

		case "sha1":
			hash := sha1.Sum([]byte(word))
			hash_string := hex.EncodeToString(hash[:])

			if hash_string != hash_target {
				fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
			} else {
				fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]\n\n", hash_string, word)
				found = true
			}
		}

		if found {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

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
