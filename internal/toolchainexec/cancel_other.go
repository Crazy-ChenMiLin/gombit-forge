//go:build !unix

package toolchainexec

import "os/exec"

// CommandContext cancels the direct process. WaitDelay bounds pipe draining.
func configureCancellation(*exec.Cmd) {}
