// string and byte
// string is made of byte
// rune is also accepted, but byte is preferred

// stinrg 声明
s := "hello"
var s = "Hello"

// byte 声明
b := s[0]
var b byte = s[0]

// string to []byte
b_slice := []byte(s)
var bs []byte = []byte(s)

// []byte to string
s2 := string(bs)