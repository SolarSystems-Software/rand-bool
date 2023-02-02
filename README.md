# rand-bool
Generate bits and booleans easily.

## Usage
This example uses the default generator, but you can always make your own generator.

### Generating random booleans

```go
package main

import randbool "github.com/SolarSystems-Software/rand-bool"

func main() {
	if randbool.Default.NextBool() {
		// true
    } else {
		// false
    }
}
```
