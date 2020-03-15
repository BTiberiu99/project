package puzzle

type Stack struct {
	configs []*Config
}

func (l *Stack) Pop() *Config {
	var c *Config
	c, l.configs = l.configs[len(l.configs)-1], l.configs[:len(l.configs)-1]
	return c
}

func (l *Stack) Append(c *Config) *Stack {
	l.configs = append(l.configs, c)

	return l
}

func (l *Stack) Len() int {
	return len(l.configs)
}

func (l *Stack) Configs() []*Config {
	return l.configs
}

func NewStack() *Stack {
	return &Stack{
		configs: make([]*Config, 0),
	}
}
