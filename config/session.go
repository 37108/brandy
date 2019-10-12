package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ObtainSession(profile *Profile) (*session.Session, error) {
	if profile.MFASerial == "" && profile.RoleArn == "" {
		sess := session.Must(session.NewSession(&aws.Config{
			Credentials: credentials.NewSharedCredentials("", profile.Name),
		}))
		return sess, nil
	}
	return nil, nil
}
