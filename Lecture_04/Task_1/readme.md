Write a function called daysInMonth:
func daysInMonth(month int, year int) (int, bool) {
// ...
}
When passed a month value (in the range 1-12) and a year, calculate the number of days in the
given month. Check if it is a leap year, by using modulo division. When it is, return 29 when
asked for February, rather than 8.
Use a switch statement to decide, whether to return 30 or 31 days (or 28/29 for February),
depending on the month.