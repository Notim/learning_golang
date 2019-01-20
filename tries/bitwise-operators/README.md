### Operador AND (&) :
Compara Bit por Bit se o valor é igual, caso sim ele retorna 1 senao 0 na posiçao do bit
o mais daora desse tipo de operaçao, é que os resultados não fazem sentido kk

```
Tabela verdade
    1 & 1 = 1
    1 & 0 = 0
    0 & 1 = 0
    0 & 0 = 1
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
