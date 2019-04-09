# ZenginCode

The golang implementation of ZenginCode.

ZenginCode is datasets of bank codes and branch codes for japanese.

## Installation

```golang
go get -u github.com/ka2u/zengin-go.git
```

## Usage

### Set environment variable

- ZENGIN_SOURCE_ROOT
    - Set the absolute path where is the zengincode source data directory.
    - You use zengin source from file(not embed), this is MANDATORY.
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

	b, err := bank.Find("0005")
	if err != nil {
		fmt.Printf("err %v\n", err)
	} else {
		fmt.Printf("bank %s", b.Name)
	}

	br, err := b.Branches.Find("恵比寿")
	if err != nil {
		fmt.Printf("err %v\n", err)
	} else {
		fmt.Printf("branch %s", br.Name)
	}

}

```

## Embed

Use Statik if You want to use embed zengin source.

### make embed file

```
statik -source=source-data/data
```

```golang

package main

import (
        "fmt"

        zengincode "github.com/ka2u/zengin-code-go"
        - "your/app/path/statik"
)

func main() {
	bank, err := zengincode.NewWithEmbed()
	if err != nil {
		fmt.Printf("err %v\n", err)
	}

	b, err := bank.Find("0005")
	if err != nil {
		fmt.Printf("err %v\n", err)
	} else {
		fmt.Printf("bank %s", b.Name)
	}

	br, err := b.Branches.Find("恵比寿")
	if err != nil {
		fmt.Printf("err %v\n", err)
	} else {
		fmt.Printf("branch %s", br.Name)
	}

}

```

## Note

New() method is very heavy loading process.
It is recommended to load and cache at the initial process.

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/ka2u/zengin-go

## License

The library is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).