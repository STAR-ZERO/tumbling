# tumbling

[![Go Report Card](https://goreportcard.com/badge/github.com/STAR-ZERO/tumbling)](https://goreportcard.com/report/github.com/STAR-ZERO/tumbling)

Command-line Tumblr image downloader in go.

## Usage

```
$ tumbling --apikey API_KEY --blog BLOG_URL --out OUT_DIR
```

* apikey
	* Register application [here](https://www.tumblr.com/oauth/apps), and get apikey(OAuth Consumer Key).
* blog
	* Tumblr blog URL without scheme, e.g. star-zero.tumblr.com
* out
	* Download directory
* limit(optional)
	* The number of results post. 1 to 20 (default 10)
* offset(optional)
	* Post number to start at (default 0)

example:

```
$ tumbling --apikey XXXXXXXXX --blog star-zero.tumblr.com --out ~/Desktop/img
```

## Installation

```
$ go get github.com/STAR-ZERO/tumbling
```

or

```
$ go get -d github.com/STAR-ZERO/tumbling
$ cd $GOPATH/src/github.com/STAR-ZERO/tumbling
$ make deps
$ make install
```

## License

The MIT License