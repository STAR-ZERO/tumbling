package tumbling

import (
	"errors"
	"os"

	"github.com/urfave/cli"
)

type Options struct {
	apikey string
	blog   string
	out    string
	limit  int
	offset int
}

func applyOption(c *cli.Context) (*Options, error) {
	apikey := c.String("apikey")
	blog := c.String("blog")
	out := c.String("out")
	limit := c.Int("limit")
	offset := c.Int("offset")

	options := Options{
		apikey,
		blog,
		out,
		limit,
		offset,
	}

	if apikey == "" {
		return nil, errors.New("required --apikey")
	}
	if blog == "" {
		return nil, errors.New("required --blog")
	}
	if out == "" {
		return nil, errors.New("required --out")
	}
	if _, err := os.Stat(out); os.IsNotExist(err) {
		return nil, err
	}

	return &options, nil
}
