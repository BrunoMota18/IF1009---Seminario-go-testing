# IF1009-Seminario-go-testing
Seminário com demo para disciplina de Teste de Software do CIn-UFPE.<br> 
Este repositório tem como objetivo demonstrar funcionalidades da biblioteca 
de testes **go-testing**. 

---

## [contaPalavras](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/contaPalavras/contaPalavras.go)
A função **contaPalavras** recebe uma string e retorna a quantidade de palavras dessa frase.

## [contaPalavras_test](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/contaPalavras/contaPalavras_test.go)
O teste checa as palavras de uma frase e verifica se há algum erro quando não há nenhuma frase.

## [contaPalavras_benchmark](https://github.com/BrunoMota18/IF1009-Seminario-go-testing/blob/master/contaPalavras/contaPalavras_benchmark.go)
O benchmark recebe uma frase e entra em um loop com um limite delimitado pela própria função onde é possível realizar a análise de segundos por operação.

# Testes com Go!
Ao fim desse tutorial você estará por dentro dos conceitos básicos de testes em Go.


# Instalação

Inicialmente será necessário a instalação no seu computador.
 A documentação original pode ser encontrada [aqui](http://www.golangbr.org/doc/instalacao), assim como o instalador.


# Testando o funcionamento no seu computador

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

# Como fazer testes com Golang
Escrever testes em Go é como escrever qualquer outra função, porém teremos algumas regras para o funcionamento correto.
- Como em Go já existe um recurso de testes nativo da linguagem, você só deve importar **testing**.

        import "testing"

- O nome do arquivo deve sempre ser o nome do seu arquivo seguido de **_test.go**. 

        touch contaPalavras_test.go

- A função de teste deve começar com **Test** seguido do seu nome. 

        TestContaPalavras(...)

- O único argumento da função de teste recebe é **t  *testing.T**. 

        TestContaPalavras(t *testing.T)

## Exemplo de teste utilizando GO
Aqui iremos demonstrar um exemplo da função **`contaPalavras()`** que tem com objetivo retornar um dicionario com as frequências de palavras de uma frase.
### hello.go
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

### hello_test.go

    package main

    import (
        "errors"
        "reflect"
        "testing"

        "github.com/stretchr/testify/assert"
    )

    func TestContaPalavras(t *testing.T) {
        assertEqualMaps := func(t *testing.T, got, want map[string]int) {
            t.Helper()
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got %q want %q", got, want)
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

            if gotErr != nil {
                t.Errorf("got %q want %q", gotErr, "nil")
            }
            assertEqualMaps(t, gotDicionario, wantDicionario)
        })

        t.Run("checa se acusa erro quando não há frase", func(t *testing.T) {
            gotDicionario, gotErr := contaPalavras("")
            wantErr := errors.New("o dicionario está vazio; certifique se digitou alguma frase")
            assert.Equal(t, 0, len(gotDicionario))
            assert.Equal(t, wantErr, gotErr)
        })
    }

No TestContaPalavras() podemos perceber que todos os passos acima foram seguidos, desde o nome do arquivo até o parâmetro da função.

A função teste verifica se de fato o que a função contaPalavras() retorna o que é o esperado que no caso é o objeto map[string]int com as configurações guardada na variável wantDicionario.

A comparação é feita a partir de `got` com `want` e caso venha a falhar um print é utilizando a partir de  `t.Errof()` ou `assert.Equal(...)` para mostrar a diferença entre o recebido e o esperado.

Para rodar o test deve ser utilizado `go test` porém ainda temos outras opções como o `go test -v` para ver com mais detalhes(verbose).

##### go test
Aqui está o comportamento do `go test`:

    PASS
    ok      contaPalavras  0.599s

##### go test -v
Aqui vemos como se comporta o `go test -v`:


    === RUN   TestContaPalavras
    === RUN   TestContaPalavras/checa_palavras_de_uma_frase
    === RUN   TestContaPalavras/checa_se_acusa_erro_quando_não_há_frase
    --- PASS: TestContaPalavras (0.00s)
    --- PASS: TestContaPalavras/checa_palavras_de_uma_frase (0.00s)
    --- PASS: TestContaPalavras/checa_se_acusa_erro_quando_não_há_frase (0.00s)
    PASS
    ok      contaPalavras  0.722s


