<img align="right" style="display: inline-block; margin: 0 auto; width: 150px" src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/svg/65.svg" text="image from https://github.com/MariaLetta/free-gophers-pack">

# go-aviation

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
opts := aviation.MetarOptions{
	Stations: "LIBP",
	HoursBeforeNow: "1",
}
metar, _, err := client.Metar.Get(opts)
```

For more sample snippet, check [examples](https://github.com/michelangelomo/go-aviation/tree/main/examples) directory.

## Contributing

## License