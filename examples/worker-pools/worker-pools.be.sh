# Наша праграма паказвае 5 задач, якія выконваюцца
# рознымі работнікамі. Праграма выконвае задачу ўсяго каля 2 секунд,
# нягледзячы на тое, што ў агульнай складанасці яна выконвае каля 5 секунд працы, таму што
# адначасова працуюць 3 работнікі.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
