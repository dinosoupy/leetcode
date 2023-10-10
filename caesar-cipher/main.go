func caesarCipher(s string, k int32) string {

    cipher:=[]byte(s)
    for i, char:=range cipher {
        if isLetter(char) {
            base:= byte('a')
            if isUppercase(char) {
                base= byte('A')
            } 
            cipher[i] = (char - base + byte(k)) % 26 + base
        }
    }
    return string(cipher)
}

func isLetter(c byte) bool {
    if c >= 'a' && c<= 'z' || c>= 'A' && c <='Z' {
        return true
    } else {
        return false
    }
}

func isUppercase(c byte) bool {
    if c >= 'A' && c <= 'Z' {
        return true
    } else {
        return false
    }


