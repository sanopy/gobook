# ex 7.7

`flag.CommandLine.Var` によるフラグを追加した際に呼び出される関数は以下に示す部分で定義されている。
<https://github.com/golang/go/blob/c2397905e027cdbab3a28d02813adcb82368422c/src/flag/flag.go#L861-L885>

上記のコードから、追加される flag 情報は、以下のコードにより生成されていることがわかる。

```go
func (f *FlagSet) Var(value Value, name string, usage string) {

	// 割愛

	// Remember the default value as a string; it won't change.
	flag := &Flag{name, usage, value, value.String()}

	// 割愛
}
```

また、Flag型は以下のように定義されている。ヘルプメッセージにおいて表示されるデフォルト値は、 `DefValue` プロパティで管理されている。

```go
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
```

以上のことから、ヘルプメッセージに `°` が含まれていた理由は、flag情報を管理する際に `.String()` が呼ばれていたためである。
