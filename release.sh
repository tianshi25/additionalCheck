#!/usr/bin/env bash

package_name="additionalCheck"
platforms=("windows/amd64" "linux/amd64" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
#    echo ${paltform}
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name=${package_name}'_'$GOOS'_'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o './out/'$output_name
done
