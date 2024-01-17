package collider

func (c *Collider[Seed]) generatorLoop() {
	c.generator(c.finished, c.seedChan)
	close(c.seedChan)
}
