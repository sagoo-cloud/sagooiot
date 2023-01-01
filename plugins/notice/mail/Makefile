BINARY_NAME=notice-mail

local:
	echo "========local============"
	go build -o ${BINARY_NAME}
	mv ${BINARY_NAME} ../../built

linux:
	echo "========linux============"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}
	mv ${BINARY_NAME} ../../built