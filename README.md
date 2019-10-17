# slack-invitation handler

slack-invitation handler implemented by Go and it used on [now](https://zeit.co/)

## Deploy

* `now`
* `now --prod`

## Configuration

* `now secrets ls`
* `now secrets add slack-subdomain <domain>`.
* `now secrets add slack-api-token <api-token>`
* `now secrets add google-captcha-secret <captcha-secret>`

## Learn more about users.admin.invite Slack API

https://github.com/ErikKalkoken/slackApiDoc/blob/master/users.admin.invite.md