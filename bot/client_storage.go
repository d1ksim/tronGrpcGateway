package bot

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"sync"
)

type Client struct {
	rwMux    sync.RWMutex
	userData map[int64]map[string]any

	// This struct could also contain:
	// - chat data, conversation data, etc
	// - pointers to database/cache connections
	// - localised strings
	// - helper methods for retrieving/caching chat settings
	// Really, any "global" data that should be accessible by your bot
}

func (c *Client) GetUserData(ctx *ext.Context, key string) (any, bool) {
	c.rwMux.RLock()
	defer c.rwMux.RUnlock()

	if c.userData == nil {
		return nil, false
	}

	userData, ok := c.userData[ctx.EffectiveUser.Id]
	if !ok {
		return nil, false
	}

	v, ok := userData[key]
	return v, ok
}

func (c *Client) SetUserData(ctx *ext.Context, key string, val any) {
	c.rwMux.Lock()
	defer c.rwMux.Unlock()

	if c.userData == nil {
		c.userData = map[int64]map[string]any{}
	}

	_, ok := c.userData[ctx.EffectiveUser.Id]
	if !ok {
		c.userData[ctx.EffectiveUser.Id] = map[string]any{}
	}
	c.userData[ctx.EffectiveUser.Id][key] = val
}
