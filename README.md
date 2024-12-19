# 1/2 functions are tested
```
go test -cover ./math

ok      mytestproject/math      0.244s  coverage: 50.0% of statements

```
# 2/3 function are tested TestAdd
Note that util.log is tested in 
```
go test -cover ./...
        mytestproject/utils             coverage: 0.0% of statements
ok      mytestproject/math      0.334s  coverage: 66.7% of statements
```