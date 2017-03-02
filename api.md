login
```
path:/api/login
method:post

{
    "password": "123",
    "username": "lxw"
}

返回数据：
{
"code": 0,
"data": {
  "id": 1,
  "username": "lxw",
  "password": "",
  "phone": "",
  "email": "",
  "pic": "",
  "money": 0,
  "token": "f208d78fc6fa3fff68a06317143cd06e",
  "CreatedAt": "0001-01-01T00:00:00Z",
  "UpdateAt": "2017-02-28T23:14:39.396053427+08:00"
},
"msg": "success"
}
```

register

```
path:/api/register
method:post
{
    "email": "935232474@qq.com",
    "password": "123",
    "username": "lxw2"
}

返回
{
    "code": 0,
    "data": {
        "id": 2,
        "username": "lxw2",
        "password": "123",
        "phone": "",
        "email": "935232474@qq.com",
        "pic": "",
        "money": 0,
        "token": "239f312e04a2cddf795495c20e0e57ce",
        "CreatedAt": "2017-02-28T23:30:34.276041887+08:00",
        "UpdateAt": "2017-02-28T23:30:34.276407928+08:00"
    },
    "msg": "success"
}

```

获取个人信息
```
path:/api/user
method:get
{
  "code": 0,
  "data": {
    "id": 2,
    "username": "lxw2",
    "phone": "",
    "email": "935232474@qq.com",
    "pic": "",
    "money": 0,
    "token": "239f312e04a2cddf795495c20e0e57ce",
    "CreatedAt": "2017-02-28T23:30:34+08:00",
    "UpdateAt": "2017-02-28T23:30:34+08:00"
  },
  "msg": "success"
}
```

更新个人信息
```
path:/api/user
method:put

{
	"username":"ss1",
	"pic":"http://a4.topitme.com/o017/1001791526cae12549.jpg"
}


{
  "code": 0,
  "data": {
    "id": 2,
    "username": "luxingwen",
    "phone": "",
    "email": "",
    "pic": "http://a4.topitme.com/o017/1001791526cae12549.jpg",
    "money": 0,
    "token": "",
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdateAt": "2017-03-01T11:02:19.599319155+08:00"
  },
  "msg": "success"
}
```

添加图书
```
path:/api/book
method:post
{
    "name": "编程人生",
    "autor": "甩甩",
    "desc": "我是描述",
    "pic": "",
    "category": 1
}

{
    "code": 0,
    "data": {
        "id": 1,
        "name": "编程人生",
        "autor": "甩甩",
        "desc": "我是描述",
        "pic": "",
        "category": 1,
        "CreatedAt": "2017-03-01T00:23:35.1330871+08:00",
        "UpdateAt": "2017-03-01T00:23:35.133094206+08:00"
    },
    "msg": "success"
}
```


更新图书信息
```
path:/api/book/id
id是图书的id
method:put
{
    "id": 1,
    "name": "代码大全",
    "autor": "甩甩2",
    "desc": "我是描述",
    "pic": "",
    "category": 2
}

 {
    "code": 0,
    "data": {
        "id": 1,
        "name": "代码大全",
        "autor": "甩甩2",
        "desc": "我是描述",
        "pic": "",
        "category": 2,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdateAt": "2017-03-01T00:33:27.227677513+08:00"
    },
    "msg": "success"
}
```


添加评论
```
path:/api/book/info/id/community
id是图书的id
method:post
{
	"content": "我是甩甩的评论。。。"
}


{
    "code": 0,
    "data": {
        "id": 1,
        "content": "我是甩甩的评论。。。",
        "bookId": 1,
        "userId": 2,
        "CreatedAt": "2017-03-01T00:46:13.912301035+08:00"
    },
    "msg": "success"
}

```


购买图书
```
path:/api/book/buy
id是图书的id
method:post

{
	"bookId":1
}
{
    "code": 0,
    "data": {
        "id": 1,
        "name": "代码大全",
        "autor": "甩甩2",
        "desc": "我是描述",
        "pic": "",
        "category": 2,
        "CreatedAt": "2017-03-01T00:23:35+08:00",
        "UpdateAt": "2017-03-01T00:33:27+08:00"
    },
    "msg": "success"
}
```


