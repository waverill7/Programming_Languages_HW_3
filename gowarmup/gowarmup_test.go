package gowarmup

import (
    "sort"
    "strings"
    "testing"
)

func anagram( s, t string ) bool {
    a1, a2 := strings.Split( s, "" ), strings.Split( t, "" )
    sort.Strings( a1 )
    sort.Strings( a2 )
    return strings.Join( a1, "" ) == strings.Join( a2, "" )
}

func TestChange ( t *testing.T ) {
    coins := change( 97 )
    
    for i, c := range [] int { 3, 2, 0, 2 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 97 should be [ 3, 2, 0, 2 ]" )
            break
        }
    }
    
    coins = change( 8 )
    
    for i, c := range [] int { 0, 0, 1, 3 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 8 should be [ 0, 0, 1, 3 ]" )
            break
        }
    }
    
    coins = change( 250 )
    
    for i, c := range [] int { 10, 0, 0, 0 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 250 should be [ 10, 0, 0, 0 ]" )
            break
        }
    }
    
    coins = change( 0 )
    
    for i, c := range [] int { 0, 0, 0, 0 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 0 should be [ 0, 0, 0, 0 ]" )
            break
        }
    }
    
    coins = change( 1 )
    
    for i, c := range [] int { 0, 0, 0, 1 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 1 should be [ 0, 0, 0, 1 ]" )
            break
        }
    }
    
    coins = change( 5 )
    
    for i, c := range [] int { 0, 0, 1, 0 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 5 should be [ 0, 0, 1, 0 ]" )
            break
        }
    }
    
    coins = change( 10 )
    
    for i, c := range [] int { 0, 1, 0, 0 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 10 should be [ 0, 1, 0, 0 ]" )
            break
        }
    }
    
    coins = change( 25 )
    
    for i, c := range [] int { 1, 0, 0, 0 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 25 should be [ 1, 0, 0, 0 ]" )
            break
        }
    }
    
    coins = change( 999 )
    
    for i, c := range [] int { 39, 2, 0, 4 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 999 should be [ 39, 2, 0, 4 ]" )
            break
        }
    }

    coins = change( 999999999 )

    for i, c := range [] int { 39999999, 2, 0, 4 } {
    
        if coins[ i ] != c {
            t.Errorf( "change of 999999999 should be [ 39999999, 2, 0, 4 ]" )
            break
        }
    }    
}

func TestStripVowels ( t *testing.T ) {
    
    if stripVowels( "Hello, world" ) != "Hll, wrld" {
        t.Errorf( "stripVowels of 'Hello, world' should be 'Hll, wrld'" )
    }
    
    if stripVowels( "aeiouAEIOU" ) != "" {
        t.Errorf( "stripVowels of 'aeiouAEIOU' should be ''" )
    }
    
    if stripVowels( "ALL YOUR BASE ARE BELONG TO US" ) != "LL YR BS R BLNG T S" {
        t.Errorf( "stripVowels of 'ALL YOUR BASE ARE BELONG TO US' should be ''LL YR BS R BLNG T S" )
    }
    
    if stripVowels( "Mississippi" ) != "Msssspp" {
        t.Errorf( "stripVowels of 'Mississippi' should be 'Msssspp'" )
    }
    
    if stripVowels( "onomatopoeia" ) != "nmtp" {
        t.Errorf( "stripVowels of 'onomatopoeia' should be 'nmtp'" )
    }
    
    if stripVowels( "1!2@3#4$5%6^7&8*90" ) != "1!2@3#4$5%6^7&8*90" {
        t.Errorf( "stripVowels of '1!2@3#4$5%6^7&8*90' should be '1!2@3#4$5%6^7&8*90'" )
    }
}

func TestScramble ( t *testing.T ) {
    data := [] string { "a", "rat", "BSOD", "BDFL", "Go testing" }
    
    for i, s := range data {
    
        if !anagram( s, scramble( s ) ) {
            t.Errorf( "scramble did not work for case #%d", i ) 
        }
    }
}


func TestPowersOfTwo ( t *testing.T ) {
    data := [] int {}
    powersOfTwo( 0, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powersOfTwo of 0 should evaluate to []" )
    }
    
    data = [] int {}
    powersOfTwo( 1, func ( x int ) { data = append( data, x ) } )
    
    for i, x := range [] int { 1 } {
    
        if data[ i ] != x {
            t.Errorf( "powersOfTwo of 1 should evaluate to [ 1 ]" )
            break
        }
    }
    
    data = [] int {}
    powersOfTwo( 2, func ( x int ) { data = append( data, x ) } )
    
    for i, x := range [] int { 1, 2 } {
    
        if data[ i ] != x {
            t.Errorf( "powersOfTwo of 2 should evaluate to [ 1, 2 ]" )
            break
        }
    }
    
    data = [] int {}
    powersOfTwo( -10, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powersOfTwo of -10 should evaluate to []" )
    }
    
    data = [] int {}
    powersOfTwo( 70, func ( x int ) { data = append( data, x ) } )
    
    for i, x := range [] int { 1, 2, 4, 8, 16, 32, 64 } {
    
        if data[ i ] != x {
            t.Errorf( "powersOfTwo of 70 should evaluate to [ 1, 2, 4, 8, 16, 32, 64 ]" )
            break
        }
    }

    data = [] int {}
    powersOfTwo( 2000, func ( x int ) { data = append( data, x ) } )
    
    for i, x := range [] int { 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024 } {
    
        if data[ i ] != x {
            t.Errorf( "powersOfTwo of 2000 should evaluate to [ 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024 ]" )
            break
        }
    }    
}

