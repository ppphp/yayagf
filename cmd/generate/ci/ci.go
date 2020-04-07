package ci

import (
	"bytes"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			tmpl, err  := template.New("jenkins").Parse(jenkinsTemplate)
			if err != nil {
				log.Printf("jenkinsTemplate parse failed: %v", err.Error())
				return 1, err
			}
			b := bytes.Buffer{}
			if err :=tmpl.Execute(&b, file.GetProjectInfo()); err != nil {
				log.Printf("jenkinsTemplate parse failed: %v", err.Error())
				return 1, err
			}
			if err := ioutil.WriteFile(filepath.Join(root, "Jenkinsfile"), b.Bytes(), 0644); err != nil {
				log.Printf("write file failed: %v", err.Error())
				return 1, err
			}
			return 0, nil
		},
	}
	return c, nil
}

const jenkinsTemplate = `
pipeline {
    agent any
    tools {
        go 'Go 1.13'
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    try {
                        switch(gitlabActionType) {
                        case "PUSH":
                            echo "repo: ${gitlabSourceRepoHomepage} user: ${gitlabUserName} action: ${gitlabActionType} before: ${gitlabBefore} after: ${gitlabAfter}"
                            sh "git checkout ${gitlabAfter}"
                            break
                        case "TAG_PUSH":
                            echo "repo: ${gitlabSourceRepoHomepage} user: ${gitlabUserName} action: ${gitlabActionType} before: ${gitlabBefore} after: ${gitlabAfter}"
                            sh "git checkout ${gitlabAfter}"
                            break
                        default:
                            echo gitlabActionType
                        }
                    } catch (Exception ex){
                        gitlabActionType = "push by hand"
                    }
                }
            }
        }

        stage('Build') {
            steps {
                sh 'GOPROXY=https://goproxy.io GOSUMDB=off GOPRIVATE=gitlab.papegames.com/* go build -a'
            }
            post {
                failure {
                    dingTalk accessToken: 'https://oapi.dingtalk.com/robot/send?access_token=90009a26b4bb8aadef33a1fee8062da3837af3b032a1f3a4494afed203f8707b', jenkinsUrl: 'http://192.168.0.97:10086/',
                            message: "TDS {{.BinName}}测试环境 构建失败", notifyPeople: 'Jenkins'
                }
            }
        }

        stage('Stage') {
            when{ not {equals actual: gitlabActionType, expected: "TAG_PUSH"} }
            steps {
            	sh 'ssh root@101.37.25.45 "mkdir -p /data/tds/{{.BinName}}"'
            	sh 'ssh root@101.37.25.45 "screen -S {{.BinName}} -X quit" || true'
                sh 'scp ./{{.BinName}} root@101.37.25.45:/data/tds/{{.BinName}}/'
                sh "JENKINS_NODE_COOKIE=dontKillMe ssh root@101.37.25.45 \"cd /data/tds/{{.BinName}}; screen -dmS {{.BinName}} bash -c \\\"./{{.BinName}} 1>>/data/tds/logs/{{.BinName}}.log 2>>/data/tds/logs/{{.BinName}}.log \\\"\""
            }
            post {
                failure {
                    dingTalk accessToken: 'https://oapi.dingtalk.com/robot/send?access_token=90009a26b4bb8aadef33a1fee8062da3837af3b032a1f3a4494afed203f8707b', jenkinsUrl: 'http://192.168.0.97:10086/',
                            message: "TDS {{.BinName}}测试环境 部署失败", notifyPeople: 'Jenkins'
                }
            }
        }

        stage("Quality Gate") {
            when{ equals actual: gitlabActionType, expected: "PUSH" }
            environment {
                scannerHome = tool 'ci_server_scanner'
            }
            steps {
                script{
                    withSonarQubeEnv('ci_server') {
                        sh "${scannerHome}/bin/sonar-scanner -Dsonar.projectKey=gift_svr"
                    }
                }
                sleep(10)
                timeout(time: 10, unit: 'SECONDS') {
                    waitForQualityGate abortPipeline: false
                }
            }
            post {
                failure {
                    dingTalk accessToken: 'https://oapi.dingtalk.com/robot/send?access_token=90009a26b4bb8aadef33a1fee8062da3837af3b032a1f3a4494afed203f8707b', jenkinsUrl: 'http://192.168.0.97:10086/',
                            message: "TDS {{.BinName}} 代码太烂了", notifyPeople: 'Jenkins'
                }
            }
        }
        stage('Package') {
            when{ equals actual: gitlabActionType, expected: "TAG_PUSH" }
            steps {
                script {
                     tag = sh(script: "git tag --points-at ${gitlabAfter}", returnStdout: true).trim()
                     if (tag==""){
                     	throw "no tag"
                     }
                     now = sh(returnStdout: true, script: "date '+%Y-%m-%d-%H-%M-%S'").trim()
                     fname = "{{.BinName}}-${tag}_${now}.tar.gz"
                     sh("tar -czf ${fname} {{.BinName}} webui/dist/webui")
                     sh "scp ${fname} admin@192.168.0.11:/volume1/SDK/tds/"
                }
            }
            post {
                success {
                	 dingTalk accessToken: 'https://oapi.dingtalk.com/robot/send?access_token=90009a26b4bb8aadef33a1fee8062da3837af3b032a1f3a4494afed203f8707b', jenkinsUrl: 'http://192.168.0.97:10086/',
                     		 message: "TDS {{.BinName}}最新tag: ${tag}", notifyPeople: 'Jenkins'

                }
                failure {
                    dingTalk accessToken: 'https://oapi.dingtalk.com/robot/send?access_token=90009a26b4bb8aadef33a1fee8062da3837af3b032a1f3a4494afed203f8707b', jenkinsUrl: 'http://192.168.0.97:10086/',
                            message: "TDS {{.BinName}} ${tag} 打包失败", notifyPeople: 'Jenkins'
                }
            }
        }
        stage('Clean') {
            steps {
                cleanWs()
            }
        }
    }
}
`
