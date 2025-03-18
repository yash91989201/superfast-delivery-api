module github.com/yash91989201/superfast-delivery-api/services/geolocation

go 1.24.0

replace (
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5
	google.golang.org/genproto/googleapis/rpc => google.golang.org/genproto/googleapis/rpc v0.0.0-20230803162519-f966b187b2e5
)

require (
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nrednav/cuid2 v1.0.1
	github.com/redis/go-redis/v9 v9.7.1
	github.com/tinrab/retry v1.0.0
	github.com/yash91989201/superfast-delivery-api/common v0.0.0-20250309055017-edc39db580b5
	google.golang.org/grpc v1.71.0
	googlemaps.github.io/maps v1.7.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	go.mongodb.org/mongo-driver/v2 v2.1.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
