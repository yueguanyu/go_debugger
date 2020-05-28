# go_debugger

go_debugger

usage:

```go
import "github.com/yueguanyu/go_debugger"

func main(){
    debug := go_debugger.Debug("Prefix1")
	debug("test: ", []int{2, 4, 6, 7})
	debug1 := go_debugger.Debug("Prefix2")
	debug1("test: ", []int{2, 4, 6, 7})
	debug2 := go_debugger.Debug("Prefix3")
	debug2("test: ", []int{2, 4, 6, 7})
	debug3 := go_debugger.Debug("Prefix4")
	debug3("test: ", []int{2, 4, 6, 7})
}
```

output:
[log](https://raw.githubusercontent.com/yueguanyu/go_debugger/master/log.png)
