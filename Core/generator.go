package core

import (
	"fmt"
	"os"
	"time"
)

func Run_generator(hash_target string, type_hash string, min int, max int) {

	fmt.Printf("\n\nExecutando a quebra do hash.\n")
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*_-" // Define novamente os caracteres que serão utilizados

	start := time.Now() // Salva o momento atual

	for length := min; length <= max; length++ { // Loop que percorre todos os comprimentos de senha desejados
		hour := time.Now().Format("15:04:05")

		fmt.Printf("[%v] [Testando senhas de %d caracteres]\n", hour, length)
		for i := 0; i < len(characters); i++ { // Loop que percorre todos os caracteres possíveis
			generatePassword(length, string(characters[i]), hash_target, type_hash, &start) // Chama a função generatePassword com um prefixo inicial e imprime a senha gerada imediatamente
		}
	}
}

func generatePassword(length int, prefix, hash_target, type_hash string, start *time.Time) bool {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*_-" // Define novamente os caracteres que serão utilizados

	if length == 1 { // Se o comprimento da senha for 1, a senha foi gerada
		stop := classification(prefix, type_hash, hash_target) // Verifica se a senha gerada é a desejada
		if stop {
			fmt.Printf(" Tempo total: [%v]\n\n", time.Since(*start)) // Imprime o tempo total desde o início da execução
			os.Exit(0)
		}
		return false // Retorna da função
	}

	for i := 0; i < len(characters); i++ { // Loop que percorre todos os caracteres possíveis
		generatePassword(length-1, prefix+string(characters[i]), hash_target, type_hash, start)    
    }

	return false // Se nenhum valor verdadeiro foi retornado até agora, retorna falso
}
