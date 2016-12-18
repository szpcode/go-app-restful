# go-app-restful
Application restful with dependency Injection, design patterns MVC, use of the library szpcode/tdb


### Example 
READ
``` sh
curl -i -X GET "http://localhost/person?id=1"
```

DELETE
``` sh
curl -i -X DELETE "http://localhost/person?id=1"
```
EDIT
``` sh
curl -i -X PUT "http://localhost/person?id=2&name=Anabel&surname=Rose&birthday=1981-01-08"
```
SEARCH
``` sh
curl -i -X GET "http://localhost/personList?surname=ro"
```
CREATE
``` sh
curl -i -X POST "http://localhost/person?id=2&name=Annie&surname=Woodward&birthday=1975-11-23"
```
