
<a name="paths"></a>
## Paths

<a name="blogposts-get"></a>
### Get all blog posts
```
GET /blogPosts
```


#### Description
Return all blog posts in memory


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|< [blogpost.blogPost](#blogpost-blogpost) > array|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### Tags

* blogPosts


<a name="blogposts-id-get"></a>
### Get a specific blog post given the ID
```
GET /blogPosts/{id}
```


#### Description
If a blog posts with specified ID exists, it will be returned else nothing will be returned


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**id**  <br>*required*|ID of the blog post to look for|integer|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|OK|[blogpost.blogPost](#blogpost-blogpost)|


#### Consumes

* `application/json`


#### Produces

* `application/json`


#### Tags

* blogPosts



