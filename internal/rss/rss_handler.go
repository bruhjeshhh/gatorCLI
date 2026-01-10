package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	_ "log"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not fetch %w", err)
	}
	req.Header.Set("User-Agent", "gator")
	res, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		return nil, fmt.Errorf("the response was ew %w", err2)
	}

	formatted_Response, err := io.ReadAll(res.Body)
	var fetched RSSFeed

	if err3 := xml.Unmarshal(formatted_Response, &fetched); err3 != nil {
		return nil, fmt.Errorf("couldnt unmarshall")
	}

	fetched.Channel.Title = html.UnescapeString(fetched.Channel.Title)
	fetched.Channel.Description = html.UnescapeString(fetched.Channel.Description)
	for i, _ := range fetched.Channel.Item {
		fetched.Channel.Item[i].Title = html.UnescapeString(fetched.Channel.Item[i].Title)
		fetched.Channel.Item[i].Description = html.UnescapeString(fetched.Channel.Item[i].Description)
	}
	return &fetched, nil

}
