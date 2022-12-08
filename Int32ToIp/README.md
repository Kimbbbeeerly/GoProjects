# Int32ToIp

[IPv4 Address](https://en.wikipedia.org/wiki/Internet_Protocol_version_4) has **4 octets**, where each octet is a single byte *(or 8 bits)*. For example lets take IPv4 Address: `192.168.0.1`
- 1st octet `192` has the binary representation: `11000000`
- 2nd octet `168` has the binary representation: `10101000`
- 3rd octet `0` has the binary representation: `00000000`
- 4th octet `1` has the binary representation: `00000001`

So `192.168.0.1` == `11000000.10101000.00000000.00000001`, we can represent it as the unsigned 32-bit number: `3232235521`

This function taking a 32-bit number, converting it to binary and building IPv4.

## Examples
```
3232235521 ==> "192.168.0.1"
2149583361 ==> "128.32.10.1"
64         ==> "0.0.0.64"
32         ==> "0.0.0.32"
0          ==> "0.0.0.0"
```
