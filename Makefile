run-wallet:
	cd ./services/wallet/app && wire
	cd ./services/wallet/app/ && go run .

run-order:
	cd ./services/order/app && wire
	cd ./services/order/app/ && go run .
