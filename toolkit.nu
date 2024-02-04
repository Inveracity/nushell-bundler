#!/usr/bin/env nu
export def main [] {
    print "use --help for more info"
}

# Bundle nushell scripts to a binary
export def "main build" [
    name: string = "mytool", # name of the binary
] {
    cp $nu.current-exe ./nu
    ^go build -ldflags="-s -w" -o $name
    rm ./nu

    # silently compress binary if UPX is available
    do {upx -1 --force-overwrite $name} | complete

    ls $name
}

# Run "go run main.go"
export def "main run" [] {
    cp $nu.current-exe ./nu
    ^go run main.go
    rm ./nu
}
