package controller

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/ProxeusApp/node-proof-of-existence/fakes"
	"github.com/ProxeusApp/node-proof-of-existence/service"
	"github.com/ProxeusApp/node-proof-of-existence/service/servicefakes"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite
	controller *Controller

	fakeTwitterService *servicefakes.FakeTwitterService
	fakeEchoContext    *fakes.FakeContext

	tweet *twitter.Tweet
}

func (me *ControllerTestSuite) SetupTest() {
	me.fakeTwitterService = &servicefakes.FakeTwitterService{}
	me.fakeEchoContext = &fakes.FakeContext{}
	me.tweet = &twitter.Tweet{
		User:     &twitter.User{Name: "Silvio"},
		FullText: "any test",
	}

	me.controller = NewController(me.fakeTwitterService)
}

func (me *ControllerTestSuite) Test_GetTweetByURLOrIDIsFound() {
	fakeRequest := &http.Request{Body: ioutil.NopCloser(strings.NewReader("{ \"tweetUrlOrID\": \"329042394002934\" }"))}
	me.fakeEchoContext.RequestReturns(fakeRequest)
	me.fakeTwitterService.GetTweetByUrlOrIdReturns(me.tweet, nil)

	err := me.controller.GetTweetByURLOrID(me.fakeEchoContext)

	assert.Nil(me.T(), err, "should not return error")
	assert.Equal(me.T(), 1, me.fakeEchoContext.RequestCallCount())
	assert.Equal(me.T(), 1, me.fakeTwitterService.GetTweetByUrlOrIdCallCount())
	assert.Equal(me.T(), "329042394002934", me.fakeTwitterService.GetTweetByUrlOrIdArgsForCall(0))
}

func (me *ControllerTestSuite) Test_GetTweetByURLOrIDEmptyBody() {
	fakeRequest := &http.Request{Body: ioutil.NopCloser(strings.NewReader(""))}
	me.fakeEchoContext.RequestReturns(fakeRequest)
	me.fakeTwitterService.GetTweetByUrlOrIdReturns(me.tweet, nil)
	me.fakeEchoContext.BindReturns(errors.New("any"))

	err := me.controller.GetTweetByURLOrID(me.fakeEchoContext)

	assert.Nil(me.T(), err, "should not return error")
	assert.Equal(me.T(), 1, me.fakeEchoContext.RequestCallCount())
	assert.Equal(me.T(), http.StatusBadRequest, me.fakeEchoContext.NoContentArgsForCall(0), "should return Bad Request")
	assert.Equal(me.T(), 0, me.fakeTwitterService.GetTweetByUrlOrIdCallCount(), "should never call twitterService")
}

func (me *ControllerTestSuite) Test_GetTweetByURLOrIDTweetNotFound() {
	fakeRequest := &http.Request{Body: ioutil.NopCloser(strings.NewReader("{ \"tweetUrlOrID\": \"any\" }"))}
	me.fakeEchoContext.RequestReturns(fakeRequest)
	me.fakeTwitterService.GetTweetByUrlOrIdReturns(nil, service.ErrNotFound)
	me.fakeEchoContext.BindReturns(nil)

	err := me.controller.GetTweetByURLOrID(me.fakeEchoContext)

	assert.Nil(me.T(), err, "should not return error")
	assert.Equal(me.T(), http.StatusNotFound, me.fakeEchoContext.NoContentArgsForCall(0), "should return Not Found")
}

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
