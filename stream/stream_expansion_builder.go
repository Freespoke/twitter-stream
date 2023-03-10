package stream

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type (
	//IStreamQueryParamsBuilder is the interface for StreamQueryParamBuilder.
	IStreamQueryParamsBuilder interface {
		AddBackFillMinutes(minutes uint) *StreamQueryParamBuilder
		AddExpansion(expansion string) *StreamQueryParamBuilder
		AddMediaField(mediaField string) *StreamQueryParamBuilder
		AddPlaceField(placeField string) *StreamQueryParamBuilder
		AddPollField(pollField string) *StreamQueryParamBuilder
		AddTweetField(tweetField string) *StreamQueryParamBuilder
		AddUserField(userField string) *StreamQueryParamBuilder
		Build() *url.Values
	}

	// StreamQueryParamBuilder is a struct used for requesting additional data from a tweet.
	// Read more at https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
	StreamQueryParamBuilder struct {
		backFillMinutes uint
		expansions []*string
		mediaFields []*string
		placeFields []*string
		pollFields []*string
		tweetFields []*string
		userFields []*string
	}

)

// NewStreamQueryParamsBuilder creeates a struct that implements IStreamQueryParamsBuilder.
// It is used to request additional data from a tweet.
// Read more at https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
func NewStreamQueryParamsBuilder() IStreamQueryParamsBuilder {
	return &StreamQueryParamBuilder{
		backFillMinutes: 0,
		expansions: []*string{},
		mediaFields: []*string{},
		placeFields: []*string{},
		pollFields: []*string{},
		tweetFields: []*string{},
		userFields: []*string{},
	}
}

// Build will build and encode the required query params.
func (s *StreamQueryParamBuilder) Build() *url.Values {
	query := new(url.URL).Query()

	s.addQuery(&query, &s.expansions, "expansions")
	s.addQuery(&query, &s.mediaFields, "media.fields")
	s.addQuery(&query, &s.placeFields, "place.fields")
	s.addQuery(&query, &s.pollFields, "poll.fields")
	s.addQuery(&query, &s.tweetFields, "tweet.fields")
	s.addQuery(&query, &s.userFields, "user.fields")

	if s.backFillMinutes > 0 {
		query.Add("backfill_minutes", strconv.Itoa(int(s.backFillMinutes)))
	}

	return &query
}


// AddExpansion adds an expansion defined in https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
// With expansions, developers can expand objects referenced in the payload. Objects available for expansion are referenced by ID.
// Add a single expansion for each invoke of `AddExpansion`.
func (s *StreamQueryParamBuilder) AddExpansion(expansion string) *StreamQueryParamBuilder {
	s.expansions = append(s.expansions, &expansion)
	return s
}

// AddMediaField adds a media field which enables you to select which specific media fields will deliver in each returned tweet.
// The Tweet will only return media fields if the Tweet contains media and if you've also included `AddExpansion("attachments.media_keys")`.
// Learn more about media fields on twitter docs https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
// Add a single media field for each invoke of `AddMediaField`.
func (s *StreamQueryParamBuilder) AddMediaField(mediaField string) *StreamQueryParamBuilder {
	s.mediaFields = append(s.mediaFields, &mediaField)
	return s
}

// AddPlaceField adds a place field which enables you to select which specific place fields will deliver in each returned tweet.
// The Tweet will only return place fields if the Tweet contains a place and if you've also included `AddExpansion("geo.place_id")`.
// Learn more about place fields on twitter docs https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
// Add a single place field for each invoke of `AddPlaceField`.
func (s *StreamQueryParamBuilder) AddPlaceField(placeField string) *StreamQueryParamBuilder {
	s.placeFields = append(s.placeFields, &placeField)
	return s
}

// AddPollField adds a poll field which enables you to select which specific poll fields will deliver in each returned tweet.
// The Tweet will only return poll fields if the Tweet contains a place and if you've also included `AddExpansion("attachments.poll_ids")`.
// Learn more about poll fields on twitter docs https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
// Add a single poll field for each invoke of `AddPollField`.
func (s *StreamQueryParamBuilder) AddPollField(pollField string) *StreamQueryParamBuilder {
	s.pollFields = append(s.pollFields, &pollField)
	return s
}

// AddTweetField This fields parameter enables you to select which specific Tweet fields will deliver in each returned Tweet object.
// Specify the desired fields in a comma-separated list without spaces between commas and fields.
// You can also include `AddExpansion("referenced_tweets.id")` to return the specified fields for both the original Tweet and any included referenced Tweets.
// The requested Tweet fields will display in both the original Tweet data object, as well as in the referenced Tweet expanded data object that will be located in the includes data object.
// Learn more about tweet fields on twitter docs https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
func (s *StreamQueryParamBuilder) AddTweetField(tweetField string) *StreamQueryParamBuilder {
	s.tweetFields = append(s.tweetFields, &tweetField)
	return s
}

// AddUserField This fields parameter enables you to select which specific user fields will deliver in each returned Tweet.
// Specify the desired fields in a comma-separated list without spaces between commas and fields.
// While the user ID will be located in the original Tweet object, you will find this ID and all additional user fields in the includes data object.
// You must also pass one of the user expansions to return the desired user field.
// `AddExpansion("author_id")`
// `AddExpansion("entities.mentions.username")`
// `AddExpansion("in_reply_to_user_id")`
// `AddExpansion("referenced_tweets.id.author_id")`
func (s *StreamQueryParamBuilder) AddUserField(userField string) *StreamQueryParamBuilder {
	s.userFields = append(s.userFields, &userField)
	return s
}


// AddBackFillMinutes will allow you to recover up to 5 minutes worth of data that might have been missed during a disconnection.
// This feature is currently only available to the academic research product track!
// Learn more about media fields on twitter docs https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream.
func (s *StreamQueryParamBuilder) AddBackFillMinutes(backFillMinutes uint) *StreamQueryParamBuilder {
	s.backFillMinutes = backFillMinutes
	return s
}

func (s StreamQueryParamBuilder) addQuery(qb *url.Values, fields *[]*string, param string) {
	if len(*fields) > 0 {
		var sb strings.Builder
		for i, expansion := range *fields {
			if i == len(*fields) - 1 {
				sb.WriteString(fmt.Sprintf("%v", *expansion))
			} else {
				sb.WriteString(fmt.Sprintf("%v,", *expansion))
			}
		}
		value := sb.String()
		qb.Add(param, value)
	}
}
