## Data Type

### Numeric
|Data Type  |coverage                                              |
|-----------|------------------------------------------------------|
|uint8      |0 to 255                                              |
|uint16     |0 to 65535                                            |
|uint32     |0 to 4294967295                                       |
|uint64     |0 to 18446744073709551615                             |
|uint       |Depending on value (`uint32` or `uint64`)             |
|byte       |alias for `uint8`                                     |
|int8       |-128 to 127                                           |
|int16      |-32768 to 32767                                       |
|int32      |-2147483648 to 2147483647                             |
|int64      |-9223372036854775808 to 9223372036854775807           |
|int        |Depending on value (`int32` or `int64`)               |
|rune       |alias for `uint32`                                    |
|float64    |IEEE-754 64-bit floating-point numbers                |
|float32    |IEEE-754 32-bit floating-point numbers                |
|complex64  |complex numbers with float32 real and imaginary parts |
|complex128 |complex numbers with float64 real and imaginary parts |

### Zero value
- Zero value of string `""`
- Zero value of boolean `false`
- Zero value of Non Decimal Numeric `0`
- Zero value of Decimal Numeric `0.0`

