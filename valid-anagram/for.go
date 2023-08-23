package main

import (
    "fmt"
)

func main() {
    var ans=isAnagram("yoyo", "ooyy")
    fmt.Println(ans)
}

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    smap:=make(map[byte]int)
    tmap:=make(map[byte]int)
    for i:=0; i<len(s); i++ {
        smap[s[i]]++
    }
    for j:=0; j<len(t); j++ {
        tmap[t[j]]++
    }
    for k, v:=range tmap{
        if v != smap[k] {
            return false
        }
    }
    return true
}