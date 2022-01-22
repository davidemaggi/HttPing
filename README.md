# HttPing
Just like ping, but for Http Rquests.
Simply provide an Url an it will send requests, as a default HttPing will send an HEAD request.

###### Please, keep in mind that the local network latency is included in the result 

## Why do i need it?
Well, I've no idea, i created this small util because i needed it and just wanted to share.

## How does it work
HttPing is a Command Line Util, just download id and you are ready to go
###### If you are using it on an unix based system ensure that it can be executed

| Parameter | Mandatory | Default | Description                                                                         |
|:---------:|:---------:|:-------:|-------------------------------------------------------------------------------------|
|     u     |     x     |         | The Url to ping, it can be passed with or without the flag                          |
|     n     |           |    4    | How many requests needs to be send                                                  |
|     t     |           |         | Contnuos Ping, it will override the n flag, to terminate the script use CTRL+C      |
|     g     |           |         | Method used in the request, if set a GET request is sent instead of an HEAD request |
|     v     |           |         | Display HttPing version's details                                                   |
|   help    |           |         | Show Help                                                                           |

#### Samples
Basic:
```console
davide@davide-pc:~/HttPing$ ./httping https://google.it
PING: google.it HEAD (): 
connected to 142.250.184.67 (-1 bytes), seq=0 time=719 ms : 200 OK 
connected to 142.250.184.67 (-1 bytes), seq=1 time=614 ms : 200 OK 
connected to 142.250.184.67 (-1 bytes), seq=2 time=620 ms : 200 OK 
connected to 142.250.184.67 (-1 bytes), seq=3 time=608 ms : 200 OK 
--- https://google.it ping statistics ---
4 connects, 4 ok, 0.00% failed, time 2561 ms
round-trip min/avg/max = 608.0/640.2/719.0 ms
```
Ping Number:
```console
davide@davide-pc:~/HttPing$ ./httping --n 10 https://google.it
PING: google.it HEAD ():
connected to 142.250.184.67 (-1 bytes), seq=0 time=875 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=1 time=614 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=2 time=712 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=3 time=619 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=4 time=814 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=5 time=614 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=6 time=615 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=7 time=620 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=8 time=609 ms : 200 OK
connected to 142.250.184.67 (-1 bytes), seq=9 time=615 ms : 200 OK
--- https://google.it ping statistics ---
10 connects, 10 ok, 0.00% failed, time 6707 ms
round-trip min/avg/max = 609.0/670.7/875.0 ms
```

```console
davide@davide-pc:~/HttPing$ ./httping  https://google.it/
PING: google.it HEAD (/): 
connected to 142.250.184.67 (-1 bytes), seq=0 time=298 ms : 200 OK 
connected to 142.250.184.67 (-1 bytes), seq=1 time=77 ms : 200 OK 
connected to 142.250.184.67 (-1 bytes), seq=2 time=183 ms : 200 OK 
connected to 142.250.184.67 (-1 bytes), seq=3 time=72 ms : 200 OK 
--- https://google.it/ ping statistics ---
4 connects, 4 ok, 0.00% failed, time 630 ms
round-trip min/avg/max = 72.0/157.5/298.0 ms
```

## Can I install it?
Yes, of course

#### Windows
//TODO
#### Linux
//TODO
#### macOS
//TODO