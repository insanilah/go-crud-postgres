### Install Commitizen dan Husky
```
npm install --save-dev commitizen cz-conventional-changelog husky
```

### Endpoint CRUD
#### POST /users
```
curl -X POST http://localhost:8080/users -d '{"username": "john_doe", "email": "john.doe@example.com"}' -H "Content-Type: application/json"
```

#### GET /users
```
curl http://localhost:8080/users
```

#### Other Info