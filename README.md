# rand-bool
Generate bits and booleans easily.

## Usage
These examples use a new generator, but the same generator should be used in real usages.

### Generating random booleans
```
generator := randbool.New()
randBoolean := randbool.NextBool(generator)

if randBoolean {
    // etc...
}
```

### Generating random bits
```
generator := randbool.New()
randBit := randbool.NextBit(generator)

if randBit == 1 {
    // etc...
}
```