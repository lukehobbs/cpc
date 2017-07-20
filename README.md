[![Go Report Card](https://goreportcard.com/badge/github.com/lukehobbs/cpc)](https://goreportcard.com/report/github.com/lukehobbs/cpc)

# CPC Pipeline Control

A CLI for your commit messages!

## Installation
Get binaries for OS X/Linux from the latest [release](https://github.com/lukehobbs/cpc/releases/latest) as one of the first steps in your pipeline.

## Usage
Example Jenkins usage for cpc using [cpc.yaml](./cpc.yaml) looks like:

```
def nodeVars =
  [
    "PATH+RUBY=/usr/bin/ruby-2.4.0/bin/",
    "PATH+GO=/usr/bin/go"
  ]

stage('First') {
  node('my-agent') {
    withEnv(nodeVars) {
      deleteDir()
      checkout scm
      MESSAGE = sh (script: "git log --pretty=%s -1", returnStdout: true).trim()
      CPC = sh (script: "./cpc ${MESSAGE}", returnStdout: true).trim()
      nodeVars = nodeVars.addAll(CPC.tokenize(', '))
      ...
    }
  }
}

stage('Second') {
  println "My CPC variables values: "
  println "full: ${env.FULL}"
  println "name: ${env.NAME}"
  println "persist: ${env.PERSIST}"
}
```

If this pipeline is run from a commit message like: "Adding my new features. cpc --full -n Luke -p 25", the console output will be:

```
[Pipeline] sh
My CPC variables values:
full: true
name: Luke
persist: 25
```

You can then implement logic to pivot upon these variables in anyway you please. The `name` variable is a popular one in order to give your pipeline's stacks unique and human-readable names.

## Constructing your own cpc.yaml

You will need a `cpc.yaml` file at the root of your project in order to define the flags for cpc to search for. These can be defined in an easy to read manner. The only types of flags currently implemented are:

- Boolean flags that default to false
- String flags
- Integer flags

These can be defined as follows:

```
FlagType:
  - Name: "long, short"
    Usage: "Usage string"
```

For example,

```
BoolFlags:
  - Name: "name, n"
    Usage: "Use this flag to associate a name with the build."
```

The current flag types are: `BoolFlags, StringFlags, IntFlags`.

----------

Feel free to contact me with any questions you may have. Pull Requests are welcome as well as feature requests if you find CPC to be missing something- just create an issue :-)
