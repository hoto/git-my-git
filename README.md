[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/git-my-git/workflows/Build%20and%20test/badge.svg?branch=master)](https://github.com/hoto/git-my-git/actions)
[![Release](https://img.shields.io/github/release/hoto/git-my-git.svg?style=flat-square)](https://github.com/hoto/git-my-git/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)

# Git my git

Navigitae locally cloned repos.

### TODO:

* Fix config vars by using goreleaser default ldflags
* Use GPG signing for commits and releases

### Installation
    
Mac:

    brew install hoto/tap/git-my-git

Mac or Linux:

    sudo curl -L \
      "https://github.com/hoto/git-my-git/releases/download/1.0.0/git-my-git_1.0.0_$(uname -s)_$(uname -m)" \
       -o /usr/local/bin/git-my-git

    sudo chmod +x /usr/local/bin/git-my-git
    
Or manually download binary from [releases](https://github.com/hoto/git-my-git/releases).
    
### Configuration and running

Add to your `~/.bashrc` or `~/.zshrc` or `~/.profile`:  

    alias g='git-my-git $@'

In terminal:

    $ g
    
Find projects by partial name:

    $ g myprojectname
        
Help:
  
    $ git-my-git --help
    
### Development

Get:

    go get github.com/hoto/git-my-git

Download dependencies:

    make dependencies

Build, test and run:

    make clean
    make build
    make test
    make run
    
    ./bin/git-my-git

Run with arguments:

    make args="myargs" run

Install to global golang bin directory:

    make install