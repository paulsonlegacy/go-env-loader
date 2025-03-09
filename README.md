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
	value, err := envloader.GetEnv(".env", "MY_SECRET_KEY")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	fmt.Println("MY_SECRET_KEY:", value)
}
```

