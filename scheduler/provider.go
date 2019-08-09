package scheduler

import (
	"encoding/json"
	"time"
)

const SignalFormat  =  "2006-01-02 15:04:05"
const SignalFormatWithZone  =  "2006-01-02 15:04:05 MST"

type IProvider interface {
	Init() error
	GetName() string
	CheckInterval(time.Time) bool
	Run(time.Time)
	String() string
}

type Provider struct {
	Name         string    `toml:"name" json:"name"`
	TimeRule     string    `toml:"interval" json:"interval"`
	Interval     *Interval `toml:"-" json:"-"`
}

func (p *Provider) GetName() string {
	return p.Name
}

func (p *Provider) CheckInterval(t time.Time) bool {
	if p.Interval == nil {
		return false
	}

	return p.Interval.Check(t)
}

func (p *Provider) GetSignal(t time.Time) string {
	return t.Format(SignalFormat)
}

func (p *Provider) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}