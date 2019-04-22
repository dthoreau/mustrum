# Targets needed:
# - build
# - test
# - fetch tools
# - swagger build
# - update glide packages 
#
swagger:
	swagger generate server -f server/swagger.yaml

build:
	go build
