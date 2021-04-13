# Install
1. Download the ISIS environment file
   ```console
   $ wget https://raw.githubusercontent.com/evindunn/go-isis/master/environment.yml
   ```
2. Install ISIS
    ```console
    $ conda env create -p .isis -f environment.yml
    ```
3. Activate the ISIS environment
   ```console
   $ conda activate ./.isis
   ```
4. Download any needed data (in this example, to `/data/disk/isisdata`)
5. Set `ISISROOT` and `ISISDATA`
    ```console
    $ export ISISROOT=.isis
    $ export ISISDATA=/data/disk/isisdata
    ```

# Usage
This example uses [GoDotEnv](https://github.com/joho/godotenv)
in order to use the following file in place of step 5 above:

.env:
```dotenv
ISISROOT=.isis
ISISDATA=/data/disk/isisdata
```

main.go:
```go
package main

import (
   "fmt"
   "github.com/evindunn/go-isis"
   "github.com/joho/godotenv"
   "os"
   "strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Run a single command
	err = isis.Isis(
		"mroctx2isis",
		map[string]string{
			"from": "testfiles/P22_009808_1436_XI_36S210W.IMG",
			"to": "testfiles/P22_009808_1436_XI_36S210W.single.cub",
		},
	)

	if err != nil {
		panic(err)
	}

	// Run using a Pool
	inputFiles := []string{
		"testfiles/P22_009808_1436_XI_36S210W.IMG",
		"testfiles/J21_052811_1983_XN_18N282W.IMG",
	}
	pool := isis.NewPool()

	for _, file := range inputFiles {
		outFile := strings.Replace(file, ".IMG", ".cub", 1)
		pool.Run(
			"mroctx2isis",
			map[string]string{
				"from": file,
				"to":   outFile,
			},
		)
	}

	errs := pool.Wait()
	hadError := false
	for _, err := range errs {
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			hadError = true
		}
	}

	if hadError {
		os.Exit(1)
	}
}
```
