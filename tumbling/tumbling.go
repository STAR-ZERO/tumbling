package tumbling

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"
)

func Run(c *cli.Context) error {
	options, err := applyOption(c)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	url := fmt.Sprintf("https://api.tumblr.com/v2/blog/%s/posts?&api_key=%s&type=photo&limit=%d&offset=%d", options.blog, options.apikey, options.limit, options.offset)

	body, err := fetch(url)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	urls, err := getImageUrls(body)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	eg := errgroup.Group{}
	for _, url := range urls {
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		eg.Go(func() error {
			return download(url, options.out)
		})
	}

	if err := eg.Wait(); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func getImageUrls(body []byte) ([]string, error) {
	var tumblr Tumblr
	err := json.Unmarshal(body, &tumblr)
	if err != nil {
		return nil, err
	}

	urls := make([]string, 0)
	for _, post := range tumblr.Response.Posts {
		for _, photo := range post.Photos {
			urls = append(urls, photo.OriginalSize.Url)
		}
	}

	return urls, nil
}

func download(url string, out string) error {
	filename := path.Base(url)
	file, err := os.Create(filepath.Join(out, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		os.Remove(file.Name())
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		os.Remove(file.Name())
		return err
	}

	return nil
}
