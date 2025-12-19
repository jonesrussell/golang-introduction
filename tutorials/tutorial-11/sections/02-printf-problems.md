# The Problem with Printf

**Duration:** 4-5 minutes

## The Anti-Pattern

```go
// BAD: Unstructured logging
log.Printf("User %d logged in from %s at %s", userID, ip, time.Now())
log.Printf("Error: %v", err)
log.Printf("Request completed in %dms", elapsed)
```

## Problems:
1. Hard to parse programmatically
2. Inconsistent formats
3. Can't filter by fields
4. Poor for log aggregation

## Structured Logging Benefits:
- Machine-parseable (JSON)
- Queryable fields
- Consistent format
- Easy aggregation
- Better performance