func TestPowers ( t *testing.T ) {
    data := [] int {}
    powers( 0, 1000, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powers of 0 and 1000 should evaluate to []" )
    }
    
    data = [] int {}
    powers( 1, 1000, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powers of 1 and 1000 should evaluate to []" )
    }
    
    data = [] int {}
    powers( 1000, 0, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powers of 1000 and 0 should evaluate to []" )
    }

    data = [] int {}
    powers( 1000, 1000, func ( x int ) { data = append( data, x ) } )

    for i, x := range [] int { 1, 1000 } {
        
        if data[ i ] != x {
            t.Errorf( "powers of 1000 and 1000 should evaluate to [ 1, 1000 ]" )
            break
        }
    }

    data = [] int {}
    powers( 3, 400, func ( x int ) { data = append( data, x ) } )

    for i, x := range [] int { 1, 3, 9, 27, 81, 243 } {
        
        if data[ i ] != x {
            t.Errorf( "powers of 3 and 400 should evaluate to [ 1, 3, 9, 27, 81, 243 ]" )
            break
        }
    }

    data = [] int {}
    powers( -2, 2000, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powers of -2 and 2000 should evaluate to []" )
    }

    data = [] int {}
    powers( -2, -2000, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powers of -2 and -2000 should evaluate to []" )
    }

    data = [] int {}
    powers( 2, -2000, func ( x int ) { data = append( data, x ) } )
    
    if len( data ) != 0 {
        t.Errorf( "powers of 2 and -2000 should evaluate to []" )
    }    
}

func TestInterleave ( t *testing.T ) {
    a := interleave( [] int {}, [] int {} )
    
    if len( a ) != len( [] int {} ) {
        t.Errorf( "interleave of [] should be []" )
    }
    
    a = interleave( [] int { 0, 2, 4, 6, 8 }, [] int { 1, 3, 5, 7, 9 } )
    
    for i, x := range [] int { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 } {
    
        if a[ i ] != x {
            t.Errorf( "interleave of [ 0, 2, 4, 6, 8 ] and [ 1, 3, 5, 7, 9 ] should be [ 1, 2, 3, 4, 5, 6, 7, 8, 9 ]" )
            break
        }
    }
    
    a = interleave( [] int { 0, 2, 4, 6, 8 }, [] int {} )
    
    for i, x := range [] int { 0, 2, 4, 6, 8 } {
    
        if a[ i ] != x {
            t.Errorf( "interleave of [ 0, 2, 4, 6, 8 ] and [] should be [ 0, 2, 4, 6, 8 ]" )
            break
        }
    }
    
    a = interleave( [] int {}, [] int { 1, 3, 5, 7, 9 } )
    
    for i, x := range [] int { 1, 3, 5, 7, 9 } {
    
        if a[ i ] != x {
            t.Errorf( "interleave of [] and [ 1, 3, 5, 7, 9 ] should be [ 1, 3, 5, 7, 9 ]" )
            break
        }
    }
    
    a = interleave( [] int { 5, 15, 25 }, [] int { 10, 20 } )
    
    for i, x := range [] int { 5, 10, 15, 20, 25 } {
    
        if a[ i ] != x {
            t.Errorf( "interleave of [ 5, 15, 25 ] and [ 10, 20 ] should be [ 5, 10, 15, 20, 25 ]" )
            break
        }
    }
    
    a = interleave( [] int { 1, 2, 3 }, [] int { 4, 5, 6, 7, 8, 9 } )
    
    for i, x := range [] int { 1, 4, 2, 5, 3, 6, 7, 8, 9 } {
    
        if a[ i ] != x {
            t.Errorf( "interleave of [ 1, 2, 3 ] and [ 4, 5, 6, 7, 8, 9 ] should be [ 1, 4, 2, 5, 3, 6, 7, 8, 9 ]" )
            break
        }
    }
}

func TestStutter ( t *testing.T ) {
    a := stutter( [] int {} )
    
    if len( a ) != len( [] int {} ) {
        t.Errorf( "stutter [] should be []" )
    }
    
    a = stutter( [] int { 4, 3 } )
    
    for i, x := range [] int { 4, 4, 3, 3 } {
    
        if a[ i ] != x {
            t.Errorf( "stutter of [ 4, 3 ] should be [ 4, 4, 3, 3 ]" ) 
            break
        }
    }
    
    a = stutter( [] int { 1, 2, 3, 4, 5, 6, 7, 8, 9 } )
    
    for i, x := range [] int { 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9 } {
    
        if a[ i ] != x {
            t.Errorf( "stutter of [ 1, 2, 3, 4, 5, 6, 7, 8, 9 ] should be [ 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9 ]" ) 
            break
        }
    }
}

