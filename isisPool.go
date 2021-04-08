package isis

import "sync"

type Pool struct {
	waitGroup *sync.WaitGroup
	errors chan error
}

func (pool *Pool) run(cmd string, args map[string]string) {
	defer pool.waitGroup.Done()
	pool.errors <- Isis(cmd, args)
}

func (pool *Pool) Run(cmd string, args map[string]string) {
	pool.waitGroup.Add(1)
	go pool.run(cmd, args)
}

func (pool *Pool) Wait() []error {
	pool.waitGroup.Wait()
	close(pool.errors)

	errs := make([]error, 0)
	for err := range pool.errors {
		errs = append(errs, err)
	}

	return errs
}

func NewPool() *Pool {
	return &Pool{
		waitGroup: &sync.WaitGroup{},
		errors:    make(chan error, 1024),
	}
}
