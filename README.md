# Share View

Quick app that queries Bloomberg's quote page and displays share details on screen. Will refresh at a specified interval (default 1 minute).

### Compile
```
go build main.go
```

### flags
* -c string
       bloomberg instrument codes (default "NFSWIX:SJ,DBXWD:SJ")
* -i duration
        (default 1m0s)

### Usage
```
./main -i=15m0s -c=NFSWIX:SJ,DBXWD:SJ,STPROP:SJ
```

The output will look similar to this:

```
ID           OPEN    NOW     CHANGE %     DATE
NFSWIX:SJ    1596    1600    -0.373599    2016-03-22 17:00:17 +0200 SAST
DBXWD:SJ     2517    2518    0.079491     2016-03-22 17:00:20 +0200 SAST
STPROP:SJ    6659    6680    1.396480     2016-03-22 17:00:06 +0200 SAST
```
