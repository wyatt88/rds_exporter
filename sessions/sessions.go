package sessions

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/percona/rds_exporter/config"
)

// Sessions is a pool of aws *session.Sessions.
type Sessions struct {
	sessions map[config.Instance]*session.Session
	sync.RWMutex
}

func (s *Sessions) Get(instance config.Instance) *session.Session {
	s.RLock()
	defer s.RUnlock()
	return s.sessions[instance]
}

func (s *Sessions) Load(instances []config.Instance) error {
	s.Lock()
	defer s.Unlock()
	return s.load(instances)
}

func (s *Sessions) load(instances []config.Instance) error {
	// Destroy known sessions
	s.sessions = map[config.Instance]*session.Session{}

	// Load new sessions
	for _, instance := range instances {
		s.loadOne(instance)
	}
	return nil
}

func (s *Sessions) loadOne(instance config.Instance) {
	awsConfig := &aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region: aws.String(instance.Region),
		// TODO HTTPClient
		// TODO Logger
	}

	// If aws_access_key or aws_secret_key is present in config
	// then use those rather than relying on automatic configuration detection in aws library.
	if instance.AWSAccessKey != "" || instance.AWSSecretKey != "" {
		awsConfig.Credentials = credentials.NewCredentials(&credentials.StaticProvider{
			Value: credentials.Value{
				AccessKeyID:     instance.AWSAccessKey,
				SecretAccessKey: instance.AWSSecretKey,
			},
		})
	}

	// Cache session
	s.sessions[instance] = session.Must(session.NewSession(awsConfig))
}
