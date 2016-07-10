package main

var currentId int
var playerId int
var market Market
var brokers Brokers

func init() {
  StateCreateCompany(Company{
    name: "Standard Oil",
    shares: 1000,
    money: 1000000,
    pricePerShare: 10,
  })

  StateCreateCompany(Company{
    name: "U.S. Steel",
    shares: 1000,
    money: 1000000,
    pricePerShare: 10,
  })

  player := CreatePlayer("mdobs")
  StateCreatePlayer(player)
}

func StateCreateCompany(c Company) Company{
  currentId += 1
  c.id = currentId
  market = append(market, c)
  return c
}

func StateCreatePlayer(p Player) Player {
  playerId += 1
  p.id = playerId
  brokers = append(brokers, p)
  return p
}

func StateFindPlayer(id int) Player {
  for _, p := range brokers {
    if p.id == id {
      return p
    }
  }

  return Player{}
}

func StateFindAllCompanies() Market {
  return market
}
