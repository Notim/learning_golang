## Operadores Bitwise (bit-a-bit) GoLang:
Em um tempo onde o hardware era escasso e o programador precisava aproveitar cada bit que estava disponível
existia o recurso de operações binárias que eram usadas em várias técnicas de programação.

GoLang não abriu mão de suportar esses tipos de operações e podem ser ultilizadas em algorítimos 
para melhorar ainda mais a performace. sendo eles:
```
Operadores bitwise 
 &   bitwise AND
 |   bitwise OR
 ^   bitwise XOR
 ^   bitwise NOT (unário)
&^   AND NOT
<<   left shift
>>   right shift
```



### Operador AND (&) :
Compara Bit por Bit se o valor é igual em ambos, caso sim ele retorna 1 senão 0 na posiçao do bit
```
Tabela verdade
    1 & 1 = 1
    1 & 0 = 0
    0 & 1 = 0
    0 & 0 = 1
```
#### Código:
```go
func main(){
    var val1 byte = 0xAA
    var val2 byte = 0x63
    var and  byte = val1 & val2

    fmt.Printf("(%d & %d = %d)\n",val1, val2, and)
    fmt.Printf("  %8b > 0x%2X > %3d\n", val1, val1, val1)
    fmt.Printf("& %8b > 0x%2X > %3d\n", val2, val2, val2)
    fmt.Printf("= %8b > 0x%2X > %3d\n", and, and, and)
}
```
#### Exemplos:
```
foo@bar ~:
(134 & 120 = 0)
  10000110 > 0x86 > 134
& 01111000 > 0x78 > 120
= 00000000 > 0x00 > 0
```
```
foo@bar ~:
(140 & 12 = 12)
  10001100 > 0x8C > 140
& 00001100 > 0x0C > 12
= 00001100 > 0x0C > 12
```
```
foo@bar ~:
(170 & 99 = 34)
  10101010 > 0xAA > 170
& 01100011 > 0x63 > 99
= 00100010 > 0x22 > 34
```



### Operador OR (|) :
Compara Bit por Bit e se o valor de pelo menos uma posiçao entre os binarios eh 1
Ele retorna 1 senao 0 na posiçao do bit
```
Tabela verdade:
    0 | 0 = 0
    0 | 1 = 1
    1 | 0 = 1
    1 | 1 = 1
```
#### Codigo:
```go
func main(){
    var val1 byte = 0xAA
    var val2 byte = 0x63
    var or   byte = val1 | val2

    fmt.Printf("(%d & %d = %d)\n",val1, val2, or)
    fmt.Printf("  %8b > 0x%2X > %3d\n", val1, val1, val1)
    fmt.Printf("| %8b > 0x%2X > %3d\n", val2, val2, val2)
    fmt.Printf("= %8b > 0x%2X > %3d\n", or, or, or)
}
```

#### Exemplos:
```
foo@bar ~:
(134 | 118 = 246)
  10000110 > 0x86 > 134
| 01110110 > 0x76 > 118
= 11110110 > 0xF6 > 246
```
```
foo@bar ~:
(140 | 12 = 140)
  10001100 > 0x8C > 140
| 00001100 > 0x0C >  12
= 10001100 > 0x8C > 140
```
```
foo@bar ~:
(170 | 99 = 235)
  10101010 > 0xAA > 170
| 01100011 > 0x63 >  99
= 11101011 > 0xEB > 235
``` 
