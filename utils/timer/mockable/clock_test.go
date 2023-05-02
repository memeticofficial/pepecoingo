// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package mockable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClockSet(t *testing.T) {
	require := require.New(t)

	clock := Clock{}
	time := time.Unix(1000000, 0)
	clock.Set(time)
	require.True(clock.faked)
	require.Equal(time, clock.Time())
}

func TestClockSync(t *testing.T) {
	require := require.New(t)

	clock := Clock{true, time.Unix(0, 0)}
	clock.Sync()
	require.False(clock.faked)
	require.NotEqual(time.Unix(0, 0), clock.Time())
}

func TestClockUnixTime(t *testing.T) {
	require := require.New(t)

	clock := Clock{true, time.Unix(123, 123)}
	require.Zero(clock.UnixTime().Nanosecond())
	require.Equal(123, clock.Time().Nanosecond())
}

func TestClockUnix(t *testing.T) {
	clock := Clock{true, time.Unix(-14159040, 0)}
	actual := clock.Unix()
	if actual != 0 {
		// We are Unix of 1970s, Moon landings are irrelevant
		t.Errorf("Expected time prior to Unix epoch to be clamped to 0, got %d", actual)
	}
}
