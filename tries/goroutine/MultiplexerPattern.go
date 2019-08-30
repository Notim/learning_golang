package main

func Forward(origem <- chan string, destino chan string){
    for {
        destino <- <-origem
    }
}

func Mix(inside1, inside2 <-chan string) <-chan string  {
    chanel := make(chan string)

    go Forward(inside1, chanel)
    go Forward(inside2, chanel)

    return chanel
}