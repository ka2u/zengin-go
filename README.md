# ZenginCode

The golang implementation of ZenginCode.

ZenginCode is datasets of bank codes and branch codes for japanese.

## Installation

```golang
go get -u github.com/ka2u/zengin-go.git
```

## Usage

### Set environment variable

- ZENGIN_SOURCE_ROOT or ZENGIN_SOURCE_INCLUDE (MANDATORY)
    - Set the absolute path where is the zengincode source data directory.
    - Set the value is TRUE or FALSE If you want to use embedded zengincode data.
- ZENGIN_SOURCE_YAML (OPTIONAL)
    - Set the value is TRUE or FALSE If you want to use YAML data. Default data is JSON.

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
        bank, err := bank.Find("2241")
        if err != nil {
            fmt.Error("not found")
        }
        fmt.Printf("bank %+v", bank)
}

```

## Embedded

I use [fileb0x](https://github.com/UnnoTed/fileb0x).

## Note

New() method is very heavy loading process.
It is recommended to load and cache at the initial process.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/ka2u/zengin-go

## License

The library is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).