[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/git-my-git/workflows/Build%20and%20test/badge.svg?branch=master)](https://github.com/hoto/git-my-git/actions)
[![Release](https://img.shields.io/github/release/hoto/git-my-git.svg?style=flat-square)](https://github.com/hoto/git-my-git/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)

# Git my git

Navigitae locally cloned repos.


# WIP TODO

    git-my-git (show help)
    git-my-git help
    git-my-git --help
    git-my-git regenerate --paths "~/projects, ~/go/src/github.com/hoto"
    git-my-git reset --frequency|last_accessed_date
    git-my-git find myprojectname
    git-my-git imfeelinglucky myprojectname (navigate to first found project)
    git-my-git imfeelinglucky --weight alphabetic|frequency myprojectname

> ~/.config/git-my-git/git-my-git.yml

```yaml
projects:
  - path: "/home/.../hello-pony"
    frequency: 0
    created_date: "2020-11-03:10:11:02:UTC"
    last_accessed_date: "2020-11-03:10:11:02:UTC"
  - path: "/home/.../project-zeta"
    frequency: 0
    created_date: "2020-11-03:10:11:02:UTC"
    last_accessed_date: "2020-11-03:10:11:02:UTC"
  - path: "/home/.../im-hungry"
    frequency: 2
    created_date: "2020-11-03:10:11:02:UTC"
    last_accessed_date: "2020-11-03:10:11:02:UTC"
```

### Installation
    
Mac:

    brew install hoto/repo/git-my-git

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

Build and test:

    go get github.com/hoto/git-my-git
    
    make dependencies build test
    
Build binary:

     make build
    ./bin/git-my-git

Run with arguments:

    make args="myargs" run

Install to global golang bin directory:

    make install
    git-my-git
    
    