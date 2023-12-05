# Makefile extensions for linux-x86_64.

# -----------------------------------------------------------------------------
# OS-ARCH specific targets
# -----------------------------------------------------------------------------

.PHONY: build-osarch-specific
build-osarch-specific: linux/amd64

# -----------------------------------------------------------------------------
# Makefile targets supported only by this platform.
# -----------------------------------------------------------------------------

.PHONY: only-linux-x86_64
only-linux-x86_64:
	@echo "Only linux-x86_64 has this Makefile target."
