# Калі мы запускаем нашу праграму, яно замяняецца на `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Звярніце ўвагу, што Go не прапануе класічнай функцыі Unix `fork`.
# Звычайна гэта не праблема, бо
# запуск goroutines, стварэнне працэсаў і выкананне
# працэсаў ахоплівае большасць выпадкаў выкарыстання `fork`.
