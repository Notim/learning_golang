package string

func (str String) Reversed() string{
    var reversed string

    for index := (str.Lenght()-1); index >= 0 ; index-- {
        reversed += string(str[index])
    }
    return reversed
}