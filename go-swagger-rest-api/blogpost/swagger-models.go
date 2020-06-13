package blogpost

// BlogPost response payload
// swagger:response blogPostRes
type swaggBlogPostRes struct {
	// in:body
	Body blogPost
}

// BlogPosts response payload
// swagger:response blogPostsRes
type swaggBlogPostsRes struct {
	// in:body
	Body []blogPost
}
