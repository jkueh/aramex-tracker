# Aramex Tracker

I wrote this script to pull package information from the Aramex (formerly Fastway) parcel tracking API, instead of
spending many hours (and bytes!) refreshing the webpage.

## Build

```
go build -o aramex-tracker
```

## Usage

Note: 'Label Number' is synonymous with 'Tracking Number' in some cases.

### Running

If you're lazy and don't want to run the `go build` command above:
```
go run . <label_number_1> [<label_number_N>]
```

Otherwise:

```
./aramex-tracker <label_number_1> [<label_number_N>]
```

### Extending / Rebuilding

All known object models are inside the `aramex/` directory, and can be imported like so:

```go
import (
	// Your other packages here
	"github.com/jkueh/aramex-tracker/aramex"
)
```

### Debugging

There's a `DEBUG` package variable that can be set by setting an equivalent environment variable that will spit out more
information about what's going on, e.g.

```
DEBUG=true ./aramex-tracker <label_number_1> [<label_number_N>]
```

## Caveats

* The structs were built using a limited sample set. It's likely that the structure will change over time.
* Some of these fields don't really make much sense on their own - e.g. the `Name` field in a scan event appears to have
  the name of locations.
