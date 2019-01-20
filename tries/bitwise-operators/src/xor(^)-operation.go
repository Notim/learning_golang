package src

import (
    "fmt"
)

func main() {
    /**
     * Operador XOR (^) :
     * Compara Bit por Bit e se o valor for igual em ambos ele zera o valor [na posiçao do bit]
     *
     * Tabela verdade
     *  1 ^ 1 = 0
     *  0 ^ 0 = 0
     *  0 ^ 1 = 1
     *  1 ^ 0 = 1
     * Exemplos:
     *
     * (134 ^ 118 = 240)
     *    10000110 > 0x86 > 134
     *  ^ 01110110 > 0x76 > 118
     *  = 11110000 > 0xF0 > 240
     *
     * (140 ^ 12 = 128)
     *    10001100 > 0x8C > 140
     *  ^ 00001100 > 0x0C >  12
     *  = 10000000 > 0x80 > 128
     *
     * (170 ^ 99 = 201)
     *    10101010 > 0xAA > 170
     *  ^ 01100011 > 0x63 >  99
     *  = 11001001 > 0xC9 > 201
     *
    **/

    val1 := 0xAA
    val2 := 0x63
    xor  := val1 ^ val2

    fmt.Printf("(%d ^ %d = %d)\n",val1, val2, xor)
    fmt.Printf("  %8b > 0x%2X > %3d\n", val1, val1, val1)
    fmt.Printf("^ %8b > 0x%2X > %3d\n", val2, val2, val2)
    fmt.Printf("= %8b > 0x%2X > %3d\n", xor, xor, xor)
}