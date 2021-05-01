Go Text Splitter
===

> Split a string to chunks of given size at whitespaces.

## Installation

``` bash
$ go get github.com/AidenHadisi/go-text-splitter/splitter
```

## Usage
``` go
import (
	"github.com/AidenHadisi/go-text-splitter/splitter"
)
```

#### func split(str, chunkSize) []string;

Splits the given string into an array of strings with maximum length of `chunkSize`

``` go
chunks := splitter.Split("hello world! What's up?", 7);
/* ["hello", "world!", "What's", "up?"] */
```

If no space is found before `chunkSize`, it will break up the word to ensure the chunks have mixmum length of `chunkSize`.
``` go
chunks := splitter.Split("foo bar test", 3);
/* ["foo", "bar", "tes", "t"] */
```

## Contributing 
Contributions are welcome. Feel free to improve or add new features.
