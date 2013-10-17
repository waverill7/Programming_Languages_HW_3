package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func lines ( fileName string ) {
    file, err := os.Open( fileName )
    
    if err != nil {
        log.Fatal( err )
    } else {
        s := bufio.NewScanner( file )
        lineCount := 0        
        
        for s.Scan() {
            line := s.Text()
            line = strings.Join( strings.Fields( line ), "" )
            
            if line != "" && strings.Index( line, "#" ) != 0 {
                lineCount += 1
            }
        }
        
        fmt.Printf( "Number Of Lines = %d\n", lineCount )
        file.Close()    
    }
}

func main () {
    fileName := os.Args[ 1 ]
    lines( fileName )
}