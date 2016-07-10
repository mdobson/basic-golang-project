package main

import (
  "crypto/md5"
  "io"
)

type Player struct {
  id int
  name string
  money int
  shares map[string]int
}

func CreatePlayer(name string) Player {
  return Player{
    name: name,
    money: 100,
    shares: map[string]int{},
  }
}

func (p *Player) createToken() string {
  h := md5.New()
  io.WriteString(h, p.name)
  buf := h.Sum(nil)
  return string(buf[:])
}

type Brokers []Player
