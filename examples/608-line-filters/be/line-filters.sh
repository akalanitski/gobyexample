# Каб паспрабаваць наш радковы фільтр, спачатку стварыце файл з некалькімі
# радкамі ў ніжнім рэгістры.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Затым выкарыстайце радковы фільтр, каб атрымаць радкі ў верхнім рэгістры.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
