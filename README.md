<h1 align="center">lenv</h1>

**lenv** (concatenation of **local** and **env**) executes a command with environment variables defined in a `.env` file. 

## Install

    go get github.com/nwjlyons/lenv
    
## Usage

    Usage:  lenv <command>
    
Example

    $ echo 'PORT=8000' > .env
    
    $ lenv env
    $ PORT=8000
