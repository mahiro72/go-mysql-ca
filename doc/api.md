# API仕様書


### 現在のタスクをすべて取得するAPI

#### request
```
GET /task/all
```

#### response
Status : 200
```
{
    "data": [
        {
            "id": 1,
            "name": "task1",
            "status": "done"
        },
        {
            "id": 2,
            "name": "task2",
            "status": "todo"
        }
    ]
}
```

<br>

### 新しいタスクを保存するAPI

#### request
```
POST /task/create
```

#### request body
name(必須)
```
{
    "name":"hoge2"
}
```

#### response
Status : 200
```
{
    "data": {
        "id": 0,
        "name": "hoge2",
        "status": ""
    }
}
```

<br>

### タスクのステータスを更新するAPI

#### request
```
PUT /task/status
```

#### request body
id(必須)
status(必須)
```
{
    "id":1,
    "status":"done"
}
```

#### response
Status : 200
```
{
    "data": {
        "id": 1,
        "name": "",
        "status": "done"
    }
}
```
