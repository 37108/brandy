package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func SessionToEnv(sess *session.Session) {
	var value = credentials.Value{}
	credentials := sess.Config.Credentials
	value, _ = credentials.Get()

	os.Setenv("AWS_ACCESS_KEY_ID", value.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", value.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", value.SessionToken)
}

func Exec() {}
