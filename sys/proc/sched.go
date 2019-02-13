package proc

import (
	"fmt"
	"sync"
)

type (
	// PID is the process ID
	PID int

	// Proc is a process interface
	Proc interface {
		Terminate()
	}

	// Sched is the process scheduler
	Sched struct {
		Procs map[PID]Proc

		mu     sync.RWMutex
		nextid int
	}
)

func NewSched() *Sched {
	return &Sched{
		Procs: make(map[PID]Proc),
	}
}

// Exec the given binary
func (s *Sched) Exec(path string) error {
	s.mu.Lock()
	s.nextid++
	pid := PID(s.nextid)
	s.mu.Unlock()

	proc, err := exec(path)
	if err != nil {
		s.mu.Lock()
		s.nextid--
		s.mu.Unlock()
		return err
	}

	s.Procs[pid] = proc
	return nil
}

// Kill the process
func (s *Sched) Kill(pid int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, proc := range s.Procs {
		if int(id) == pid {
			proc.Terminate()
			delete(s.Procs, id)
			return nil
		}
	}

	return fmt.Errorf("pid %d not found", pid)
}
