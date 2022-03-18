<h1 align="center">lenv</h1>

**lenv** (concatenation of **local** and **env**) executes a command with environment variables defined in a `.env` file. 

## Install

    go install github.com/nwjlyons/lenv@latest
    
## Usage

    Usage:  lenv <command>
    
Example

    $ echo 'PORT=8000' > .env
    
    // Show all environment variables
    $ lenv env
    $ PORT=8000
    
    // Run a local HTTP server
    $ lenv go run main.go
    $ Listening on port :8000
