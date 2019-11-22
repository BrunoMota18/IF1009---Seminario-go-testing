# IF1009-Seminario-go-testing
Seminário com demo para disciplina de Teste de Software do CIn-UFPE.  
Este repositório tem como objetivo demonstrar funcionalidades da biblioteca 
de testes **go-testing**. 

---

## [contaPalavras](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/contaPalavras/contaPalavras.go)
A função **contaPalavras** recebe uma string e retorna um dicionário com as frequências de cada palavra em uma frase passada como parâmetro. Caso for passado um string vazio deve retornar um erro.

## [contaPalavras_test](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/contaPalavras/contaPalavras_test.go)
O teste checa o dicionário retornado e se existe algum erro assim como também se ele retorna um erro quando deve.

## [dummy](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/dummy/dummy.go)
Uma função simples que retorna `true` se o parâmetro for maior que 5 e `false` caso contrário.

## [dummy_test](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/dummy/dummy_test.go)
Uma função de teste criada para a função mencionada acima. Só cobre um único caso. Nós veremos mais disto abaixo na seção de Test Coverage.

# Testes com Go!
Ao fim desse tutorial você estará por dentro dos conceitos básicos de testes em Go.

[Passo 1: Instalação](#passo-1:-instalação)  
* [Passo 1.1: Testando o funcionamento no seu computador](#passo-1.1:-testando-o-funcionamento-no-seu-computador)

[Passo 2: Como fazer testes com Golang](#passo-2:-como-fazer-testes-com-golang)  
* [Passo 2.1: Exemplo de teste utilizando GO](#passo-2.1:-exemplo-de-teste-utilizando-go)
* [Passo 2.2: Rodando testes em Go](#passo-2.2:-rodando-testes-em-go)
    * [go test](#go-test)
    * [go test -v](#go-test--v)
    * [go test -run](#go-test--run)
    * [gotest](#gotest)

[Passo 3: Test Coverage](#passo-3:-test-coverage)

[Passo 4: Benchmarks](#passo-4:-benchmarks)
* [Passo 4.1: Exemplo de benchmark utilizando GO](#passo-4.1-exemplo-de-benchmark-utilizando-go)

# Passo 1: Instalação

Inicialmente será necessário a instalação no seu computador.
 A documentação original pode ser encontrada [aqui](http://www.golangbr.org/doc/instalacao), assim como o instalador.


# Passo 1.1: Testando o funcionamento no seu computador

Para testar o funcionamento da instalação, você pode criar a seguinte função em **Go**:

    package main
	import "fmt"
	
	func main() {
	    fmt.Printf("olá, mundo\n")
    }

Após ter escrito, deve chamá-la conforme descrito abaixo:

    $ go run ola.go
    olá, mundo

Caso consiga rodar sem erros, está tudo certo!

# Passo 2: Como fazer testes com Golang
Escrever testes em Go é como escrever qualquer outra função, porém teremos algumas regras para o funcionamento correto.
- Como em Go já existe um recurso de testes nativo da linguagem, você só deve importar **testing**.

        import "testing"

- O nome do arquivo deve sempre ser o nome do seu arquivo seguido de **_test.go**. Na primeira linha criamos via terminal para sistemas Unix e abaixo uma alternativa usando Powershell do Windows:

        touch contaPalavras_test.go #para sistemas Unix
        echo > contaPalavras_test.go #para Unix ou Windows Powershell

- A função de teste deve começar com **Test** seguido do seu nome. 

        TestContaPalavras(...)

- O único argumento da função de teste recebe é **t  *testing.T**. 

        TestContaPalavras(t *testing.T)

## Passo 2.1: Exemplo de teste utilizando GO
Aqui iremos demonstrar um exemplo da função **`contaPalavras()`** que tem com objetivo retornar um dicionário com as frequências de palavras de uma frase. Comece criando um arquivo e adicionando o código abaixo:
### contaPalavras.go
    package main

    import (
        "errors"
        "fmt"
        "log"
        "strings"
    )

    func contaPalavras(frase string) (map[string]int, error) {
        dicionario := make(map[string]int)
        for _, palavra := range strings.Fields(frase) {
            dicionario[palavra]++
        }
        //fmt.Println(len(dicionario))
        if len(dicionario) == 0 {
            return nil, errors.New("o dicionario está vazio; certifique se digitou alguma frase")
        }
        return dicionario, nil
    }

    func main() {
        fmt.Println(contaPalavras("Hasta la vista baby"))
        fmt.Println(contaPalavras("Hasta la vista baby vista la baby hasta la"))
        if _, err := contaPalavras("   "); err != nil {
            log.Fatal("Erro!: ", err)
        }
    }

### contaPalavras_test.go

Em Go cada caso de teste da função a ser testada é especificado pelo comando `t.Run(...)`, logo você especificar diferentes caminhos a serem testados em uma única função.

Não existem asserts nativos mas podem ser criados como uma função Helper por `t.Helper()` abstraindo o código e deixando os testes propriamenente ditos mas objetivos.

Ainda podemos utilizar bibliotecas implementadas por diferentes usuários da comunidade para auxiliar na criação de testes como é o caso de **assert** que pode ser adquirida pelo comando no terminal:

    go get github.com/stretchr/testify/assert

A biblioteca supracitada já vem com asserts prontos para serem usados como `assert.Equal(t, expected, actual)`, semelhante a diversas implementações nativas em frameworks de testes em outras linguagens. Depois de dar o comando acima é necessária importar no começo do código. 

    package main

    import (
        "errors"
        "reflect"
        "testing"

        "github.com/stretchr/testify/assert"
    )

    //Esta função é um helper criado para ser usado logo abaixo
    func TestContaPalavras(t *testing.T) {
        assertEqualMaps := func(t *testing.T, got, want map[string]int) {
            t.Helper()
            if !reflect.DeepEqual(got, want) { //checa se os Maps tem valores iguais
                t.Errorf("got %q want %q", got, want) //se sim esse comando indicará um erro
            }
        }

        t.Run("checa palavras de uma frase", func(t *testing.T) {
            gotDicionario, gotErr := contaPalavras("Hasta la vista baby vista la baby hasta la")
            wantDicionario := map[string]int{
                "Hasta": 1,
                "baby":  2,
                "la":    3,
                "vista": 2,
                "hasta": 1,
            }
            
            //modo clássico de asserts em go(não usadas a princípio)
            if gotErr != nil {
                t.Errorf("got %q want %q", gotErr, "nil")
            }
            assertEqualMaps(t, gotDicionario, wantDicionario) //uso do helper acima
        })

        t.Run("checa se acusa erro quando não há frase", func(t *testing.T) {
            gotDicionario, gotErr := contaPalavras("")
            wantErr := errors.New("o dicionario está vazio; certifique se digitou alguma frase")
            assert.Equal(t, 0, len(gotDicionario)) //assert requerida ao importar a biblioteca
            assert.Equal(t, wantErr, gotErr) //biblioteca pode ser requerida com o comando go get
        })
    }

No TestContaPalavras() podemos perceber que todos os passos acima foram seguidos, desde o nome do arquivo até o parâmetro da função.

A função teste verifica se de fato o que a função contaPalavras() retorna o que é o esperado que no caso é o objeto map[string]int com as configurações guardada na variável wantDicionario.

A comparação é feita a partir de `got` com `want` e caso venha a falhar um print é utilizando a partir de  `t.Errof()` ou `assert.Equal(...)`(que abstrai as funções de erro) para mostrar a diferença entre o recebido e o esperado.

## Passo 2.2: Rodando testes em Go

Para rodar o test deve ser utilizado `go test` porém ainda temos outras opções como o `go test -v` para ver com mais detalhes(verbose).

#### go test
Aqui está o comportamento do `go test`:

    PASS
    ok      contaPalavras  0.599s

#### go test -v
Aqui vemos como se comporta o `go test -v`:


    === RUN   TestContaPalavras
    === RUN   TestContaPalavras/checa_palavras_de_uma_frase
    === RUN   TestContaPalavras/checa_se_acusa_erro_quando_não_há_frase
    --- PASS: TestContaPalavras (0.00s)
    --- PASS: TestContaPalavras/checa_palavras_de_uma_frase (0.00s)
    --- PASS: TestContaPalavras/checa_se_acusa_erro_quando_não_há_frase (0.00s)
    PASS
    ok      contaPalavras  0.722s

#### go test -run

Caso tenhamos mais que uma função de teste podemos usar o comando 

    go test -run <NomedaFuncaodeTeste>

para rodar uma específica.

#### gotest
Uma opção para customizar o terminal é utilizar a biblioteca **gotest**(desta vez tudo junto) que sinaliza os testes que passarem pela cor verde e os que não pela cor vermelha:

    go get -u github.com/rakyll/gotest

Uma vez executado este comando você pode aplicar o comando **gotest** ou **gotest -v**(análagos aos comandos anteriores) para visualizar se os testes passaram ou não com cores.

# Passo 3: Test Coverage

É possível mensurar nossa statement coverage com o comando 

    go test -cover

que retornará este valor em porcentagem.

    PASS
    coverage: 61.5% of statements
    ok      
    contaPalavras  0.820s
    
Muitas vezes enquanto escrevendo testes não conseguiremos uma cobertura que nos irá satisfazer para determinda feature. Um jeito de checar quais statements específicos ficaram faltando ser cobertos é utilizando um comando de cobertura visual que gerará um html sinalizando statements os quais não foram cobertos. Para isso nós iremos usar o comando:

    go test -coverprofile="cover.txt"

Este comando acima colocará os dados sobre a cobertura em um arquivo específico(poderia ser de outro formato). Em seguida executamos:

	go test cover -html="cover.txt" -o cover.html

Por fim, uma página html será gerada com highlighting verde para os statements cobertos e vermelha para os não cobertos.

Adicionando a função Dummy e seu teste em seus respectivos arquivos deixará a explicação ficará mais clara.

### dummy.go

    func Dummy(x int) bool {
	  if x > 5 {
	    return true
	  } else {
	    return false
	  }
    }

### dummy_test.go

    func TestDummy(t *testing.T) {
	  got := Dummy(6)
	  want := true

	  if got != want {
	    t.Errorf("got %t want %t", got, want)
	  }
    }

Muitas vezes enquanto escrevendo testes não conseguiremos uma cobertura que nos irá satisfazer para determinda feature. Um jeito de checar quais statements específicos ficaram faltando ser cobertos é utilizando um comando de cobertura visual que gerará um html sinalizando statements os quais não foram cobertos. Para isso nós iremos usar o comando:

    go test -coverprofile="cover.txt"

Este comando acima colocará os dados sobre a cobertura em um arquivo específico(poderia ser de outro formato). Perceba que um novo arquivo será gerado em seu workspace. Em seguida executamos:

    go tool -cover -html="cover.txt" -o cover.html

Por fim, uma página html será gerada com highlighting verde para os statements cobertos e vermelha para os não cobertos. Caso você opte por abrir em seu browser:

![alt text](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/images/statement_coverage.png "Statement Coverage")

Após rodar os comandos fica notável que o caso do `else` não está sendo coberto no teste indicando que o mesmo pode ser melhorado.

# Passo 4: Benchmarks

É possível realizar bechmarks dos testes em Golang para checar o seu tempo de execução. Os resultados poderão ser coletados pelo analista/engenheiro de testes afim de melhorar o desempenho dos mesmos.

 - Colocamos a função de benchmark junto com nossos testes.

- A função de teste deve começar com **Benchmark** seguido do seu nome. 

        BenchmarkContaPalavras(...)

- O único argumento da função de teste recebe é **b  *testing.B**. 

        BenchmarkContaPalavras(b *testing.B)

## Passo 4.1: Exemplo de benchmark utilizando GO
Aqui iremos demonstrar um exemplo fazendo o benchmark da função **`contaPalavras()`** que será executada N vezes onde este valor é definido pelo framework. Adicionamos a função no nosso arquivo **contaPalavras_test.go**:

    func BenchmarkcontaPalavras(b *testing.B) {
        for i := 0; i < b.N; i++ {
            contaPalavras(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
                        sed do eiusmod tempor incididunt ut labore et dolore magna 
                        aliqua. Ac turpis egestas sed tempus urna et pharetra. Duis 
                        at consectetur lorem donec massa sapien faucibus. Viverra ipsum 
                        nunc aliquet bibendum enim. Dui accumsan sit amet nulla facilisi 
                        morbi tempus iaculis. Blandit volutpat maecenas volutpat blandit 
                        aliquam etiam erat. Augue ut lectus arcu bibendum at. Pharetra diam 
                        sit amet nisl suscipit adipiscing bibendum. Pharetra diam sit amet 
                        nisl suscipit adipiscing bibendum est ultricies. Dolor sit amet consectetur 
                        adipiscing elit pellentesque.`)
        }
    }

Para rodar o benchmark deve ser utilizado `go test -bench=.` se usamos sistemas Unix ou se `go test -bench="."` se Windows Powershell. Abaixo um exemplo de execução da função acima:

    goos: windows
    goarch: amd64
    BenchmarkContaPalavras-4           60949             19087 ns/op
    PASS
    ok     
    contaPalavras  2.610s



