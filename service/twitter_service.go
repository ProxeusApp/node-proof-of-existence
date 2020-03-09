package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2/clientcredentials"
)

var ErrNotFound = errors.New("not found")

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TwitterService
type TwitterService interface {
	GetTweetByUrlOrId(urlOrId string) (*twitter.Tweet, error)
}

type httpTwitterService struct {
	client *twitter.Client
}

const tweetUrlRegexp = `https?:\/\/twitter\.com\/(?:\#!\/)?(\w+)\/status(es)?\/(?P<ID>\d+)`

func NewDefaultTwitterService(ctx context.Context, consumerKey, consumerSecret string) *httpTwitterService {
	config := &clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	httpClient := config.Client(ctx)

	return &httpTwitterService{
		client: twitter.NewClient(httpClient),
	}
}

func (me *httpTwitterService) GetTweetByUrlOrId(urlOrId string) (*twitter.Tweet, error) {

	id, err := me.parseID(urlOrId)
	if err != nil {
		return nil, err
	}

	tweet, _, err := me.client.Statuses.Show(id, nil)
	return tweet, err
}

// Given an url of a status, will return its id, parsed.
// Also just an id can be given, that will simply be returned in int format
func (me *httpTwitterService) parseID(url string) (int64, error) {
	log.Printf("Parsing %s", url)

	// Try to convert to number first
	id, err := strconv.Atoi(url)
	if err == nil {
		// User gave an id, not an url
		return int64(id), nil
	}

	rxp := regexp.MustCompile(tweetUrlRegexp)

	if !rxp.MatchString(url) {
		return 0, errors.New("no match")
	}

	vars := rxp.FindStringSubmatch(url)
	log.Println(vars)
	if len(vars) != 4 {
		return 0, fmt.Errorf("returned different amount of ids than expected: %v (total: %d)", vars, len(vars))
	}

	return strconv.ParseInt(vars[3], 10, 64)
}
