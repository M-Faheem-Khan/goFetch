# goFetch

A utilitiy to display system information for linux distributions.


The utility displays the following information. 
- user@hostname
- OS name
- Kernel version
- CPU Model
- Memory (MB)

```sh
env GOOS=linux GOARCH=amd64 go build -o goFetch main.go
```


<!-- EOF -->