# go swagger sample

## install

### download

```shell
git clone {current-repo}
```

### gen code

```shell
rm -rf gen && mkdir gen && swagger generate server -t gen -A greetings -f swagger/swagger.yml
```

## run

```shell
go run main.go --port 8899
```