# Develop Guild

running commands
```
consuladdr=http://localhost:8500 go run app/service/wallet_key/cmd/main.go

v_address=http://127.0.0.1:8200 v_token=s.I0kFnO89SY5vZZ2VAWCJW79V v_path=tk/wallet go run app/service/wallet_management/cmd/main.go

v_address=http://127.0.0.1:8200 v_token=s.I0kFnO89SY5vZZ2VAWCJW79V v_path=tk/tx go run app/service/transaction/cmd/main.go

v_address=http://127.0.0.1:8200 v_token=s.I0kFnO89SY5vZZ2VAWCJW79V v_path=tk/dashboard go run app/service/dashboard/cmd/main.go

v_address=http://127.0.0.1:8200 v_token=s.I0kFnO89SY5vZZ2VAWCJW79V v_path=tk/chainsquery go run app/service/chains_query/cmd/main.go

v_address=http://127.0.0.1:8200 v_token=s.I0kFnO89SY5vZZ2VAWCJW79V v_path=tk/account go run app/service/account/cmd/main.go

consuladdr=http://localhost:8500 go run app/gateway/webapi/cmd/main.go

npm run dev
```
