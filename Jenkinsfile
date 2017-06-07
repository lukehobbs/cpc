#! /usr/bin/env groovy

node {
  stage('Version Control and Branching Strategy') {
    deleteDir()
    checkout scm
    env.GIT_MSG = sh(
      script: "git log --pretty=%s -1",
      returnStdout: true
    ).trim()
    echo env.GIT_MSG
  }
  stage('commitArgs Parsing') {
    sh "go get ./..."
    sh "/usr/bin/go install"
    sh "commitArgs ${env.GIT_MSG}"
  }
}
