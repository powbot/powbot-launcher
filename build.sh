#!/bin/bash

platforms=("windows/amd64" "windows/386" "darwin/amd64")
for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	
	output_name="PowBotLauncher-$GOOS-$GOARCH"
	
	if [ $GOOS = "windows" ]; then
	    output_name+='.exe'
	fi

	GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
done
