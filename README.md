# Programmers Playground
A platform for programmers to easy their programming experience..

## What is it?
Programmers Playground is a platform where programmers can track their programming path

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
