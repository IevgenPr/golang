```
git init
git remote add origin git@github.com:ipr0/golang
git push --set-upstream origin master
go mod init github.com/ipr0/golang21

```

### Github actions

- install incoming webhooks app on Slack
- put webhoot url into github settings/secrets section as SLACK_WEBHOOK_URL
- put actions config into `.github/workflows/xxx.yml` file
