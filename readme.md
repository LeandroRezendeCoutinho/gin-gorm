### Gin + GORM

Simple implementation REST API for Gin with GORM

## Explaining test
- Use benchmarck tool in route /dogs
    - `autocannon http://localhost:4500/dogs`
- Populate database with 10 records
```curl
curl --location --request GET 'localhost:4500/dogs' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Ralph",
    "breed":"Pastor",
    "age": 12,
    "isGoodBoy":true
}'
```

#### Benchmark
![bench](https://github.com/LeandroRezendeCoutinho/gin-gorm/blob/main/img/gin_gorm_bench.png)

#### Memory and CPU
![bench](https://github.com/LeandroRezendeCoutinho/gin-gorm/blob/main/img/gin_mem.png)

References:
- https://gin-gonic.com/
- https://gorm.io/
