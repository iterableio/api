ENV_KEY := ITERABLE_ENVIRONMENT
CONFIG_DIR_KEY := ITERABLE_CONFIG_DIR

root_dir := $(shell pwd)

test:
	export $(CONFIG_DIR_KEY)=$(root_dir)/config; export $(ENV_KEY)=test; go test -cover ./...
