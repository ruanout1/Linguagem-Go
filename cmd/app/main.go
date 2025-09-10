// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/hello"
	"github.com/seu-usuario/meu-projeto-go/internal/hello/fibonacci"
)

// Função principal do programa
func main() {

	hello.SayHello()
	posicao := 10
	resultado := fibonacci.Fibonacci(posicao)

	fmt.Printf("O %dº número de Fibonacci é: %d\n", posicao, resultado)
}
