# BlogPost API


<a name="overview"></a>
## Overview
Here we provide a detailed overview of how to use BlogPost API.


### Version information
*Version* : 1.0.0


### Contact information
*Contact* : Matteo Baiguini  
*Contact Email* : bygui86@github.com


### License information
*License* : Apache 2.0  
*License URL* : http://www.apache.org/licenses/LICENSE-2.0.html  
*Terms of service* : http://swagger.io/terms/


### URI scheme
*Host* : localhost:8080  
*BasePath* : /




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




<a name="definitions"></a>
## Definitions

<a name="blogpost-blogpost"></a>
### blogpost.blogPost

|Name|Schema|
|---|---|
|**body**  <br>*optional*|string|
|**id**  <br>*optional*|string|
|**title**  <br>*optional*|string|





