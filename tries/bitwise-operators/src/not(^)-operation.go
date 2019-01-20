package src

import (
    "fmt"
)

func main() {
    /**
     * Operador (NOT) > XOR (^):
     * O Operador NOT serve para negar todos os bits de um inteiro, o que era 0 passa a ser 1
     * e 1 passa a ser 0
     * O Golang nao possui um operador unario especifico para operacoes de NOT como "~" em C, C++ e afins
     * Porem o Go pode ser usado o XOR ^ de modo unario dessa forma ^x, ele faz o xor com seu proprio valor
     * sendo assim a açao eh a mesma do NOT em outras linguagens.
     * usa-se "^X" mas o compilador faz desse modo "X ^ X",
     * Quando a unidade eh a mesma ele retora todos os bits inversos,
     *
     * Tabela Verdade
     *  ^1 = 0
     *  ^0 = 1
     *
     * Exemplos:
     *
     * (^255 = 0)
     *   11111111 > 0xFF > 255
     * ^ 00000000 > 0x00 > 000
     *
     * (^175 = 80)
     *   10101111 > 0xAF > 175
     * ^ 01010000 > 0x50 > 080
     *
     * (^241 = 14)
     *   11110001 > 0xF1 > 241
     * ^ 00001110 > 0x0E > 014
     *
     * (^161 = 94)
     *    10100001 > 0xA1 > 161
     *  ^ 01011110 > 0x5E > 094
     *
    **/

    var val    byte  = 0xA1
    var not    byte  = ^val // val ^ val

    fmt.Printf("(^%d = %d)\n",val, not)
    fmt.Printf("   %08b > 0x%02X > %03d\n", val, val, val)
    fmt.Printf("^  %08b > 0x%02X > %03d\n", not, not, not)
}