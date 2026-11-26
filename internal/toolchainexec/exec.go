// Package toolchainexec runs build and generator commands with bounded cancellation.
package toolchainexec

import (
	"context"
	"os/exec"
	"time"
)

func CombinedOutput(ctx context.Context, dir, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	// A descendant can keep a copy of an output pipe after the command exits.
	cmd.WaitDelay = time.Second
	configureCancellation(cmd)
	return cmd.CombinedOutput()
}
