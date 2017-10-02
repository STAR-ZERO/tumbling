package tumbling

type Tumblr struct {
	Meta     Meta     `json:"meta"`
	Response Response `json:"response"`
}

type Meta struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type Response struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Photos []Photo `json:"photos"`
}

type Photo struct {
	OriginalSize Size `json:"original_size"`
}

type Size struct {
	Url string `json:"url"`
}
