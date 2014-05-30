OS_LIST="linux darwin windows"
ARCH_LIST="amd64"

for os in $OS_LIST; do
	for arch in $ARCH_LIST; do
		echo "building for $os $arch"
		bin="fluentail"
		[ $os == "windows" ] && bin="${bin}.exe"
		GOOS=$os GOARCH=$arch go build

		zip fluentail_$os_$arch.zip $bin > /dev/null
		rm $bin
	done
done
#release:
#	GOOS=linux GOARCH=amd64 go build
#	zip fluentail_linux_amd64.zip fluentail
#	rm fluentail
#	GOOS=darwin GOARCH=amd64 go build
#	zip fluentail_darwin_amd64.zip fluentail
#	rm fluentail
