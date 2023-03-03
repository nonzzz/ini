# Ini

A simple ini parser write with golang.



```in

[Address]

s=127.0.0.1

[Account]

name=kanno

```


``` go

package main

import "fmt"
import "io/ioutil"
import "github/nonzzz/ini/pkg/api"

func main(){
 buffer :=  ioutil.ReadFile("./test.ini")
 code := string(buffer)
 result := api.Transform(code)

}

```