package rollbar

import (
	"fmt"
	"github.com/davidji99/simpleresty"
	"sync"
	"time"
)

const (
	// DefaultAPIBaseURL is the base URL when making API calls.
	DefaultAPIBaseURL = "https://api.rollbar.com/api/1"

	// DefaultUserAgent is the user agent used when making API calls.
	DefaultUserAgent = "rollbar-go"

	// RollbarAuthHeader is the Authorization header.
	RollbarAuthHeader = "x-rollbar-access-token"
)

// A Client manages communication with the Rollbar API.
type Client struct {
	// clientMu protects the client during calls that modify the CheckRedirect func.
	clientMu sync.Mutex

	// HTTP client used to communicate with the API.
	http *simpleresty.Client

	// BaseURL for API
	BaseURL string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// User agent used when communicating with the Rollbar API.
	UserAgent string

	// Services used for talking to different parts of the Rollbar API.
	Notifications       *NotificationsService
	Projects            *ProjectsService
	ProjectAccessTokens *ProjectAccessTokensService
	Teams               *TeamsService
	Users               *UsersService

	// Custom HTTPHeaders
	customHTTPHeaders map[string]string

	// Account access token
	accountAccessToken string

	// Project access token
	projectAccessToken string
}

// service represents the api service client.
type service struct {
	client *Client
}

// TokenAuthConfig represents options when initializing a new API http.
type TokenAuthConfig struct {
	// ProjectAccessToken is a Rollbar project access token.
	ProjectAccessToken *string

	// AccountAccessToken is a Rollbar account access token.
	AccountAccessToken *string

	// Custom HTTPHeaders
	CustomHTTPHeaders map[string]string
}

// NewClientTokenAuth constructs a new client to interact with the Rollbar API using a project or account access token.
func NewClientTokenAuth(config *TokenAuthConfig) (*Client, error) {
	// Validate that either ProjectAccessToken or AccountAccessToken are set in TokenAuthConfig.
	if config.GetProjectAccessToken() == "" && config.GetAccountAccessToken() == "" {
		return nil, fmt.Errorf("please set an account access token and/or a project access token for authentication")
	}

	// Construct new client.
	c := &Client{
		http: simpleresty.New(), BaseURL: DefaultAPIBaseURL, UserAgent: DefaultUserAgent,
		customHTTPHeaders: config.CustomHTTPHeaders, accountAccessToken: config.GetAccountAccessToken(),
		projectAccessToken: config.GetProjectAccessToken(),
	}
	c.injectServices()

	// Setup client
	c.setupClient()

	return c, nil
}

// injectServices adds the services to the client.
func (c *Client) injectServices() {
	c.common.client = c
	c.Notifications = (*NotificationsService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	c.ProjectAccessTokens = (*ProjectAccessTokensService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
}

// setupClient sets common headers and other configurations.
func (c *Client) setupClient() {
	// We aren't setting an authentication header initially here as certain API resources require specific access_tokens.
	// Per Rollbar API documentation, each individual resource will set the access_token parameter when constructing
	// the full API endpoint URL.
	c.http.SetHeader("Content-type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", c.UserAgent).
		SetTimeout(5 * time.Minute).
		SetAllowGetMethodPayload(true)

	// Set additional headers
	if c.customHTTPHeaders != nil {
		c.http.SetHeaders(c.customHTTPHeaders)
	}
}

func (c *Client) setAuthTokenHeader(token string) {
	c.http.SetHeader(RollbarAuthHeader, token)
}
