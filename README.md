# Go migration

## For migrate schema
#### Up
```
go run main.go
# or
go run main.go -a up
```
#### Down
```
go run main.go -a down
```


## For migrate data
#### Up
```
go run main.go -m data
```
#### Down
```
go run main.go -a down -m data
```