#! /usr/bin/env groovy

node {
  stage('Version Control and Branching Strategy') {
    deleteDir()
    checkout scm
    env.GIT_MSG = sh(
      script: "git log --pretty=%s -1",
      returnStdout: true
    ).trim()
  }
  stage('commitArgs Parsing') {
    sh "export GOPATH=$GOROOT/bin"
    sh "go get ./..."
    sh "/usr/bin/go install"
    sh "echo env.GIT_MSG"
    sh "commitArgs ${env.GIT_MSG}"
  }
}
