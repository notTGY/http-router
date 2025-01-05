# http-router
Turn your router into a webserver



## Compilation

```
docker build -t http-server-mips .

docker run --name http-server-container http-server-mips
docker cp http-server-container:/app/http_server .
docker stop http-server-container
docker rm http-server-container
```

## Serve file on the computer

```
python -m http.server 8080
```


## Download binary on the router

(assuming you connected into router and it has wget)
In my case, my computer ip on local network is 192.168.1.87

```
wget -O http_server http://192.168.1.87:8080/http_server
```

## Execute binary

First give execution permisson to the file
```
chmod +x http_server
```

Then execute
```
./http_server
```

## Outcome

In my case everything works except the last step.
The output is as follows:
```
./http_server: line 1: EL: not found
./http_server: line 2: syntax error: unexpected "("
```
