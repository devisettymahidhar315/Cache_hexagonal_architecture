# multi backend caching library implemented in hexagonal architecture and gin framework

 **Run the Program**
   ```bash
      go run main.go length
   ```
   length means number of elements are storing
   ```bash
      go run main.go
   ```
###  without length if we run the program, it takes the default length has 2

# Functions Present in the Project
### `get`  
### `post`
### `delete`

# Accessing the Functions we are using the postman
## Get Function
### you can select the get method 
### printing the particular key ```http://localhost:8080/key```
### printing the entire data ```http://localhost:8080/print```

## Delete Function
### delete the data
### you can select the delete method
### delete partcular key  ```http://localhost:8080/key```
### delete entire data ```http://localhost:8080/all```

## Post Function
### store the data
### Implement the "time to live" (TTL) feature. If the time is -1, it means no time is set. If the time is greater than 1, the key-value pair will be deleted after that many seconds
### select the post method and type the following command ```http://localhost:8080/key/value/time```
