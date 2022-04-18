localstack-setup:
	docker-compose exec localstack aws --endpoint-url=http://localhost:4572 s3 mb s3://s3-example