# envloader 📂

A simple and lightweight Go package to read `.env` files without modifying system environment variables.

### 🚀 Installation
```sh
go get github.com/paulsonlegacy/go-env-loader
```

### 🛠 Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/paulsonlegacy/go-env-loader"
)

func main() {
	value, err := envloader.GetEnv(".env", "MY_SECRET_KEY", "OPTIONAL_DEFAULT_VALUE")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	fmt.Println("MY_SECRET_KEY:", value)
}
```

*NB - .env files store values as strings, so any fetched value is naturally a string. Even if a value looks like a number (PORT=8080), it's still a string ("8080").*