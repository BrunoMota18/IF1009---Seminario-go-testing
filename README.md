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
