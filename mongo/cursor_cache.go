package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver"

	"github.com/rdeavilafloqast/mongobetween/lruttl"
)

// on a 64-bit machine, 1 million cursors uses around 480mb of memory
const maxCursors = 1024 * 1024

// one day expiry
const expiry = 24 * time.Hour

type cursorCache struct {
	c *lruttl.Cache
}

func newCursorCache() *cursorCache {
	return &cursorCache{
		c: lruttl.New(maxCursors, expiry),
	}
}

func (c *cursorCache) count() int {
	return c.c.Len()
}

func (c *cursorCache) peek(cursorID int64) (server driver.Server, ok bool) {
	v, ok := c.c.Peek(cursorID)
	if !ok {
		return
	}
	return v.(driver.Server), true
}

func (c *cursorCache) add(cursorID int64, server driver.Server) {
	c.c.Add(cursorID, server)
}

func (c *cursorCache) remove(cursorID int64) {
	c.c.Remove(cursorID)
}
