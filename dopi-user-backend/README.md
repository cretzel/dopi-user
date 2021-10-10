# dopi-user

user backend for dopi

## run dev
```
gow run main.go
```

## build
```
docker build -t cretzel/dopi-user-backend .
docker push cretzel/dopi-user-backend
```

## login
```
curl -v -X POST -d '{"username": "admin", "password":"admin"}' http://localhost:8081/api/user/login
```