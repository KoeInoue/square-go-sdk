package square

type Sandbox string
type Production string

// Environment has part of the domain that varies with the environment.
type Env interface {
    Sandbox | Production
}

// Config struct contains access token and environment
type Config[T Env] struct {
    AccessToken string
    Environment T
}

const (
	ENV_SANDBOX    = Sandbox("https://connect.squareupsandbox.com")
	ENV_PRODUCTION = Production("https://connect.squareup.com")
)

// Environment has part of the domain that varies with the environment.
var Environments = struct{
    Sandbox Sandbox
    Production Production
}{
    Sandbox: ENV_SANDBOX,
    Production: ENV_PRODUCTION,
}
