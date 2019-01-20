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
[notim@Lenovo-ideapad:~]
(134 & 120 = 0)
  10000110 > 0x86 > 134
& 01111000 > 0x78 > 120
= 00000000 > 0x00 > 000
```
```
[notim@Lenovo-ideapad:~]
(140 & 12 = 12)
  10001100 > 0x8C > 140
& 00001100 > 0x0C > 012
= 00001100 > 0x0C > 012
```
```
[notim@Lenovo-ideapad:~]
(170 & 99 = 34)
  10101010 > 0xAA > 170
& 01100011 > 0x63 > 099
= 00100010 > 0x22 > 034
```



### Operador OR (|) :
Compara Bit por Bit e se o valor de pelo menos uma posiçao entre os binarios é 1
Ele retorna 1 senão 0 na posição do bit
```
Tabela verdade:
    0 | 0 = 0
    0 | 1 = 1
    1 | 0 = 1
    1 | 1 = 1
```
#### Código:
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
[notim@Lenovo-ideapad:~]
(134 | 118 = 246)
  10000110 > 0x86 > 134
| 01110110 > 0x76 > 118
= 11110110 > 0xF6 > 246
```
```
[notim@Lenovo-ideapad:~]
(140 | 12 = 140)
  10001100 > 0x8C > 140
| 00001100 > 0x0C > 012
= 10001100 > 0x8C > 140
```
```
[notim@Lenovo-ideapad:~]
(170 | 99 = 235)
  10101010 > 0xAA > 170
| 01100011 > 0x63 > 099
= 11101011 > 0xEB > 235
``` 


### Operador XOR (^) :
Compara Bit por Bit e se o valor for igual em ambos ele zera o valor na posiçao do bit
```
Tabela verdade
    1 ^ 1 = 0
    0 ^ 0 = 0
    0 ^ 1 = 1
    1 ^ 0 = 1
```

```go
func main(){
    var val1 uint8 = 0xAA
    var val2 uint8 = 0x63
    var xor  uint8 = val1 ^ val2

    fmt.Printf("(%d ^ %d = %d)\n",val1, val2, xor)
    fmt.Printf("  %8b > 0x%2X > %3d\n", val1, val1, val1)
    fmt.Printf("^ %8b > 0x%2X > %3d\n", val2, val2, val2)
    fmt.Printf("= %8b > 0x%2X > %3d\n", xor, xor, xor)
}
```
#### Exemplos:
```
[notim@Lenovo-ideapad:~]
(134 ^ 118 = 240)
  10000110 > 0x86 > 134
^ 01110110 > 0x76 > 118
= 11110000 > 0xF0 > 240
```
```
[notim@Lenovo-ideapad:~]
(140 ^ 12 = 128)
  10001100 > 0x8C > 140
^ 00001100 > 0x0C > 012
= 10000000 > 0x80 > 128
```
```
[notim@Lenovo-ideapad:~]
(170 ^ 99 = 201)
  10101010 > 0xAA > 170
^ 01100011 > 0x63 > 099
= 11001001 > 0xC9 > 201
```



### Operador NOT (XOR ^):
O Operador NOT serve para negar todos os bits de um inteiro, o que era 0 passa a ser 1
e 1 passa a ser 0
O Golang nao possui um operador unario especifico para operacoes de NOT como "~" em C, C++ e afins
Porem o Go pode ser usado o XOR ^ de modo unario dessa forma "^x", ele faz o xor com seu proprio valor
sendo assim a açao eh a mesma do NOT em outras linguagens.
usa-se "^X" mas o compilador faz desse modo "X ^ X", Quando a unidade eh a mesma ele retora todos os bits inversos,

```
Tabela Verdade
    ^1 = 0
    ^0 = 1
```

#### Exemplos:
```
[notim@Lenovo-ideapad:~]
(^255 = 0)
  11111111 > 0xFF > 255
^ 00000000 > 0x00 > 000
```
```
[notim@Lenovo-ideapad:~]
(^175 = 80)
  10101111 > 0xAF > 175
^ 01010000 > 0x50 > 080
```
```
[notim@Lenovo-ideapad:~]
(^241 = 14)
  11110001 > 0xF1 > 241
^ 00001110 > 0x0E > 014
```
```
[notim@Lenovo-ideapad:~]
(^161 = 94)
  10100001 > 0xA1 > 161
^ 01011110 > 0x5E > 094
```


### Operador AND NOT &^:
Na verdade isso eh um atalho para os Operadores AND XOR(or)
serve para comparar os bits com AND depois de inverter os bits
operacao real (xa & ^xb)

```
Tabela Verdade
    (a1 == (^b0 = 1)) = 1
    (a0 == (^b0 = 1)) = 0
    (a1 == (^b1 = 0)) = 0
    (a0 == (^b1 = 0)) = 1
```

#### Exemplos:
```
[notim@Lenovo-ideapad:~]
(175 &^= 80 = 175)
a      10101111 > 0xAF > 175
b      01010000 > 0x50 > 080
^b     10101111 > 0xAF > 175
a&(^b) 10101111 > 0xAF > 175
```
```
[notim@Lenovo-ideapad:~]
(255 &^= 0 = 255)
a      11111111 > 0xFF > 255
b      00000000 > 0x00 > 000
^b     11111111 > 0xFF > 255
a&(^b) 11111111 > 0xFF > 255
```
```
[notim@Lenovo-ideapad:~]
(103 &^= 50 = 69)
a      01100111 > 0x67 > 103
b      00110010 > 0x32 > 050
^b     11001101 > 0xCD > 205
a&(^b) 01000101 > 0x45 > 069
```