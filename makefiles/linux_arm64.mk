# Makefile extensions for linux-arm64.

# -----------------------------------------------------------------------------
# OS-ARCH specific targets
# -----------------------------------------------------------------------------

.PHONY: build-osarch-specific
build-osarch-specific: linux/arm64

# -----------------------------------------------------------------------------
# Makefile targets supported only by this platform.
# -----------------------------------------------------------------------------

.PHONY: only-linux-arm64
only-linux-arm64:
	@echo "Only linux-arm64 has this Makefile target."
