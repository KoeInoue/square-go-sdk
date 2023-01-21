package square

type Sandbox string
type Production string

// Environment has part of the domain that varies with the environment.
type Env struct {
	Sandbox    Sandbox
	Production Production
}

// Config struct contains access token and environment
type Config struct {
	AccessToken string
	Environment Env
}

const (
	ENV_SANDBOX    = Sandbox("https://connect.squareupsandbox.com")
	ENV_PRODUCTION = Production("https://connect.squareup.com")
)

// Environment has part of the domain that varies with the environment.
var Environment = Env{
	Sandbox:    ENV_SANDBOX,
	Production: ENV_PRODUCTION,
}
