package source

import "context"

type Syncer interface {
	Sync(context.Context) error
}
