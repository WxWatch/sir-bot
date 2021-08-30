package leveling

// Adapted from https://github.com/jtagt/mee6/blob/0543c2c641bd99e1ef7aaae622ae066af901b8ee/chat-bot/plugins/levels.py
// \frac{1}{5}*(x^{2})+50*x+100
func ExpForLevel(level uint) uint {
	return (1/5)*(level^2) + 50*level + 100
}
