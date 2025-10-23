package event_handlers

import "context"

type EventHandle func(ctx context.Context) (interface{}, error)
