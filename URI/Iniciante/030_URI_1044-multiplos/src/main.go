package main

	import io "fmt"
	import sys "os"

	func main() {
		var  A, B int

		io.Scanf("%d %d", &A, &B)

		A, B = swap(A, B)

		if  ((A % B) == 0){
			io.Println("Sao Multiplos")
		}else{
			io.Println("Nao sao Multiplos")
		}
		sys.Exit(0)
	}
	func swap(A ,B int) (sa , sb int){
		if B > A {
			swap := A
			A = B
			B = swap
		}
		sa, sb = A, B

		return
	}