# Калі вы запусціце `exit.go` з дапамогай `go run`,
# каманда `go` атрымае і вывядзе вынік exit
$ go run exit.go
exit status 3

# Збіраючы і выконваючы бінарныя-файл, можна ўбачыць
# стан у тэрмінале.
$ go build exit.go
$ ./exit
$ echo $?
3

# Звярніце ўвагу, што сімвал `!` з нашай праграмы ніколі не выводзіцца.
