# ex04

## System

```bash
$ uname -sr
Linux 5.10.60.1-microsoft-standard-WSL2

$ free -h
              total        used        free      shared  buff/cache   available
Mem:           24Gi       620Mi        24Gi       0.0Ki       161Mi        24Gi
Swap:         7.0Gi       400Mi       6.6Gi
```

## Result

```bash
$ go run main.go
number of channels: 1000, 353.3µs
number of channels: 100000, 52.0385ms
number of channels: 1000000, 417.9659ms
signal: killed
```
