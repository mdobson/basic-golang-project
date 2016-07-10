package main

type Company struct {
  id int
  name string
  shares int
  money int
  pricePerShare int
}

func (c *Company) price(desiredShares int) int {
  return desiredShares * c.pricePerShare
}

type Market []Company
