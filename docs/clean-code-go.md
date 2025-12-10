# Clean code guide for Go (Short)

1. Use descriptive, compact names and self-explanatory code.
2. Always try to write test code.
3. Minimize the comments.
   - Try to use comments only when it is for `go doc`, or when it is necessary to explain the logic.
   - Self-explanatory code is better than comments.
   - Use `t.Run` or other descriptive methods for test functions.
4. Categorize by function.
   - Go standard is to group functions by their purpose, not their location.
   - Test and types are also in the scope.
