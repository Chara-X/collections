package collections

import (
	"container/list"
	"context"
)

type MsgQueue[T any] struct{ subscribers *list.List }

func NewMsgQueue[T any]() *MsgQueue[T] { return &MsgQueue[T]{subscribers: list.New()} }
func (q *MsgQueue[T]) Subscribe(ctx context.Context) <-chan T {
	var subscriber = q.subscribers.PushBack(make(chan T))
	go func() {
		<-ctx.Done()
		q.subscribers.Remove(subscriber)
		close(subscriber.Value.(chan T))
	}()
	return subscriber.Value.(chan T)
}
func (q *MsgQueue[T]) Publish(msg T) {
	for e := q.subscribers.Front(); e != nil; e = e.Next() {
		e.Value.(chan T) <- msg
	}
}
