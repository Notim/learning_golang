package main

import "fmt"

func main()  {

    /*
        Operador shift left (>>)
        Ele empurra todos os bits para uma posiçao a esquerda quantas vezes forem sinalizadas,

        Segue alguns exemplos:

        00001<<1 =    00010 //  02 <<1 = 4
        00001<<2 =    00100 //  02 <<1 = 8
        00001<<3 =    01000 //  02 <<1 = 16
        00001<<4 =    10000 //  02 <<1 = 32
     10000001<<1 = 00000010 // 129 <<1 = 02 -> ele joga fora o 1 que compoe o 128 pois a variavel de exemplo s� tera espaço para 255 numeros(uint8)

        (255 <<1 = 254)
        <<0 11111111  0xFF  255
        <<1 11111110  0xFE  254
        <<2 11111100  0xFC  252
        <<3 11111000  0xF8  248
        <<4 11110000  0xF0  240
        <<5 11100000  0xE0  224
        <<6 11000000  0xC0  192
        <<7 10000000  0x80  128
        <<8 00000000  0x00  000

        ( 1 <<1 =  2)
        <<0 00000001  0x01  001
        <<1 00000010  0x02  002
        <<2 00000100  0x04  004
        <<3 00001000  0x08  008
        <<4 00010000  0x10  016
        <<5 00100000  0x20  032
        <<6 01000000  0x40  064
        <<7 10000000  0x80  128
        <<8 00000000  0x00  000
    */
    var val byte = 0x01

    fmt.Printf("(%2d <<1 = %2d)\n", val, val<<1)
    for i := 0; i <= 8; i++{
        fmt.Printf("<<%d %08b  0x%02X  %03d\n", i, val, val, val)
        val = val<<1
    }
}
