# Understand how go test -coverpkg works

## go test a specific package without -coverpkg
```
go test -cover ./math

ok      mytestproject/math      0.244s  coverage: 50.0% of statements

```
It makes sense. If we only test package /math, Only add function is tested and substract is not. So it is 50%.

## go test ./... with/out -coverpkg

### Case 1: 
Without -coverpkg to include util package pattern, the function from utils package will not included into the denominator for the test coverage of math package. so it will just be `Add` function tested out of `Add`, `Substract` functions(50%). For utils package, there is no test case at all so it is 0%.
```
go test -cover ./... -coverprofile=cover.out
        mytestproject/utils             coverage: 0.0% of statements
ok      mytestproject/math      0.372s  coverage: 50.0% of statements

go tool cover -func=cover.out
mytestproject/math/add.go:3:            Add             100.0%
mytestproject/math/substract.go:3:      Subtract        0.0%
mytestproject/utils/logger.go:5:        Log             100.0%
total:                                  (statements)    66.7%

```
go tool cover will still calculate the overall test coverage which includes util.log function.
Note that util.log function is used/tested in TestAdd as well

### Case 2
With -coverpkg=./..., it will include functions imported from utils package. Therefore, 2/3 functions are tested in math package. 
- Add function: tested
- Log function: tested
- Substract function: NOT tested

For package utils, the denomintor is also 3 functions but none of them are tested in utils package.

  
```
go test -coverpkg=./...  ./...  -coverprofile=cover.out
ok      mytestproject/math      (cached)        coverage: 66.7% of statements in ./...
        mytestproject/utils             coverage: 0.0% of statements
go tool cover -func=cover.out
mytestproject/math/add.go:3:            Add             100.0%
mytestproject/math/substract.go:3:      Subtract        0.0%
mytestproject/utils/logger.go:5:        Log             100.0%
total:                                  (statements)    66.7%
```

### Case 3
with -coverpkg=./utils, only Log function from utils package is included as denominator in math package so it is 100% for math package.
```
go test -coverpkg=./utils  ./...  -coverprofile=cover.out
        mytestproject/utils             coverage: 0.0% of statements
ok      mytestproject/math      0.407s  coverage: 100.0% of statements in ./utils

go tool cover -func=cover.out
mytestproject/utils/logger.go:5:        Log             100.0%
total:                                  (statements)    100.0%
```

### Case 4
If we comment the Log function in math.TestAdd and keep `-coverpkg=./...`, 1/3 functions are tested in math as now function `Log` is not tested anymore.

```
go test -coverpkg=./...  ./...  -coverprofile=cover.out
        mytestproject/utils             coverage: 0.0% of statements
ok      mytestproject/math      0.396s  coverage: 33.3% of statements in ./...

go tool cover -func=cover.out
mytestproject/math/add.go:3:            Add             100.0%
mytestproject/math/substract.go:3:      Subtract        0.0%
mytestproject/utils/logger.go:5:        Log             0.0%
total:                                  (statements)    33.3%

```
