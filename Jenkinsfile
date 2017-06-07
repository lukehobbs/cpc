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
    sh "/usr/bin/go install"
    sh "commitArgs ${env.GIT_MSG}"
  }
}
