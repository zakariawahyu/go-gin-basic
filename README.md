# API Documentation

### Get All Albums
```http request
GET /albums
```
```javascript
Response
[
    {
        "id": "1",
        "title": "Resah",
        "artist": "Payung Teduh",
        "price": 50.5
    },
    {
        "id": "2",
        "title": "Senja di Ambang Pilu",
        "artist": "Danilla",
        "price": 78.999
    }
]
```
### Get Albums By ID
```http request
GET /albums/:id
```
```javascript
Response
{
    "id": "4",
    "title": "Menuju Senja",
    "artist": "Payung Teduh",
    "price": 45.788
}
```

### Delete Albums By ID
```http request
DELETE /albums/:id
```

### Create Albums
```http request
POST /albums
```
```javascript
Body Request
{
    "id": "3",
    "title": "Senja di Ambang Pilu",
    "artist": "Danilla",
    "price": 78.999
}
```