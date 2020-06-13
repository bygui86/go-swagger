package blogpost

import "time"

const (
	// configurations
	httpServerHostFormat          = "%s:%d"
	httpServerWriteTimeoutDefault = time.Second * 15
	httpServerReadTimeoutDefault  = time.Second * 15
	httpServerIdelTimeoutDefault  = time.Second * 60

	// url
	urlSeparator           = "/"
	rootUrl                = urlSeparator
	blogPostsEndpoint      = "blogPosts"
	blogPostsRootUrl       = rootUrl + blogPostsEndpoint
	blogPostIdEndpointPath = urlSeparator + idValue
	routesEndpoint         = "routes"
	routesRootUrl          = rootUrl + routesEndpoint
	// swaggerUrl             = urlSeparator + "swagger-ui" + urlSeparator

	// url variables
	idKey   = "id"
	idValue = "{" + idKey + ":[0-9]+}"

	// headers keys
	contentTypeHeaderKey = "Content-Type"
	acceptHeaderKey      = "Accept"

	// headers values
	contentTypeAppJson = "application/json"

	// messages
	rootMessage = "Welcome to blog-post REST application"
)
