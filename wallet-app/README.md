# itk-academy-test-app

## Установка
* make up_build
* make up

## Benchmark
* make benchmark

## Доступные UUID
* 550e8400-e29b-41d4-a716-446655440002
* 550e8400-e29b-41d4-a716-446655440003
* 550e8400-e29b-41d4-a716-446655440001

## Методы API
* POST http://localhost:8080/api/v1/wallet
```json
{
    "wallet_id": "550e8400-e29b-41d4-a716-446655440001",
    "amount": 100,
	"operation_type": "DEPOSIT"
}
```
* GET http://localhost:8080/api/v1/wallet/550e8400-e29b-41d4-a716-446655440001