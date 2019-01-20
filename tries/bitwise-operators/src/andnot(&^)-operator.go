package main

import (
    "fmt"
)

func main() {
    /**
     * Operador AND NOT &^:
     * Na verdade isso eh um atalho para os Operadores AND XOR(or)
     * serve para comparar os bits com AND depois de inverter os bits
     * operacao real (xa & ^xb)
     *
     * Tabela Verdade
     * (a1 == (^b0 = 1)) = 1
     * (a0 == (^b0 = 1)) = 0
     * (a1 == (^b1 = 0)) = 0
     * (a0 == (^b1 = 0)) = 1
     *
     *
     * Exemplos:
     * (175 &^= 80 = 175)
     * a      10101111 > 0xAF > 175
     * b      01010000 > 0x50 > 080
     * ^b     10101111 > 0xAF > 175
     * a&(^b) 10101111 > 0xAF > 175
     *
     * (255 &^= 0 = 255)
     * a      11111111 > 0xFF > 255
     * b      00000000 > 0x00 > 000
     * ^b     11111111 > 0xFF > 255
     * a&(^b) 11111111 > 0xFF > 255
     *
     * (103 &^= 50 = 69)
     * a      01100111 > 0x67 > 103
     * b      00110010 > 0x32 > 050
     * ^b     11001101 > 0xCD > 205
     * a&(^b) 01000101 > 0x45 > 069
     *
    **/

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