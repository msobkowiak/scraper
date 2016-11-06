# Amazon movies scraper
 
Amazon movies scraper is a service that makes a background request, fetching the respective Amazon web site, parsing it and giving back a valid result to the requester. When requesting an Amazon ID from this API, we will get back accessible and meaningful results in the JSON format provided below.
 
```json
{
	"title": "Um Jeden Preis",
	"release_year": 2013,
	"actors": ["Dennis Quaid", "Zac Efron"],
	"poster": "http://ecx.images-amazon.com/images/I/51UZ8st2OdL._SX200_QL80_.jpg",
	"similar_ids": ["B00SWDQPOC", "B00RBPBO1G", "B00S2EMECI", "B00M5GH53M", "B00IH8BA3S", "B00M5JP1DA"]
}
```

# Build and run 

```
# build an executable file
go build 

# run the application
./screper
```
