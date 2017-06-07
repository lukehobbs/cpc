#! /usr/bin/env groovy

node {

  def goHome = tool 'go 1.8.3'

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
    sh "${goHome}/bin/go get ./..."
    sh "${goHome}/bin/go install"
    sh "commitArgs ${env.GIT_MSG}"
  }
}
