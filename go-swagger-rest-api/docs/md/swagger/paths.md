
<a name="paths"></a>
## Paths

<a name="getblogposts"></a>
### Get all blog posts
```
GET /blogPosts
```


#### Description
Return all blog posts in memory


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|BlogPosts response payload|< [blogPost](#blogpost) > array|


#### Tags

* blogPosts


<a name="getblogpostbypath"></a>
### Get a specific blog post given the ID
```
GET /blogPosts/{id}
```


#### Description
If a blog posts with specified ID exists, it will be returned else nothing will be returned


#### Parameters

|Type|Name|Description|Schema|
|---|---|---|---|
|**Path**|**id**  <br>*required*|id of the blog post|string|


#### Responses

|HTTP Code|Description|Schema|
|---|---|---|
|**200**|BlogPost response payload|[blogPost](#blogpost)|


#### Tags

* blogPosts



