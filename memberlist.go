package swim

import "time"

type Config struct {
	ProbeTimeout time.Duration
	Interval     time.Duration
}

type MemberList struct {
	conf *Config
}

func Create(conf *Config) (*MemberList, error) {

}

func (m *MemberList) newMemberList(conf *Config) (*MemberList, error) {

}

func (m *MemberList) Join() (int, error) {

}

func (m *MemberList) Stop() {

}
