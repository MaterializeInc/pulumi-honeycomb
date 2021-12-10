VERSION ?= $(patsubst v%,%,$(shell git describe))

bin/pulumi-resource-honeycomb: provider/cmd/pulumi-resource-honeycomb/schema.json provider/resources.go
	cd provider && go build -o ../bin/pulumi-resource-honeycomb ./cmd/pulumi-resource-honeycomb

bin/pulumi-tfgen-honeycomb: provider/cmd/pulumi-tfgen-honeycomb/*.go provider/go.sum provider/resources.go
	cd provider && go build -o ../bin/pulumi-tfgen-honeycomb  -ldflags "-X github.com/MaterializeInc/pulumi-honeycomb/provider/pkg/version.Version=${VERSION}" ./cmd/pulumi-tfgen-honeycomb

provider/cmd/pulumi-resource-honeycomb/schema.json: bin/pulumi-tfgen-honeycomb provider/resources.go
	bin/pulumi-tfgen-honeycomb schema --out ./provider/cmd/pulumi-resource-honeycomb
	(cd provider && VERSION=$(VERSION) go generate cmd/pulumi-resource-honeycomb/main.go)

schema: provider/cmd/pulumi-resource-honeycomb/schema.json

python-sdk: bin/pulumi-tfgen-honeycomb provider/cmd/pulumi-resource-honeycomb/schema.json
	rm -rf sdk
	bin/pulumi-tfgen-honeycomb python
	cp README.md sdk/python/
	cd sdk/python/ && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" setup.py && \
		rm setup.py.bak
