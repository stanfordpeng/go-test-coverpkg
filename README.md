# Without substract_test.go

## without coverpkg 1/2 functions are tested
```
go test -cover ./math

ok      mytestproject/math      0.244s  coverage: 50.0% of statements

```
## Test with coverpkg
Note that util.log is tested in TestAdd
```
// without coverpkg, for math package, imported package is not included 
go test -cover ./... -coverprofile=cover.out
        mytestproject/utils             coverage: 0.0% of statements
ok      mytestproject/math      0.372s  coverage: 50.0% of statements

go tool cover -func=cover.out
mytestproject/math/add.go:3:            Add             100.0%
mytestproject/math/substract.go:3:      Subtract        0.0%
mytestproject/utils/logger.go:5:        Log             100.0%
total:                                  (statements)    66.7%

// for all packages, 2/3 functions are tested in math
go test -coverpkg=./...  ./...  -coverprofile=cover.out
ok      mytestproject/math      (cached)        coverage: 66.7% of statements in ./...
        mytestproject/utils             coverage: 0.0% of statements
go tool cover -func=cover.out
mytestproject/math/add.go:3:            Add             100.0%
mytestproject/math/substract.go:3:      Subtract        0.0%
mytestproject/utils/logger.go:5:        Log             100.0%
total:                                  (statements)    66.7%

// for utils, log function is tested 100% in math
go test -coverpkg=./utils  ./...  -coverprofile=cover.out
        mytestproject/utils             coverage: 0.0% of statements
ok      mytestproject/math      0.407s  coverage: 100.0% of statements in ./utils

go tool cover -func=cover.out
mytestproject/utils/logger.go:5:        Log             100.0%
total:                                  (statements)    100.0%

// comment log package in math, 1/3 functions are tested in math
go test -coverpkg=./...  ./...  -coverprofile=cover.out
        mytestproject/utils             coverage: 0.0% of statements
ok      mytestproject/math      0.396s  coverage: 33.3% of statements in ./...

go tool cover -func=cover.out
mytestproject/math/add.go:3:            Add             100.0%
mytestproject/math/substract.go:3:      Subtract        0.0%
mytestproject/utils/logger.go:5:        Log             0.0%
total:                                  (statements)    33.3%

```