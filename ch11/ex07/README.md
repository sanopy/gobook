# ex07

## 最適化前

最適化前のコードは、 `ch11/ex02/intset.go` の内容で計測。

```bash
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/sanopy/gobook/ch11/ex07
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkIntSetHas_1e3-6                130277146                9.793 ns/op           0 B/op          0 allocs/op
BenchmarkIntSetHas_1e5-6                100000000               10.15 ns/op            0 B/op          0 allocs/op
BenchmarkIntSetHas_1e7-6                100000000               10.07 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetHas_1e3-6             44717218                25.24 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetHas_1e5-6             49470459                23.32 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetHas_1e7-6             45078210                23.56 ns/op            0 B/op          0 allocs/op
BenchmarkIntSetAdd_1e3-6                  784621              1578 ns/op             248 B/op          5 allocs/op
BenchmarkIntSetAdd_1e5-6                   84861             14450 ns/op           40184 B/op         13 allocs/op
BenchmarkIntSetAdd_1e7-6                     638           1743586 ns/op         7306306 B/op         32 allocs/op
BenchmarkMapIntSetAdd_1e3-6               161005              7307 ns/op            3434 B/op         18 allocs/op
BenchmarkMapIntSetAdd_1e5-6               163104              7343 ns/op            3493 B/op         19 allocs/op
BenchmarkMapIntSetAdd_1e7-6               163945              7413 ns/op            3493 B/op         19 allocs/op
BenchmarkIntSetUnionWith_1e3-6          71895032                16.05 ns/op            0 B/op          0 allocs/op
BenchmarkIntSetUnionWith_1e5-6            754123              1330 ns/op               0 B/op          0 allocs/op
BenchmarkIntSetUnionWith_1e7-6              7909            152058 ns/op            1675 B/op          0 allocs/op
BenchmarkMapIntSetUnionWith_1e3-6         449094              2773 ns/op               0 B/op          0 allocs/op
BenchmarkMapIntSetUnionWith_1e5-6         401194              2862 ns/op               0 B/op          0 allocs/op
BenchmarkMapIntSetUnionWith_1e7-6         410791              2780 ns/op               0 B/op          0 allocs/op
PASS
ok   github.com/sanopy/gobook/ch11/ex07 23.478s
```

## 最適化後

```bash
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/sanopy/gobook/ch11/ex07
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkIntSetHas_1e3-6                129865420                9.378 ns/op           0 B/op          0 allocs/op
BenchmarkIntSetHas_1e5-6                125913118                9.713 ns/op           0 B/op          0 allocs/op
BenchmarkIntSetHas_1e7-6                129062330                9.320 ns/op           0 B/op          0 allocs/op
BenchmarkMapIntSetHas_1e3-6             55935960                22.63 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetHas_1e5-6             46125283                21.68 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetHas_1e7-6             48387096                21.66 ns/op            0 B/op          0 allocs/op
BenchmarkIntSetAdd_1e3-6                  899071              1273 ns/op             258 B/op          2 allocs/op
BenchmarkIntSetAdd_1e5-6                  165891              6923 ns/op           28600 B/op          2 allocs/op
BenchmarkIntSetAdd_1e7-6                    2743            465572 ns/op         3046335 B/op          3 allocs/op
BenchmarkMapIntSetAdd_1e3-6               170313              6926 ns/op            3434 B/op         18 allocs/op
BenchmarkMapIntSetAdd_1e5-6               171108              7006 ns/op            3492 B/op         19 allocs/op
BenchmarkMapIntSetAdd_1e7-6               167239              7250 ns/op            3492 B/op         19 allocs/op
BenchmarkIntSetUnionWith_1e3-6          93104089                12.76 ns/op            0 B/op          0 allocs/op
BenchmarkIntSetUnionWith_1e5-6           1000000              1052 ns/op               0 B/op          0 allocs/op
BenchmarkIntSetUnionWith_1e7-6              9631            111708 ns/op             669 B/op          0 allocs/op
BenchmarkMapIntSetUnionWith_1e3-6         467440              2496 ns/op               0 B/op          0 allocs/op
BenchmarkMapIntSetUnionWith_1e5-6         445722              2641 ns/op               0 B/op          0 allocs/op
BenchmarkMapIntSetUnionWith_1e7-6         441252              2643 ns/op               0 B/op          0 allocs/op
PASS
ok   github.com/sanopy/gobook/ch11/ex07 25.367s
```
