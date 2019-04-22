# Targets needed:
# - build
# -  test
# -  fetch tools
# -  swagger build
#
swagger:
	swagger generate server -f server/swagger.yaml
