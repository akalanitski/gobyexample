# Каб паэксперыментаваць з параметрамі нашай каманды, лепш за ўсё
# спачатку скампіляваць яе, а потым запусціць атрыманы
# бінарны файл
$ go build command-line-flags.go

# Паспрабуйце сабраную каманду, спачатку задаўшы ёй значэнні
# усіх параметраў
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Звярніце ўвагу, што калі вы прапусціце параметр, яны аўтаматычна атрымаюць
# стандартнае значэнне.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# Не апісаныя парметры можна праверрыць наступным чынам
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Звярніце ўвагу, што пакет `flag` патрабуе, каб усе сцягі
# стаялі перад пазіцыйнымі аргументамі (інакш сцягі
# будуць інтэрпрэтавацца як пазіцыйныя аргументы).
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# Выкарыстоўвайце сцягі `-h` або `--help`, каб атрымаць аўтаматычна
# згенераваны тэкст даведкі для праграмы каманднага радка.
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Калі вы падасце сцяг, які не быў указаны ў пакете `flag`, праграма вывядзе паведамленне пра памылку
# і зноў пакажа тэкст даведкі.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
