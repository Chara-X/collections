package collections

import (
	"container/list"
	"context"
)

type MsgQueue[T any] struct{ subscribers *list.List }

func NewMsgQueue[T any]() *MsgQueue[T] { return &MsgQueue[T]{subscribers: list.New()} }
func (c *MsgQueue[T]) Subscribe(ctx context.Context) <-chan T {
	var subscriber = c.subscribers.PushBack(make(chan T))
	go func() {
		<-ctx.Done()
		c.subscribers.Remove(subscriber)
		close(subscriber.Value.(chan T))
	}()
	return subscriber.Value.(chan T)
}
func (c *MsgQueue[T]) Publish(msg T) {
	for e := c.subscribers.Front(); e != nil; e = e.Next() {
		e.Value.(chan T) <- msg
	}
}
