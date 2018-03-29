package sessions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"

	"github.com/percona/rds_exporter/client"
	"github.com/percona/rds_exporter/config"
)

// Sessions is a pool of AWS sessions.
type Sessions struct {
	sessions map[*config.Instance]*session.Session
}

// New creates a new sessions pool for given configuration.
func New(cfg *config.Config, client *client.Client, logger *logrus.Entry) (*Sessions, error) {
	res := &Sessions{
		sessions: make(map[*config.Instance]*session.Session, len(cfg.Instances)),
	}

	for _, instance := range cfg.Instances {
		// use given credentials, or default credential chain
		var creds *credentials.Credentials
		if instance.AWSAccessKey != "" || instance.AWSSecretKey != "" {
			creds = credentials.NewCredentials(&credentials.StaticProvider{
				Value: credentials.Value{
					AccessKeyID:     instance.AWSAccessKey,
					SecretAccessKey: instance.AWSSecretKey,
				},
			})
		}

		s, err := session.NewSession(&aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   creds,
			Region:                        aws.String(instance.Region),
			HTTPClient:                    client.HTTP(),
			LogLevel:                      aws.LogLevel(aws.LogDebug),
			Logger:                        aws.LoggerFunc(logger.Debug),
		})
		if err != nil {
			return nil, err
		}
		res.sessions[instance] = s
	}

	return res, nil
}

// Get returns sessions for given instance.
func (s *Sessions) Get(instance *config.Instance) *session.Session {
	return s.sessions[instance]
}
