# Programmers Playground
A platform for programmers to easy their programming journey..

## Getting started
```bash
./bootstrap
```

## Run
```bash
cd $GOPATH/src/github.com/0xTanvir/pp
go get . && pp run
```

## Before commit please make sure to check this

### check gometalinter to your working package
```
gometalinter \
            --vendor \
            --exclude=../go/src \
            --deadline 60s \
            --enable-all \
            ./...
```

### import
- first go core packages
- second pp packages
- third vendor packages

each separated by a new line
```
import (
	"fmt"
	"strings"

	"github.com/0xTanvir/pp/cfg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
```
