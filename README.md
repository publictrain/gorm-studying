# testコマンド
```jsx
curl -X POST http://localhost:8080/add -H "Content-Type: application/json" -d '{
  "CompanyName": "Example Company",
  "Status": "Active",
  "IndustryID": 1,
  "CountryID": 2
}'
```

```jsx
curl -X POST http://localhost:8080/filter -H "Content-Type: application/json" -d '{
  "Status": "Active"
}'

curl -X POST http://localhost:8080/filter -H "Content-Type: application/json" -d '{
  "IndustryID": 1
}'

curl -X POST http://localhost:8080/filter -H "Content-Type: application/json" -d '{
  "CountryID": 2
}'

curl -X POST http://localhost:8080/filter -H "Content-Type: application/json" -d '{
  "CompanyName": "Example Company",
  "Status": "Active"
}'
```
