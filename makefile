db-create:
	docker run --name bill-service -e POSTGRES_PASSWORD=passwd -p 5432:5432 -d postgres

