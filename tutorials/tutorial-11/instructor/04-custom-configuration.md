# Instructor Notes: Custom Configuration

## Teaching Techniques
- Show Zap configuration options
- Demonstrate different encoders
- Show development vs production configs
- Emphasize: "Configuration matters"

## Key Emphasis
- **Development**: Human-readable (console encoder)
- **Production**: Machine-readable (JSON encoder)
- **Levels**: Configure minimum log level
- **Output**: Configure where logs go

## Common Questions
- "What encoder should I use?" - Console for dev, JSON for prod
- "How do I set log level?" - Use zap.NewProduction/Development or custom
- "Where do logs go?" - Configure output destination

## Engagement
- "Notice the difference between encoders"
- "Development needs readability, production needs structure"
- "Configuration is important"

## Real-World Context
- Different configs for dev/prod
- JSON for production (easy to parse)
- Console for development (human-readable)

## Transition
- "Let's see best practices..."
