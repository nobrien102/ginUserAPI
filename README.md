# ginUserAPI
https://www.youtube.com/watch?v=lrIR3caBlaM
Followed quick gin golang tutroial in order to start learning about creating APIs
No database was used so to receive users you need to first add users using the CreateUser endpoint.

## Endpoints

### /users/getUsers

  GET a slice of users
  
 ### /users/CreateUsers
 
 POST a User in the JSON format
 `{
  "name": [NAME_STRING],
  "age": [AGE_INT],
 }`
  A user Id is created from the uuid package and added to the user object
  
  ### user/EditUser/:id
  
   PUT: replace a user with the given id, if id is not found return an error
    
   ### user/ Delete/:id
   
   DELETE a user of the given id, if id is not found return an eror
  
