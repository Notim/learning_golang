package main

import (
    "fmt"
)

/*
    bitwise operators

     &   bitwise AND
     |   bitwise OR
     ^   bitwise XOR
    &^   AND NOT
    <<   left shift
    >>   right shift
*/
func main() {
    var byte1 byte = 0xAA // 170(10) 0x0000AA(16) 10101010(2)
    var byte2 byte = 0xFF // 255(10) 0x0000FF(16) 11111111(2)
    var byte3 byte = 0x05 // 005(10) 0x000005(16) 00000101(2)

    fmt.Printf("%3d > %2X > %8b\n",byte1, byte1, byte1)
    fmt.Printf("%3d > %2X > %8b\n",byte2, byte2, byte2)
    fmt.Printf("%3d > %2X > %8b\n",byte3, byte3, byte3)

    /**
     * Operador AND (&) :
     * Compara Bit por Bit se o valor eh igual, caso sim ele retorna 1 senao 0 [na posiçao do bit]
     * o mais daora desse tipo de operaçao, eh que os resultados nao fazem sentido kk
     *
     * Tabela verdade
     *  1 & 1 = 1
     *  1 & 0 = 0
     *  0 & 1 = 0
     *  0 & 0 = 1
     *
     * Exemplos:
     *
     * (134 & 120 = 0)
     *    10000110 > 0x86 > 134
     *  & 01111000 > 0x78 > 120
     *  = 00000000 > 0x00 > 0
     *
     * (140 & 12 = 12)
     *     10001100 > 0x8C > 140
     *   & 00001100 > 0x0C > 12
     *   = 00001100 > 0x0C > 12
     *
     * (170 & 99 = 34)
     *    10101010 > 0xAA > 170
     *  & 01100011 > 0x63 > 99
     *  = 00100010 > 0x22 > 34
     *
    **/

    val1 := 0xAA
    val2 := 0x63
    and := val1 & val2

    fmt.Printf("(%d & %d = %d)\n",val1, val2, and)
    fmt.Printf("  %8b > 0x%2X > %3d\n", val1, val1, val1)
    fmt.Printf("& %8b > 0x%2X > %3d\n", val2, val2, val2)
    fmt.Printf("= %8b > 0x%2X > %3d\n", and, and, and)
}