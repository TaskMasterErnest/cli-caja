
# Word Counter

This application is a CLI program that to used to count the number of words in a stream if text passed into it.
## Installation

Install wordcounter as `wc` with Go

```bash
  cd wordcounter
  go build -o wc main.go
```
## Running Tests

To run tests, run the following command

```bash
  go test -v
```


## Usage/Examples

```bash
  echo "This is wordcounter" | ./wc

  # you can count lines and bytes too

  ./wc -l # to count lines
  ./wc -b # to count bytes
```

![wordcounter](/assets/wordcounter.gif)



