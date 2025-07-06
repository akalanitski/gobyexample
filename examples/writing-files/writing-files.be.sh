# Паспрабуйце запусціць код запісу ў файл.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Затым праверце змесціва запісаных файлаў.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Далей мы разгледзім прымяненне некаторых ідэй файлнага ўводу/вываду
# якія мы толькі што разглядалі да патокаў `stdin` і `stdout`
