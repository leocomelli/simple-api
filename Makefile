
#!/usr/bin/make
 
.DEFAULT_GOAL := releases

DIST_DIR := dist
PLATFORMS := linux/amd64 darwin/amd64 windows/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

releases: $(PLATFORMS)

$(PLATFORMS):
	go mod download
	GOOS=$(os) GOARCH=$(arch) go build $(LD_FLAGS) -o 'dist/simple-api_$(os)-$(arch)'
