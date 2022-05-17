<img align="right" style="display: inline-block; margin: 0 auto; width: 150px" src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/svg/65.svg" text="image from https://github.com/MariaLetta/free-gophers-pack">

# go-aviation

[![Go Reference](https://pkg.go.dev/badge/github.com/michelangelomo/go-aviation.svg)](https://pkg.go.dev/github.com/michelangelomo/go-aviation)
[![Go Report Card](https://goreportcard.com/badge/github.com/michelangelomo/go-aviation)](https://goreportcard.com/report/github.com/michelangelomo/go-aviation)
[![codebeat badge](https://codebeat.co/badges/bbc89b3b-0f6e-4467-911d-f968ad35cb07)](https://codebeat.co/projects/github-com-michelangelomo-go-aviation-main)
![Go](https://github.com/michelangelomo/go-aviation/actions/workflows/go.yml/badge.svg)

**:construction: go-aviation is still under development**

A dead-simple Go library for METAR/TAF/SIGMET

## Install

```
go get github.com/michelangelomo/go-aviation
```

## Usage

```go
import "github.com/michelangelomo/go-aviation"
```

Create a new API Client with no http client options, for example:

```go
client := aviation.NewClient(nil)
```

Some methods have custom parameters, like METAR:

```go
opts := aviation.Options{}
opts.SetStations("LIBP")
opts.SetMostRecent(true)
opts.SetHoursBeforeNow(1)
metar, _, err := client.Metar.Get(opts)
// LIBP 081750Z AUTO VRB02KT 9999 OVC068/// 16/14 Q1019
```

Options can be combined together, for example:

```go
opts = aviation.Options{}
opts.SetStations("LIBP")
opts.SetStartTime(time.Now().Add(-time.Hour * 12)) // last 12 hours
opts.SetEndTime(time.Now())
taf, _, err := client.Taf.Get(opts)
// TAF LIBP 081700Z 0818/0918 VRB05KT 9999 BKN060 TEMPO 0900/0906 3000 BR
// TAF LIBP 081100Z 0812/0912 02010KT 9999 SCT015 BKN060 PROB40 TEMPO 0812/0815 4000 TSRA FEW015CB BKN060 TEMPO 0815/0820 RA BECMG 0816/0818 VRB05KT
// ...
```

You can use `opts.SetMostRecent(true)` option and only one TAF will returns.

For more sample snippets, check [examples](https://github.com/michelangelomo/go-aviation/tree/main/examples) directory.

## Contributing

Pull requests are welcome :hugs:

## License

[MIT](https://github.com/michelangelomo/go-aviation/blob/main/LICENSE)

## Authors

- [@michelangelomo](https://github.com/michelangelomo)