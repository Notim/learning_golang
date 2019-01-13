The golang does not seem to make lambda expressions, but you can use a literal anonymous function, I wrote some examples when I was studying comparing the equivalent in JS, I hope it helps !!

### no args return string:
```golang
func() string {
    return "some String Value"
}
//Js similar: () => 'some String Value'
```
### with string args and return string
```golang
func(arg string) string {
    return "some String" + arg
}
//Js similar: (arg) => "some String Value" + arg
```
### no arguments and no returns (void)
```golang
func() {
   fmt.Println("Some String Value")
} 
//Js similar: () => {console.log("Some String Value")}
```

### no Arguments and no returns (void)
```golang
func(arg string) {
    fmt.Println("Some String " + arg)
}
//Js: (arg) => {console.log("Some String Value" + arg)}
```
