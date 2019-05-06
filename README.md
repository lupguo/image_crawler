# image_crawler
From a URL, analyze the HTML, download all the images in the HTML document you are interested in to the specified directory

## Build

```bash
// download
go get -u github.com/tkstorm/image_crawler

// build
go build github.com/tkstorm/image_crawler
```

## Usage

```
$ ./cmd/image_crawler -h
Usage of ./cmd/image_crawler:
  -analyzed string
    	url page analyzed method (node|regex) (default "regex")
  -d string
    	download image storage dirname (default "/tmp")
  -url string
    	page url request by crawler (default "http://tmall.com/")
```

## Example
```
[root@gearbest-web01-test_10 tmp]# ./image_crawler
total 28 images need to be download...
ok https://blog.golang.org/survey2018/fig16.svg => /tmp/fig16.svg
ok https://blog.golang.org/survey2018/fig18.svg => /tmp/fig18.svg
ok https://blog.golang.org/survey2018/fig17.svg => /tmp/fig17.svg
ok https://blog.golang.org/survey2018/fig14.svg => /tmp/fig14.svg
ok https://blog.golang.org/survey2018/fig13.svg => /tmp/fig13.svg
ok https://blog.golang.org/survey2018/fig12.svg => /tmp/fig12.svg
ok https://blog.golang.org/survey2018/fig5.svg => /tmp/fig5.svg
ok https://blog.golang.org/survey2018/fig29.svg => /tmp/fig29.svg
ok https://blog.golang.org/survey2018/fig28.svg => /tmp/fig28.svg
ok https://blog.golang.org/survey2018/fig11.svg => /tmp/fig11.svg
ok https://blog.golang.org/survey2018/fig15.svg => /tmp/fig15.svg
ok https://blog.golang.org/survey2018/fig22.svg => /tmp/fig22.svg
ok https://blog.golang.org/survey2018/fig19.svg => /tmp/fig19.svg
ok https://blog.golang.org/survey2018/fig27.svg => /tmp/fig27.svg
ok https://blog.golang.org/survey2018/fig10.svg => /tmp/fig10.svg
ok https://blog.golang.org/survey2018/fig26.svg => /tmp/fig26.svg
...
```

## Bugs

1. In the case of too much concurrency, the server will have connection rejection

```
...
Get http://img.alicdn.com/tps/TB1Yn4KMXXXXXaJXVXXXXXXXXXX-400-398.jpg: dial tcp 184.31.46.100:80: i/o timeout
 Get http://img.alicdn.com/tps/TB1kk0CMXXXXXXtapXXXXXXXXXX-400-398.jpg: dial tcp 184.31.46.100:80: i/o timeout
 Get http://img.alicdn.com/tps/TB1EfiHLVXXXXbOXXXXXXXXXXXX-140-140.jpg: dial tcp 184.31.46.100:80: i/o timeout
 Get http://img.alicdn.com/tps/TB1KTyFLVXXXXcRXXXXXXXXXXXX-140-140.jpg: dial tcp 184.31.46.100:80: i/o timeout
 Get http://img.alicdn.com/tps/TB1S4g1JFXXXXXJXpXXXXXXXXXX-394-418.jpg: dial tcp 184.31.46.100:80: i/o timeout
 Get http://img.alicdn.com/tps/TB17HxNMXXXXXXmXVXXXXXXXXXX-400-398.jpg: dial tcp 184.31.46.100:80: i/o timeout
 ...
```

2. Currently, HTTP requests based on single IP are easy to be limited by the server