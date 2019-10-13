package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
)

func Client(profile *Profile) (*aws.Client, error) {
	cfg, err := external.LoadDefaultAWSConfig(
		external.WithMFATokenFunc(stscreds.StdinTokenProvider),
		external.WithSharedConfigProfile(profile.Name),
	)
	if err != nil {
		return nil, err
	}

	svc := sts.New(cfg)
	input := &sts.GetCallerIdentityInput{}
	req := svc.GetCallerIdentityRequest(input)
	_, err = req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	return svc.Client, nil
}
