package main
	import io "fmt"

	func main(){
		var opt string

		io.Scanf("%s", &opt)

		switch opt {
			case "61" :
				io.Println("Brasilia")
				break
			case "71" :
				io.Println("Salvador")
				break
			case "11" :
				io.Println("Sao Paulo")
				break
			case "21" :
				io.Println("Rio de Janeiro")
				break
			case "32" :
				io.Println("Juiz de Fora")
				break
			case "19" :
				io.Println("Campinas")
				break
			case "27" :
				io.Println("Vitoria")
				break
			case "31" :
				io.Println("Belo Horizonte")
				break
			default:
				io.Println("DDD nao cadastrado")
				break
		}
	}