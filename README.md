# rand-bool
Generate bits and booleans easily.

## Usage
These examples use a new generator, but the same generator should be used in real usages.

### Generating random booleans
```
generator := RandBool.New()
randBoolean := RandBool.NextBoolean(generator)

if randBoolean {
    // etc...
}
```

### Generating random bits
```
generator := RandBool.New()
randBit := RandBool.NextBit(generator)

if randBit == 1 {
    // etc...
}
```