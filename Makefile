VERSION ?= $(patsubst v%,%,$(shell git describe))

bin/pulumi-resource-honeycomb: provider/cmd/pulumi-resource-honeycomb/schema.json provider/resources.go
	cd provider && go build -o ../bin/pulumi-resource-honeycomb -ldflags "-X github.com/MaterializeInc/pulumi-honeycomb/provider/pkg/version.Version=${VERSION}" ./cmd/pulumi-resource-honeycomb

bin/pulumi-tfgen-honeycomb: provider/cmd/pulumi-tfgen-honeycomb/*.go provider/go.sum provider/resources.go
	cd provider && go build -o ../bin/pulumi-tfgen-honeycomb -ldflags "-X github.com/MaterializeInc/pulumi-honeycomb/provider/pkg/version.Version=${VERSION}" ./cmd/pulumi-tfgen-honeycomb

provider/cmd/pulumi-resource-honeycomb/schema.json: bin/pulumi-tfgen-honeycomb provider/resources.go
	bin/pulumi-tfgen-honeycomb schema --out ./provider/cmd/pulumi-resource-honeycomb
	(cd provider && VERSION=$(VERSION) go generate cmd/pulumi-resource-honeycomb/main.go)

schema: provider/cmd/pulumi-resource-honeycomb/schema.json

python-sdk: bin/pulumi-tfgen-honeycomb provider/cmd/pulumi-resource-honeycomb/schema.json
	rm -rf sdk/python
	bin/pulumi-tfgen-honeycomb python
	cp README.md sdk/python/
	cd sdk/python/ && \
		sed -i.bak -e 's/VERSION = "0.0.0"/VERSION = "${VERSION}"/g' -e 's/PLUGIN_VERSION = "0.0.0"/PLUGIN_VERSION = "${VERSION}"/g' setup.py && \
		rm setup.py.bak

nodejs-sdk: bin/pulumi-tfgen-honeycomb provider/cmd/pulumi-resource-honeycomb/schema.json
	rm -rf sdk/nodejs
	bin/pulumi-tfgen-honeycomb nodejs
	cd sdk/nodejs && \
		sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./package.json
	# Need to ensure we have copied the LICENSE and README to the bin folder before publishing
	# also need to inject the correct version string

dotnet-sdk: bin/pulumi-tfgen-honeycomb provider/cmd/pulumi-resource-honeycomb/schema.json
	rm -rf sdk/dotnet
	bin/pulumi-tfgen-honeycomb dotnet
	cd sdk/dotnet/ && \
		echo "${VERSION}" >version.txt && \
		dotnet build /p:Version=${ERSION}
	# Add the build step and inject the version to VERSION.txt

go-sdk: bin/pulumi-tfgen-honeycomb provider/cmd/pulumi-resource-honeycomb/schema.json
	rm -rf sdk/go
	bin/pulumi-tfgen-honeycomb go
