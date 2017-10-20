package main

import (
	"os"

	"github.com/STAR-ZERO/tumbling/tumbling"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "tumbling"
	app.Usage = "Download Tumblr images"
	app.Version = "0.1.0"
	app.UsageText = "tumbling --apikey API_KEY --blog BLOG_URL --out OUT_DIR"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "apikey",
			Usage: "API Key",
		},
		cli.StringFlag{
			Name:  "blog",
			Usage: "Tumblr blog URL",
		},
		cli.StringFlag{
			Name:  "out",
			Usage: "Download directory",
		},
		cli.IntFlag{
			Name:  "limit",
			Value: 10,
			Usage: "Number of posts. 1 to 20",
		},
		cli.IntFlag{
			Name:  "offset",
			Value: 0,
			Usage: "Post number to start at",
		},
	}

	app.Action = tumbling.Run
	app.Run(os.Args)
}
