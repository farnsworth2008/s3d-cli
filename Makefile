s3d: m
	cp m s3d

m: *.go 
	go mod tidy && go build . && cp m s3d && ./s3d version
