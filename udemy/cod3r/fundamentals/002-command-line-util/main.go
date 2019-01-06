/*
ok, the code below means that go has its own package manager! (more or less)
it gets the source code of the github repository and places it in $ GOPATH (env variable).
the default is "~/user/go" in unix-like and "C:\Users\username\go" in win
    ~$ go get -u "github.com/go-sql-driver/mysql"

that command verify syntax
    ~$ go vet [source path]

that command make bin from your source
    ~$ go build [source path]

that command make bin and run from your source
    ~$ go build [source path]

for more information see https://golang.org or put "~$ go --help" on terminal
*/

package main

import "fmt"

func main() {
	fmt.Println("please, see the source to understand the chapter!!")
}
