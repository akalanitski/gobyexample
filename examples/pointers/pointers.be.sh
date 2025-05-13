# `zeroval` не змяняе `i` у `main`, але
# `zeroptr` змяняе, бо мае спасылку на
# адрас памяці для гэтай зменнай.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
