# ZenginCode

The golang implementation of ZenginCode.

ZenginCode is datasets of bank codes and branch codes for japanese.

## Installation

```golang
go get -u github.com/ka2u/zengin-go.git
```

## Usage

```golang

package main

import (
        "fmt"

        zengincode "github.com/ka2u/zengin-code-go"
)

func main() {
        bank, err := zengincode.New()
        if err != nil {
                fmt.Printf("err %v\n", err)
        }
        fmt.Printf("bank %+v", bank["2241"])
}

```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/ka2u/zengin-go

## License

The library is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).