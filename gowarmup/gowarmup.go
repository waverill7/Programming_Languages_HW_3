package gowarmup

import (
    "math"
    "math/rand"
    "regexp"
    "strings"
    "time"
)

// Problem a:
func change ( cents int ) [] int {
    coinValues := [] int { 25, 10, 5, 1 }
    coins := make( [] int, 4 )
    
    for i, coinValue := range coinValues {
        coins[ i ] , cents = cents / coinValue, cents % coinValue
    }
    
    return coins
}

// Problem b:
func stripVowels ( s string ) string {
    regularExpression := regexp.MustCompile( "[AaEeIiOoUu]" )
    return regularExpression.ReplaceAllString( s, "" ) 
}

// Problem c:
func scramble ( s string ) string {
    permuted := strings.Split( s, "" )
    rand.Seed( time.Now().UTC().UnixNano() )
    
    for i := range permuted {
        randomIndex := int ( math.Floor( rand.Float64() * float64( i ) ) )
        c := permuted[ i ]
        permuted[ i ] = permuted[ randomIndex ]
        permuted[ randomIndex ] = c
    }
    
    return strings.Join( permuted, "" )
}

// Problem d:
func powersOfTwo ( limit int, g func ( int ) ) {
    base := 2
    
    for x := 1; x <= limit; x *= base {
        g( x )
    }
}

// Problem e:
func powers ( base int, limit int, g func ( int ) ) {
    
    if ( base > 1 ) {
    
        for x := 1; x <= limit; x *= base {
            g( x )
        }
    }
}

// Problem f:
func interleave ( L1 [] int, L2 [] int ) [] int {
    interleavedLists := make( [] int, len( L1 ) + len( L2 ) )
    i := 0
    j := 0
    k := 0
    
    for len( L1[ i:len( L1 ) ] ) > 0 || len( L2[ j:len( L2 ) ] ) > 0 {
        
        if len( L1[ i:len( L1 ) ] ) > 0 {
            interleavedLists[ k ] = L1[ i ]
            i += 1
            k += 1
        }
        
        if len( L2[ j:len( L2 ) ] ) > 0 {
            interleavedLists[ k ] = L2[ j ]
            j += 1
            k += 1
        }
    }
    
    return interleavedLists
}

// Problem g: 
func stutter ( sequence [] int ) [] int {
    stutteredSequence := make( [] int, 2 * len( sequence ) )
    i := 0
    
    for _, x := range sequence {
        stutteredSequence[ i ] = x
        i += 1
        stutteredSequence[ i ] = x
        i += 1
    }

    return stutteredSequence    
}