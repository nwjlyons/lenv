<h1 align="center">lenv</h1>

**lenv** (concatenation of **local** and **env**) executes a command with environment variables defined in a `.env` file. 

## Install

    go get github.com/nwjlyons/lenv
    
## Usage

    Usage:  lenv <command>
    
Example

    $ echo 'PORT=8000' > .env
    
    // Show all enviroment variables
    $ lenv env
    $ PORT=8000
    
    // Run a local HTTP server
    $ lenv go run main.go
    $ Listening on port :8000
