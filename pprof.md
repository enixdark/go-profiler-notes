# Go's pprof tool & format

The various profilers built into Go are designed to work with the [pprof visualization tool](https://github.com/google/pprof). The upstream pprof tool is designed to work with C++, Java and Go programs, but it's recommended to access the tool via the `go tool pprof` version that's [bundled](https://github.com/golang/go/tree/master/src/cmd/pprof) with the Go core. It's largely the same except for a few tweaks.

## pprof Format

The pprof tool defines a [protocol buffer](https://developers.google.com/protocol-buffers) output format, which is described in great detail in this [README](https://github.com/google/pprof/blob/master/proto/README.md) as well as the [profile.proto](https://github.com/google/pprof/blob/master/proto/profile.proto) definition file itself. The format is used for all profiling in Go, and the protocol buffer data files themselves are always comressed using gzip.

TODO: Give a better high level introduction

## Decoding pprof Files

### Using `go tool pprof`

The easiest way to decode a pprof file and see its contents is to use  `go tool pprof -raw`. The output is formatted for human readability, so arguabiliy it's not as  `-raw` as the `protoc` output shown later on.

Let's have a look at the [examples/cpu/pprof.samples.cpu.001.pb.gz](./examples/cpu/pprof.samples.cpu.001.pb.gz) CPU profile included in this repository:

```
$ go tool pprof -raw examples/cpu/pprof.samples.cpu.001.pb.gz

PeriodType: cpu nanoseconds
Period: 10000000
Time: 2021-01-08 17:10:32.116825 +0100 CET
Duration: 3.13
Samples:
samples/count cpu/nanoseconds
         19  190000000: 1 2 3
          5   50000000: 4 5 2 3
          1   10000000: 6 7 8 9 10 11 12 13 14
          1   10000000: 15 16 17 11 18 14
          2   20000000: 6 7 8 9 10 11 18 14
          7   70000000: 19 20 21 22 23 24 14
          3   30000000: 25 26 27 28
Locations
     1: 0x1372f7f M=1 main.computeSum /Users/felix.geisendoerfer/go/src/github.com/felixge/go-profiler-notes/examples/cpu/main.go:39 s=0
     2: 0x13730f2 M=1 main.run.func2 /Users/felix.geisendoerfer/go/src/github.com/felixge/go-profiler-notes/examples/cpu/main.go:31 s=0
     3: 0x1372cf8 M=1 golang.org/x/sync/errgroup.(*Group).Go.func1 /Users/felix.geisendoerfer/go/pkg/mod/golang.org/x/sync@v0.0.0-20201207232520-09787c993a3a/errgroup/errgroup.go:57 s=0
     ...
Mappings
1: 0x0/0x0/0x0   [FN]
```

The output above is truncated, [examples/cpu/pprof.samples.cpu.001.pprof.txt](./examples/cpu/pprof.samples.cpu.001.pprof.txt) has the full version.

### Using `protoc`

For those interested in seeing data closer to the raw binary storage, we need the `protoc` protocol buffer compiler. On macOS you can use `brew install protobuf` to install it, for other platform take a look at the [README's install section](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation).

Now let's take a look at the same CPU profile from above:

```
$ gzcat examples/cpu/pprof.samples.cpu.001.pb.gz | \
  protoc --decode perftools.profiles.Profile  ./profile.proto

sample_type {
  type: 1
  unit: 2
}
sample_type {
  type: 3
  unit: 4
}
sample {
  location_id: 1
  location_id: 2
  location_id: 3
  value: 19
  value: 190000000
}
sample {
  location_id: 4
  location_id: 5
  location_id: 2
  location_id: 3
  value: 5
  value: 50000000
}
...
mapping {
  id: 1
  has_functions: true
}
location {
  id: 1
  mapping_id: 1
  address: 20393855
  line {
    function_id: 1
    line: 39
  }
}
location {
  id: 2
  mapping_id: 1
  address: 20394226
  line {
    function_id: 2
    line: 31
  }
}
...
function {
  id: 1
  name: 5
  system_name: 5
  filename: 6
}
function {
  id: 2
  name: 7
  system_name: 7
  filename: 6
}
...
string_table: ""
string_table: "samples"
string_table: "count"
string_table: "cpu"
string_table: "nanoseconds"
string_table: "main.computeSum"
string_table: "/Users/felix.geisendoerfer/go/src/github.com/felixge/go-profiler-notes/examples/cpu/main.go"
...
time_nanos: 1610122232116825000
duration_nanos: 3135113726
period_type {
  type: 3
  unit: 4
}
period: 10000000
```

The output above is truncated also, [pprof.samples.cpu.001.protoc.txt](./examples/cpu/pprof.samples.cpu.001.protoc.txt) has the full version.