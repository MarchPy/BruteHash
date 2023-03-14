package core

import "os"

// Importa a biblioteca fmt para imprimir as senhas geradas

func Run_generator(hash_target string, type_hash string, min int, max int) {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Define os caracteres que serão utilizados na geração de senhas

	for length := min; length <= max; length++ { // Loop que percorre todos os comprimentos de senha desejados
		for i := 0; i < len(characters); i++ { // Loop que percorre todos os caracteres possíveis
			generatePassword(length, string(characters[i]), hash_target, type_hash) // Chama a função generatePassword com um prefixo inicial e imprime a senha gerada imediatamente
		}
	}
}

func generatePassword(length int, prefix, hash_target, type_hash string) {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Define novamente os caracteres que serão utilizados

	if length == 1 { // Se o comprimento da senha for 1, a senha foi gerada
		stop := classification(prefix, type_hash, hash_target) // Imprime a senha gerada
		if stop {
            os.Exit(0)
        }
        return                                         // Retorna da função
	}

	for i := 0; i < len(characters); i++ { // Loop que percorre todos os caracteres possíveis
		generatePassword(length-1, prefix+string(characters[i]), hash_target, type_hash) // Chama a função generatePassword recursivamente com um caractere a mais no prefixo e imprime a senha gerada imediatamente
	}
}
