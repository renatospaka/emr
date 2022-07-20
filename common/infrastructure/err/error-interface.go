package err

type ValidationEntityInterface interface {
	IsValid() bool
	Err() []string
}