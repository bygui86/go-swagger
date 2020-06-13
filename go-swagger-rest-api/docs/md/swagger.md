# BlogPost API


<a name="overview"></a>
## Overview
Here we provide a detailed overview of how to use BlogPost API.


### Version information
*Version* : 1.0.0


### Contact information
*Contact* : Matteo Baiguini  
*Contact Email* : bygui86@github.com


### URI scheme
*Host* : localhost:8080  
*Schemes* : HTTP


### Consumes

* `application/json`


### Produces

* `application/json`




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




<a name="definitions"></a>
## Definitions

<a name="blogpost"></a>
### blogPost
BlogPost
{
"id": "1",
"title": "sample title",
"body":  "sample body",
}


|Name|Schema|
|---|---|
|**body**  <br>*optional*|string|
|**id**  <br>*optional*|string|
|**title**  <br>*optional*|string|





