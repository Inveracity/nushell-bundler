use boop

def main [
    --blue (-b), # hello world, but in blue
]: nothing -> nothing {
    if $blue {
        print $"(ansi blue)hello world(ansi reset)"
        return
    }

    print "hello world"
}

# Demonstrate that modules also work
def "main boop" [] {
    boop
}

def "main error" [] {
    error make {msg: "this is an error" }
}
