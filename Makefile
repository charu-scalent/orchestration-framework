run-wallet:
	cd ./services/wallet/app && wire
	cd ./services/wallet/app/ && go run .