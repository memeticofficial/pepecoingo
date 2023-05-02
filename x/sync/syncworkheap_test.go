// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/ids"
)

// Tests heap.Interface methods Push, Pop, Swap, Len, Less.
func Test_SyncWorkHeap_Heap_Methods(t *testing.T) {
	require := require.New(t)

	h := newSyncWorkHeap()
	require.Zero(h.Len())

	item1 := &heapItem{
		workItem: &syncWorkItem{
			start:       nil,
			end:         nil,
			priority:    highPriority,
			LocalRootID: ids.GenerateTestID(),
		},
	}
	h.Push(item1)
	require.Equal(1, h.Len())
	require.Len(h.priorityHeap, 1)
	require.Equal(item1, h.priorityHeap[0])
	require.Zero(h.priorityHeap[0].heapIndex)
	require.Equal(1, h.sortedItems.Len())
	gotItem, ok := h.sortedItems.Get(item1)
	require.True(ok)
	require.Equal(item1, gotItem)

	h.Pop()
	require.Zero(h.Len())
	require.Empty(h.priorityHeap)
	require.Zero(h.sortedItems.Len())

	item2 := &heapItem{
		workItem: &syncWorkItem{
			start:       []byte{0},
			end:         []byte{1},
			priority:    highPriority,
			LocalRootID: ids.GenerateTestID(),
		},
	}
	h.Push(item1)
	h.Push(item2)
	require.Equal(2, h.Len())
	require.Len(h.priorityHeap, 2)
	require.Equal(item1, h.priorityHeap[0])
	require.Equal(item2, h.priorityHeap[1])
	require.Zero(item1.heapIndex)
	require.Equal(1, item2.heapIndex)
	require.Equal(2, h.sortedItems.Len())
	gotItem, ok = h.sortedItems.Get(item1)
	require.True(ok)
	require.Equal(item1, gotItem)
	gotItem, ok = h.sortedItems.Get(item2)
	require.True(ok)
	require.Equal(item2, gotItem)

	require.False(h.Less(0, 1))

	h.Swap(0, 1)
	require.Equal(item2, h.priorityHeap[0])
	require.Equal(item1, h.priorityHeap[1])
	require.Equal(1, item1.heapIndex)
	require.Zero(item2.heapIndex)

	require.False(h.Less(0, 1))

	item1.workItem.priority = lowPriority
	require.True(h.Less(0, 1))

	gotItem = h.Pop().(*heapItem)
	require.Equal(item1, gotItem)

	gotItem = h.Pop().(*heapItem)
	require.Equal(item2, gotItem)

	require.Zero(h.Len())
	require.Empty(h.priorityHeap)
	require.Zero(h.sortedItems.Len())
}

// Tests Insert and GetWork
func Test_SyncWorkHeap_Insert_GetWork(t *testing.T) {
	require := require.New(t)
	h := newSyncWorkHeap()

	item1 := &syncWorkItem{
		start:       []byte{0},
		end:         []byte{1},
		priority:    lowPriority,
		LocalRootID: ids.GenerateTestID(),
	}
	item2 := &syncWorkItem{
		start:       []byte{2},
		end:         []byte{3},
		priority:    medPriority,
		LocalRootID: ids.GenerateTestID(),
	}
	item3 := &syncWorkItem{
		start:       []byte{4},
		end:         []byte{5},
		priority:    highPriority,
		LocalRootID: ids.GenerateTestID(),
	}
	h.Insert(item3)
	h.Insert(item2)
	h.Insert(item1)
	require.Equal(3, h.Len())

	// Ensure [sortedItems] is in right order.
	got := []*syncWorkItem{}
	h.sortedItems.Ascend(
		func(i *heapItem) bool {
			got = append(got, i.workItem)
			return true
		},
	)
	require.Equal([]*syncWorkItem{item1, item2, item3}, got)

	// Ensure priorities are in right order.
	gotItem := h.GetWork()
	require.Equal(item3, gotItem)
	gotItem = h.GetWork()
	require.Equal(item2, gotItem)
	gotItem = h.GetWork()
	require.Equal(item1, gotItem)
	gotItem = h.GetWork()
	require.Nil(gotItem)

	require.Zero(h.Len())
}

func Test_SyncWorkHeap_remove(t *testing.T) {
	require := require.New(t)

	h := newSyncWorkHeap()

	item1 := &syncWorkItem{
		start:       []byte{0},
		end:         []byte{1},
		priority:    lowPriority,
		LocalRootID: ids.GenerateTestID(),
	}

	h.Insert(item1)

	heapItem1 := h.priorityHeap[0]
	h.remove(heapItem1)

	require.Zero(h.Len())
	require.Empty(h.priorityHeap)
	require.Zero(h.sortedItems.Len())

	item2 := &syncWorkItem{
		start:       []byte{2},
		end:         []byte{3},
		priority:    medPriority,
		LocalRootID: ids.GenerateTestID(),
	}

	h.Insert(item1)
	h.Insert(item2)

	heapItem2 := h.priorityHeap[0]
	require.Equal(item2, heapItem2.workItem)
	h.remove(heapItem2)
	require.Equal(1, h.Len())
	require.Len(h.priorityHeap, 1)
	require.Equal(1, h.sortedItems.Len())
	require.Zero(h.priorityHeap[0].heapIndex)
	require.Equal(item1, h.priorityHeap[0].workItem)

	heapItem1 = h.priorityHeap[0]
	require.Equal(item1, heapItem1.workItem)
	h.remove(heapItem1)
	require.Zero(h.Len())
	require.Empty(h.priorityHeap)
	require.Zero(h.sortedItems.Len())
}

func Test_SyncWorkHeap_Merge_Insert(t *testing.T) {
	require := require.New(t)

	// merge with range before
	syncHeap := newSyncWorkHeap()

	syncHeap.MergeInsert(&syncWorkItem{start: nil, end: []byte{63}})
	require.Equal(1, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{127}, end: []byte{192}})
	require.Equal(2, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{193}, end: nil})
	require.Equal(3, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{63}, end: []byte{126}, priority: lowPriority})
	require.Equal(3, syncHeap.Len())

	// merge with range after
	syncHeap = newSyncWorkHeap()

	syncHeap.MergeInsert(&syncWorkItem{start: nil, end: []byte{63}})
	require.Equal(1, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{127}, end: []byte{192}})
	require.Equal(2, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{193}, end: nil})
	require.Equal(3, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{64}, end: []byte{127}, priority: lowPriority})
	require.Equal(3, syncHeap.Len())

	// merge both sides at the same time
	syncHeap = newSyncWorkHeap()

	syncHeap.MergeInsert(&syncWorkItem{start: nil, end: []byte{63}})
	require.Equal(1, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{127}, end: nil})
	require.Equal(2, syncHeap.Len())

	syncHeap.MergeInsert(&syncWorkItem{start: []byte{63}, end: []byte{127}, priority: lowPriority})
	require.Equal(1, syncHeap.Len())
}
