package core

import (
	"bufio"
	"log"
	"os"
)

func Bruteforce(wordlist_path, type_hash, hash_target string) {
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
		found := classification(word, type_hash, hash_target)
		if found {
			break
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)

		}
	}
}
