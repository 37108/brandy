package config

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/defaults"
	"gopkg.in/ini.v1"
)

type Profile struct {
	Name          string
	MFASerial     string
	RoleArn       string
	SourceProfile string
}

func Parser() []*Profile {
	file := defaults.SharedConfigFilename()
	cfg, _ := ini.Load(file)

	sections := cfg.Sections()
	profiles := make([]*Profile, 0, len(sections))
	for _, section := range sections {
		name := section.Name()
		name = strings.TrimSpace(strings.Replace(name, "profile ", "", 1))
		mfa := section.Key("mfa_serial").String()
		roleArn := section.Key("role_arn").String()
		sourceProfile := section.Key("source_profile").String()

		profiles = append(profiles, &Profile{
			Name:          name,
			RoleArn:       roleArn,
			MFASerial:     mfa,
			SourceProfile: sourceProfile,
		})
	}
	return profiles
}
