# gopassgen

Policy based password generator based on GO

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/sayanarijit/gopassgen"
)

func main() {
    p := gopassgen.NewPolicy()
    p.MinDigits = 10
    p.MinSpclChars = 5
    password := gopassgen.Generate(p)
}
```

### Available policy options

| Policy | Variable | Type | Default |
| ------ | -------- | ---- | ------- |
| Minimum length | MinLength | int | 6 |
| Maximum length | MaxLength | int | 16 |
| Minimum capital letters | MinCapsAlpha | int | 0 |
| Minimum small letters | MinSmallAlpha | int | 0 |
| Minimum digits | MinDigits | int | 0 |
| Minimum special characters | MinSpclChars | int | 0 |
