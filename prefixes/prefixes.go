package main

import (
    "fmt"
    "os"
    "strings"
)

func prefixes ( s string ) {
    p := strings.Split( s, "" )
    
    for i := 0; i <= len( p ); i++ {
        fmt.Println( strings.Join( p[ :i ], "" ) )
    }    
}

func main() {
    s := os.Args[ 1 ]
    prefixes( s )
}