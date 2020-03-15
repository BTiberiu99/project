package puzzle

type Queue struct {
	configs []*Config
}

func (q *Queue) Pop() *Config {
	var c *Config
	c, q.configs = q.configs[0], q.configs[1:]
	return c
}

func (q *Queue) Append(c *Config) *Queue {
	q.configs = append(q.configs, c)

	return q
}

func (q *Queue) Len() int {
	return len(q.configs)
}

func (q *Queue) Configs() []*Config {
	return q.configs
}

func NewQueue() *Queue {
	return &Queue{
		configs: make([]*Config, 0),
	}
}
