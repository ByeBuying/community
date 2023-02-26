package model

import (
	"fmt"
	"go-common/klay/elog"
	"reflect"
	"sync"

	"community/conf"
)

type IRepository interface {
	Start() error
}

type RepositoryConstructor func(conf *conf.Config, root *Repositories) (IRepository, error)

type Repositories struct {
	lock  sync.RWMutex
	conf  *conf.Config
	elems map[reflect.Type]reflect.Value
}

func NewRepositories(cfg *conf.Config) (*Repositories, error) {
	r := &Repositories{
		conf:  cfg,
		elems: make(map[reflect.Type]reflect.Value),
	}

	for _, c := range []struct {
		constructor RepositoryConstructor
		config      *conf.Config
	}{
		{NewRedisDB, cfg},
		{NewCommunityDB, cfg},
		//{NewProdu, cfg},
	} {
		if err := r.Register(c.constructor, c.config); err != nil {
			return nil, err
		}
	}

	if err := func() error {
		r.lock.Lock()
		defer r.lock.Unlock()

		for t, e := range r.elems {
			if err := e.Interface().(IRepository).Start(); err != nil {
				elog.Error("NewRepositories", "repository", t, "error", err)
				return err
			}
		}
		return nil
	}(); err != nil {
		return nil, err
	}
	return r, nil
}

func (p *Repositories) Get(rs ...interface{}) error {
	p.lock.RLock()
	defer p.lock.RUnlock()

	notFounds := []reflect.Type{}

	for _, v := range rs {
		elem := reflect.ValueOf(v).Elem()
		if e, ok := p.elems[elem.Type()]; ok == true {
			elem.Set(e)
		} else {
			notFounds = append(notFounds, elem.Type())
		}
	}

	if len(notFounds) > 0 {
		err := fmt.Errorf("unknown repository ")
		for _, e := range notFounds {
			err = fmt.Errorf("%v, %v ", err.Error(), e)
		}
		return err
	}

	return nil
}

func (p *Repositories) Register(constructor RepositoryConstructor, config *conf.Config) error {
	if r, err := constructor(config, p); err != nil {
		return err
	} else if r != nil {
		p.lock.Lock()
		defer p.lock.Unlock()

		if _, ok := p.elems[reflect.TypeOf(r)]; ok == true {
			return fmt.Errorf("duplicated instance of %v", reflect.TypeOf(r))
		} else {
			p.elems[reflect.TypeOf(r)] = reflect.ValueOf(r)
		}
	}
	return nil
}
