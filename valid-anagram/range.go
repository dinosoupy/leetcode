func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    smap:=make(map[rune]int)
    tmap:=make(map[rune]int)
    for _, v:=range s {
        smap[v]++
    }
    for _, v:=range t {
        tmap[v]++
    }
    for k, _:=range tmap{
        if tmap[k] != smap[k] {
            return false
        }
    }
    return true
