# UIU programmers
A platform for UIU programmers.

## What is it?
UP (UIU programmers) is a platform where UIU programmers can track programming contest path

## Before commit please make sure to check this

### gometalinter
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
- second up packages
- third vendor packages

each separated by a new line
```
import (
	"fmt"
	"strings"

	"github.com/0xTanvir/up/cfg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
```
