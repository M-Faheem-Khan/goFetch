# goFetch

A utilitiy to display system information for linux distributions.

[![Go](https://github.com/M-Faheem-Khan/goFetch/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/M-Faheem-Khan/goFetch/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/M-Faheem-Khan/goFetch)](https://goreportcard.com/report/github.com/M-Faheem-Khan/goFetch)

The utility displays the following information. 
- user@hostname
- OS name
- Kernel version
- CPU Model
- Memory (MB)

![image](https://user-images.githubusercontent.com/17150767/220039830-dfa5e0f5-36e8-4c04-bee6-2da2d8e501f2.png)


### Build
```sh
env GOOS=linux GOARCH=amd64 go build -o goFetch main.go
```


<!-- EOF -->
