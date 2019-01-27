package enc

import (
    "encoding/base64"
    "fmt"
)

func Base64Encode(src []byte) []byte {
    return []byte(base64.StdEncoding.EncodeToString(src))
}

func Base64Decode(src []byte) []byte {
    b, er := base64.StdEncoding.DecodeString(string(src))
    if er!= nil {
        fmt.Println("Deu bosta: ", er)
    }
    return b
}

func main() {
    hello := "pass"
    debyte := Base64Encode([]byte(hello))
    fmt.Println(debyte)

    enbyte := Base64Decode(debyte)
    fmt.Println(string(enbyte))
}
