# go-gitcmdwrapper

* This is wrapper for git.
  
* Maybe, you installed git. :wink:

## Install

```bash
go get github.com/withnic/go-gitcmdwrapper
```

## Usage

```
import(
    gitwrapper "github.com/withnic/go-gitcmdwrapper"
    )


func main() {
    _, err := gitwrapper.Can()
    if err != nil {
        // Error, Maybe not found
    }

    // git exec
    gitwrapper.Exec(os.Args...)

    // checkout develop branch
    gitwrapper.Checkout("develop")

    // name equals current branch name.
    name := gitwrapper.CurrentBranchName()
}
