ENV_KEY := ITERABLE_ENVIRONMENT
CONFIG_DIR_KEY := ITERABLE_CONFIG_DIR

root_dir := $(shell pwd)

test:
	db/migrate reset; export $(CONFIG_DIR_KEY)=$(root_dir)/config; go test -v -cover ./...
