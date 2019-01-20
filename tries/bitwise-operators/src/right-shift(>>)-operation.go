package main

import "fmt"

func main()  {

    /*
        Operador shift right (>>)1
        esse operador eh bem legal, ele empurra todos os bits para uma posiçao a direita quantas vezes forem sinalizadas,
        e onde haviam 1's viram zeros a esquerda

        Segue alguns exemplos:

        11000>>1 = 01100 // 24 >>1 = 12
        11000>>2 = 00110 // 24 >>2 = 6
        11000>>3 = 00011 // 24 >>3 = 3
        11001>>1 = 01100 // 25 >>1 = 12 -> ele empura o ultimo 1 que seria o resto

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
    */
    var val byte = 128

    fmt.Printf("(%2d >>1 = %2d)\n", val, val>>1)
    for i :=0; i <= 8; i++{
        fmt.Printf(">>%d %08b  0x%02X  %03d\n", i, val, val, val)
        val = val>>1
    }
}
