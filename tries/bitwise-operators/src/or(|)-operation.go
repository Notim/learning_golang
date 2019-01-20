package main

import (
    "fmt"
)
func main() {
    /**
     * Operador OR (|) :
     * Compara Bit por Bit e se o valor de pelo menos uma posiçao entre os binarios eh 1
     * Ele retorna 1 senao 0 [na posiçao do bit]
     *
     * Tabela verdade:
     *  0 | 0 = 0
     *  0 | 1 = 1
     *  1 | 0 = 1
     *  1 | 1 = 1
     *
     * Exemplos:
     *
     * (134 | 118 = 246)
     *    10000110 > 0x86 > 134
     *  | 01110110 > 0x76 > 118
     *  = 11110110 > 0xF6 > 246
     *
     * (140 | 12 = 140)
     *    10001100 > 0x8C > 140
     *  | 00001100 > 0x0C >  12
     *  = 10001100 > 0x8C > 140
     *
     * (170 | 99 = 235)
     *    10101010 > 0xAA > 170
     *  | 01100011 > 0x63 >  99
     *  = 11101011 > 0xEB > 235
     *
    **/

    val1 := 0x86
    val2 := 0x76
    or := val1 | val2

    fmt.Printf("(%d | %d = %d)\n",val1, val2, or)
    fmt.Printf("  %8b > 0x%2X > %3d\n", val1, val1, val1)
    fmt.Printf("| %8b > 0x%2X > %3d\n", val2, val2, val2)
    fmt.Printf("= %8b > 0x%2X > %3d\n", or, or, or)
}