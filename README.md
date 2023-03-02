# go-idna
IDNA conversion in Go

```
$ go run idn.go 🕵💻.st
name: 🕵💻.st
Simple Punycode conversion: xn--3s8htl.st
MapForLookup, Transitional: xn--3s8htl.st
ValidateForRegistration: xn--3s8htl.st

$ go run idn.go blåbærgrød.no
name: blåbærgrød.no
Simple Punycode conversion: xn--blbrgrd-fxak7p.no
MapForLookup, Transitional: xn--blbrgrd-fxak7p.no
ValidateForRegistration: xn--blbrgrd-fxak7p.no

$ go run idn.go *.faß.com
name: *.faß.com
Simple Punycode conversion: *.xn--fa-hia.com
2023/03/02 13:10:02 idna: disallowed rune U+002A
exit status 1
```