图书信息
```
path:/api/book/info/id
id是图书的id
method:get
{
  "code": 0,
  "data": {
    "book": {
      "id": 1,
      "name": "代码大全",
      "autor": "甩甩2",
      "desc": "我是描述",
      "pic": "",
      "money": 0,
      "category": 2,
      "CreatedAt": "2017-03-01T00:23:35+08:00",
      "UpdateAt": "2017-03-01T00:33:27+08:00"
    },
    "community": [
      {
        "id": 1,
        "content": "我是甩甩的评论。。。",
        "user": {
          "id": 2,
          "username": "luxingwen",
          "password": "123",
          "phone": "",
          "email": "935232474@qq.com",
          "pic": "http://a4.topitme.com/o017/1001791526cae12549.jpg",
          "money": 0,
          "token": "239f312e04a2cddf795495c20e0e57ce",
          "CreatedAt": "2017-02-28T23:30:34+08:00",
          "UpdateAt": "2017-03-01T11:02:19+08:00"
        },
        "CreatedAt": "2017-03-01T00:46:13+08:00"
      }
    ]
  },
  "msg": "success"
}
```

获取图书列表
```
path:/api/books
method:get

{
    "code": 0,
    "data": [
        {
            "id": 1,
            "name": "代码大全",
            "autor": "甩甩2",
            "desc": "我是描述",
            "pic": "",
            "category": 2,
            "CreatedAt": "2017-03-01T00:23:35+08:00",
            "UpdateAt": "2017-03-01T00:33:27+08:00"
        }
    ],
    "msg": "success"
}
```

通过分类获取图书列表 
```
path:/api/books/id
id是分类id
method:get
{
    "code": 0,
    "data": [
        {
            "id": 1,
            "name": "代码大全",
            "autor": "甩甩2",
            "desc": "我是描述",
            "pic": "",
            "category": 2,
            "CreatedAt": "2017-03-01T00:23:35+08:00",
            "UpdateAt": "2017-03-01T00:33:27+08:00"
        }
    ],
    "msg": "success"
}
```

获取我的图书列表

```
{
  "code": 0,
  "data": [
    {
      "id": 1,
      "name": "代码大全",
      "autor": "甩甩2",
      "desc": "我是描述",
      "pic": "",
      "money": 0,
      "category": 2,
      "CreatedAt": "2017-03-01T00:23:35+08:00",
      "UpdateAt": "2017-03-01T00:33:27+08:00"
    }
  ],
  "msg": "success"
}
```




添加购物车
```
path:/api/shopcar/id
id是图书id
method:post
{
	"bookId":1
}

{
    "code": 0,
    "data": {
        "id": 0,
        "bookId": 1,
        "userId": 2,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdateAt": "0001-01-01T00:00:00Z"
    },
    "msg": "success"
}
```


购物車列表
```
path:/api/shopcar
method:get

{
  "code": 0,
  "data": [
    {
      "id": 1,
      "name": "代码大全",
      "autor": "甩甩2",
      "desc": "我是描述",
      "pic": "",
      "money": 0,
      "category": 2,
      "CreatedAt": "2017-03-01T00:23:35+08:00",
      "UpdateAt": "2017-03-01T00:33:27+08:00"
    }
  ],
  "msg": "success"
}

```


获取图书分类信息

```
path:/api/book/category
method:get

{
  "code": 0,
  "data": [
    {
      "id": 1,
      "content": "IT",
      "fatherId": 0
    },
    {
      "id": 2,
      "content": "文学",
      "fatherId": 0
    }
  ],
  "msg": "success"
}
```

文件上传方式一
```
path:/api/file/filename
method:post
filename是文件名称

用法示范
curl -X POST \
  -H "Content-Type: image/jpeg" \
  --data-binary '@myPicture.jpg' \
  /api/file/filename?token=?
  
```


文件上传方式2
```
path:/api/file2
method:post

用法：
浏览器打开 /api/file2

```