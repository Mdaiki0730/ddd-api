.PHONY: docker-up
docker-up:
	@docker-compose -f docker-compose.local.yml up -d

.PHONY: docker-down
docker-down:
	@docker-compose -f docker-compose.local.yml down

.PHONY: docker-stop
docker-stop:
	@docker-compose -f docker-compose.local.yml stop

.PHONY: docker-build
docker-build:
	@docker-compose -f docker-compose.local.yml build

.PHONY: watch-log
watch-log:
	@docker-compose -f docker-compose.local.yml logs -f api

.PHONY: gen-proto
gen-proto:
	@sh ./scripts/gen-proto.sh
