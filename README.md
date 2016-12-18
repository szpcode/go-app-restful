# go-app-restful
Application restful with dependency Injection, design patterns MVC, use of the library szpcode/tdb


### Example 
Read
``` sh
curl -i -X GET "http://localhost/person?id=1"
```

Delete
``` sh
curl -i -X DELETE "http://localhost/person?id=1"
```
Edit
``` sh
curl -i -X PUT "http://localhost/person?id=2&name=Anabel&surname=Rose&birthday=1981-01-08"
```
Search
``` sh
curl -i -X GET "http://localhost/personList?surname=ro"
```
Create
``` sh
curl -i -X POST "http://localhost/person?id=2&name=Annie&surname=Woodward&birthday=1975-11-23"
```
