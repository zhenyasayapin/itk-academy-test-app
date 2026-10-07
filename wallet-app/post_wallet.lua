wrk.method = "POST"

wrk.body = [[
{
  "wallet_id": "550e8400-e29b-41d4-a716-446655440001",
  "amount": 5,
  "operation_type": "DEPOSIT"
}
]]

wrk.headers["Content-Type"] = "application/json"
wrk.headers["Accept"] = "application/json"
