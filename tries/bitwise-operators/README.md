## Operadores Bitwise (bit-a-bit) GoLang:
Houve um tempo em que os recursos de hardware era escasso e o programador precisava aproveitar cada bit que estava disponível mara manter a usabilidade do programa.
Existem em linguagem como C e C++ o recurso de operações bit-a-bit que eram usadas em várias técnicas de programação.
Elas consistem em usar uma represantação binária de um dado (tudo é binário em computação) e usa as operações em cada bit da representação.

GoLang não abriu mão de suportar esses tipos de operações e podem ser ultilizadas em algorítimos 
para melhorar ainda mais a performace. ou resolver problemas específicos.

Operações de bitwise são muito usadas em sistemas de criptografia por exemplo.

As existentes em Golang São:
```
Operadores bitwise  
&    bitwise AND  
|    bitwise OR 
^    bitwise XOR 
^    bitwise NOT (unário) 
&^   bitwise AND NOT 
<<   bitwise left shift  
>>   bitwise right shift 
```

### Operador AND (&):
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
```console
[notim@Lenovo-ideapad:~]$
(134 & 120 = 0)
  10000110 > 0x86 > 134
& 01111000 > 0x78 > 120
= 00000000 > 0x00 > 000
```
```console
[notim@Lenovo-ideapad:~]$
(140 & 12 = 12)
  10001100 > 0x8C > 140
& 00001100 > 0x0C > 012
= 00001100 > 0x0C > 012
```
```console
[notim@Lenovo-ideapad:~]$
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
```console
[notim@Lenovo-ideapad:~]$
(134 | 118 = 246)
  10000110 > 0x86 > 134
| 01110110 > 0x76 > 118
= 11110110 > 0xF6 > 246
```
```console
[notim@Lenovo-ideapad:~]$
(140 | 12 = 140)
  10001100 > 0x8C > 140
| 00001100 > 0x0C > 012
= 10001100 > 0x8C > 140
```
```console
[notim@Lenovo-ideapad:~]$
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

#### Código:
```go
func main(){
    var val1 byte = 0xAA
    var val2 byte = 0x63
    var xor  byte = val1 ^ val2

    fmt.Printf("(%d ^ %d = %d)\n",val1, val2, xor)
    fmt.Printf("  %8b > 0x%2X > %3d\n", val1, val1, val1)
    fmt.Printf("^ %8b > 0x%2X > %3d\n", val2, val2, val2)
    fmt.Printf("= %8b > 0x%2X > %3d\n", xor, xor, xor)
}
```
#### Exemplos:
```console
[notim@Lenovo-ideapad:~]$
(134 ^ 118 = 240)
  10000110 > 0x86 > 134
^ 01110110 > 0x76 > 118
= 11110000 > 0xF0 > 240
```
```console
[notim@Lenovo-ideapad:~]$
(140 ^ 12 = 128)
  10001100 > 0x8C > 140
^ 00001100 > 0x0C > 012
= 10000000 > 0x80 > 128
```
```console
[notim@Lenovo-ideapad:~]$
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

#### Código:
```go
func main(){
    var val    byte  = 0xA1
    var not    byte  = ^val // val ^ val

    fmt.Printf("(^%d = %d)\n",val, not)
    fmt.Printf("   %08b > 0x%02X > %03d\n", val, val, val)
    fmt.Printf("^  %08b > 0x%02X > %03d\n", not, not, not)
}
```
#### Exemplos:
```console
[notim@Lenovo-ideapad:~]$
(^255 = 0)
  11111111 > 0xFF > 255
^ 00000000 > 0x00 > 000
```
```console
[notim@Lenovo-ideapad:~]$
(^175 = 80)
  10101111 > 0xAF > 175
^ 01010000 > 0x50 > 080
```
```console
[notim@Lenovo-ideapad:~]$
(^241 = 14)
  11110001 > 0xF1 > 241
^ 00001110 > 0x0E > 014
```
```console
[notim@Lenovo-ideapad:~]$
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
#### Codigo
```go
func main(){
    var val1    byte = 0x67
    var val2    byte = 0x32

    var andnot byte = val1
    andnot &^= val2

    fmt.Printf("(%d &^= %d = %d)\n",val1,val2, andnot)
    fmt.Printf("a      %08b > 0x%02X > %03d\n", val1, val1, val1)
    fmt.Printf("b      %08b > 0x%02X > %03d\n", val2, val2, val2)
    fmt.Printf("^b     %08b > 0x%02X > %03d\n", ^val2, ^val2, ^val2)
    fmt.Printf("a&(^b) %08b > 0x%02X > %03d\n", val1 & ^val2, val1 & ^val2, val1 & ^val2)
}
```
#### Exemplos:
```console
[notim@Lenovo-ideapad:~]$
(175 &^= 80 = 175)
a      10101111 > 0xAF > 175
b      01010000 > 0x50 > 080
^b     10101111 > 0xAF > 175
a&(^b) 10101111 > 0xAF > 175
```
```console
[notim@Lenovo-ideapad:~]$
(255 &^= 0 = 255)
a      11111111 > 0xFF > 255
b      00000000 > 0x00 > 000
^b     11111111 > 0xFF > 255
a&(^b) 11111111 > 0xFF > 255
```
```console
[notim@Lenovo-ideapad:~]$
(103 &^= 50 = 69)
a      01100111 > 0x67 > 103
b      00110010 > 0x32 > 050
^b     11001101 > 0xCD > 205
a&(^b) 01000101 > 0x45 > 069
```

### Operador shift right (>>)
Esse operador eh bem legal, ele empurra todos os bits para uma posiçao a direita quantas vezes forem sinalizadas,
e onde haviam 1's viram zeros a esquerda.

```go
func main(){
    var val byte = 128
    
    fmt.Printf("(%2d >>1 = %2d)\n", val, val>>1)
    for i :=0; i <= 8; i++{
        fmt.Printf(">>%d %08b  0x%02X  %03d\n", i, val, val, val)
        val = val>>1
    }
}
```
Seguem alguns exemplos:
```console
[notim@Lenovo-ideapad]$
11000>>1 = 01100 | 24 >>1 = 12
11000>>2 = 00110 | 24 >>2 = 6
11000>>3 = 00011 | 24 >>3 = 3
11001>>1 = 01100 | 25 >>1 = 12 ==> ele empura o ultimo 1 que seria o resto
```
```console
[notim@Lenovo-ideapad]$
(255 >>1 = 127)
    11111111  0xFF  255
>>1 01111111  0x7F  127
>>2 00111111  0x3F  063
>>3 00011111  0x1F  031
>>4 00001111  0x0F  015
>>5 00000111  0x07  007
>>6 00000011  0x03  003
>>7 00000001  0x00  001
>>8 00000000  0x00  000
```
```console
[notim@Lenovo-ideapad]$
(128 >>1 = 64)
    10000000  0x80  128
>>1 01000000  0x40  064
>>2 00100000  0x20  032
>>3 00010000  0x10  016
>>4 00001000  0x08  008
>>5 00000100  0x04  004
>>6 00000010  0x02  002
>>7 00000001  0x00  001
>>8 00000000  0x00  000
```

ref :  
[bit-hacking-with-go](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)   
[Bitwise-Escovando-os-bits](https://www.vivaolinux.com.br/artigo/Bitwise-Escovando-os-bits)   

#### notim @ 2019 <joao.paulino@yahoo.com.br>