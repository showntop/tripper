# api specification
==========

apps share and perfect

##create user

```
method: post
url:    /users
request body:
    {
        mobile: "",
        password: ""
    }
response body:
    {
        state_code:
        message:
        data:{
            id: 
            username:
            mobile:
        }
    }
```

## user login

```
method: post
url:    /sessions
request body:
    {
        field_value: username/mobile/email,
        password: ""
    }
response body:
    {
        state_code:
        message:
        data:{
            id: 
            username:
            mobile:
        }
    }
```

## get category list

```
method: get
url:    /categories
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,
              "name": ""
            },
        }
    }
```

## get app list

```
method: get
url:    /apps?category_id={:id}&page_no={:no}&page_num={:num}
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,
              "category_id": 16110000,
              "name": "",
              "version": "1.09",
              "description": "",
              "size": 16714301,
              "logo_url": "",
              "download_url": "",
              "Assets": null,
              "tags": [
                "t1",
                "t2",
              ]
            },
        }
    }
```

## get app detail info

```
method: get
url:    /apps/{:app_id}
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,
              "category_id": 16110000,
              "name": "",
              "version": "1.09",
              "description": "",
              "size": 16714301,
              "logo_url": "",
              "download_url": "",
              "Assets": null,
              "tags": [
                "t1",
                "t2",
              ]
            },
        }
    }
```

## get app post list

```
method: get
url:    /apps/{:app_id}/posts?page_no={1}&page_num={30}
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,

            },
        }
    }
```

## create a app post

```
method: post
url:    /apps/{:app_id}/posts
request body:
    request body:
    {
        project_id: 123456,
        password: ""
    }
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,

            },
        }
    }
```

## get app post list

```
method: post
url:    /apps/{:app_id}/posts
response body:
  {
    "data": [
      {
        "id": 10,
        "project_id": 16110019,
        "content": "henhao很好a",
        "user_id": 9,
        "comments": null,
        "like_num": 0,
        "comment_num": 0
      },
      {
        "id": 11,
        "project_id": 16110019,
        "content": "henhao很好a",
        "user_id": 9,
        "comments": null,
        "like_num": 0,
        "comment_num": 0
      },
      {
        "id": 12,
        "project_id": 16110019,
        "content": "henhao很好a",
        "user_id": 9,
        "comments": null,
        "like_num": 0,
        "comment_num": 0
      },
      {
        "id": 13,
        "project_id": 16110019,
        "content": "henhao很好a",
        "user_id": 9,
        "comments": null,
        "like_num": 0,
        "comment_num": 0
      },
      {
        "id": 14,
        "project_id": 16110019,
        "content": "henhao很好a",
        "user_id": 9,
        "comments": null,
        "like_num": 0,
        "comment_num": 0
      }
    ],
    "message": "成功",
    "state_code": 200
  }
```

## like a app post

```
method: put
url:    /posts/{:id}/likes
response body:
  {
    "message": "成功",
    "state_code": 200
  }

## create a app post's post

```
method: put
url:    /posts/{:id}/comments
response body:
  {
    data:{
      id: 
    }
    "message": "成功",
    "state_code": 200
  }
```

