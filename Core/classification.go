package core

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func classification(word, type_hash, hash_target string) bool {
	var found bool

	switch type_hash {

	case "sha256":
		hash := sha256.Sum256([]byte(word))
		hash_string := hex.EncodeToString(hash[:])

		if hash_string != hash_target {
			// fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
		} else {
			fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]", hash_string, word)
			found = true
		}

	case "sha512":
		hash := sha512.Sum512([]byte(word))
		hash_string := hex.EncodeToString(hash[:])

		if hash_string != hash_target {
			// fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
		} else {
			fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]", hash_string, word)
			found = true
		}

	case "md5":
		hash := md5.Sum([]byte(word))
		hash_string := hex.EncodeToString(hash[:])

		if hash_string != hash_target {
			// fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
		} else {
			fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]", hash_string, word)
			found = true
		}

	case "sha1":
		hash := sha1.Sum([]byte(word))
		hash_string := hex.EncodeToString(hash[:])

		if hash_string != hash_target {
			// fmt.Printf("[Hash] -\033[92m%s\033[0m- [\033[91mSem correspondência\033[0m] [Palavra] \033[1m%s\033[0m\n", hash_string, word)
		} else {
			fmt.Printf("\n[Hash] -\033[92m%s\033[0m- [Senha: \033[92m%s\033[0m] [\033[92mSucesso\033[0m]", hash_string, word)
			found = true
		}
	}
	return found
}
