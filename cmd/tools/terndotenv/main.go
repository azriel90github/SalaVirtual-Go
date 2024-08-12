package main //mecanismo para organizar o código em módulos reutilizáveis

import (//tem a função fundamental de incorporar funcionalidades de outros módulos ou bibliotecas em seu código
	"os/exec" // Importa o pacote para executar comandos externos
	
	"github.com/joho/godotenv" // Importa o pacote para carregar variáveis de ambiente
)

func main() {//funcao para ponto de entrada do programa
		if err := godotenv.Load(); err != nil { // Tenta carregar as variáveis de ambiente do arquivo .env
			panic(err) // Se houver algum erro, o programa para com uma mensagem de erro
		}

		cmd := exec.Command( // Cria um novo comando para ser executado
			"tern", // Nome do comando a ser executado (provavelmente uma ferramenta de migração)
			"migrate", // Subcomando para executar a migração
			"--migrations", // Opção para especificar o diretório das migrações
			"./internal/store/pgstore/migrations", // Caminho para o diretório das migrações
			"--config", // Opção para especificar o arquivo de configuração
			"./internal/store/pgstore/migrations/tern.conf", // Caminho para o arquivo de configuração
		)
		if err := cmd.Run(); err != nil { // Executa o comando e verifica se houve algum erro
			panic(err) // Se houver algum erro, o programa pára com uma mensagem de erro
		}
}
