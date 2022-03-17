
# Go ipsum

a small tool to get geo information for an IP address




## Usage

fetch a desired IP address
```
go-ipsum --ip=8.8.8.8
```

fetch your own IP address
```
go-ipsum
```
## Output

```
IP:  8.8.8.8
country:  United States
region:  Virginia
city:  Ashburn
zip:  20149
lat:  39.03
long:  -77.5
isp:  Google LLC

```
## Build

force GOPATH way (optional)
```
go env -w GO111MODULE=off

```

build for your (target) system
```
go build

# build for Linux ARM
GOOS=linux GOARCH=arm64 go build

# build for Windows 64 Bit
GOOS=windows GOARCH=amd64 go build

# build for MacOS ARM
GOOS=darwin GOARCH=arm64 go build

```