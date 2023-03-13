# BruteHash
O pacote core contém duas funções para processamento de hashes:
- A função Bruteforce realiza a tentativa de quebra de um hash usando uma wordlist de palavras, procurando por correspondências com o hash fornecido. Atualmente, a função é capaz de quebrar hashes MD5, SHA-1, SHA-256 e SHA-512.

- A função Identifier identifica qual é o tipo de hash de acordo com o comprimento da string fornecida como entrada.

Essas funções podem ser utilizadas em ferramentas para testes de penetração e auditoria de segurança para encontrar senhas frágeis ou vulnerabilidades em sistemas que utilizam hashes para armazenar informações sensíveis.

# Utilização

## A função Bruteforce requer três parâmetros:

- **-w** (Wordlist path) é o caminho para o arquivo contendo a lista de palavras a serem testadas.
- **--type** (Type Hash) é o tipo de hash que será quebrado. Atualmente, os valores aceitos são "md5", "sha1", "sha256" e "sha512".
- **-t** (hash_target) é o hash que se deseja quebrar.

Exemplo:
####
    ./BruteHash bruteforce -w /usr/share/wordlists/rockyou.txt -t 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918 --type sha256
    
## A função Identifier requer um parâmetros:

- **-t** (hash_target) é o hash que se deseja quebrar.

Exemplo:
####
    ./BruteHash identifier -t 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918
    

# Screenshot
![image](https://user-images.githubusercontent.com/62616207/224595117-02036fc0-db47-41c3-a81e-3ca5d218b17a.png)


# Contribuição
Contribuições são bem-vindas. Se você encontrar um bug ou tiver alguma ideia de melhoria, por favor, abra uma issue ou submeta uma pull request.
