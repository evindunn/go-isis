package isis

import (
	"runtime"
	"sync"
)

type Pool struct {
	resultGroup  *sync.WaitGroup
	errors       chan error
	poolSize	 chan struct{}
}

func (pool *Pool) run(cmd string, args map[string]string) {
	defer pool.resultGroup.Done()

	pool.poolSize <- struct{}{}
	pool.errors <- Isis(cmd, args)
	<- pool.poolSize
}

func (pool *Pool) Run(cmd string, args map[string]string) {
	pool.resultGroup.Add(1)
	go pool.run(cmd, args)
}

func (pool *Pool) Wait() []error {
	pool.resultGroup.Wait()
	close(pool.errors)

	errs := make([]error, 0)
	for err := range pool.errors {
		errs = append(errs, err)
	}

	return errs
}

func NewPool() *Pool {
	return &Pool{
		poolSize: 	  make(chan struct{}, runtime.NumCPU()),
		resultGroup:  &sync.WaitGroup{},
		errors:       make(chan error, 1024),
	}
}
