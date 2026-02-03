package sys

// Peforms a system healthcheck (e.g. Neovim)
// before allowing the CLI to be used to ensure
// that all required checks pass. If something
// does not pass, it should blow up.
func Healthcheck() {
}
