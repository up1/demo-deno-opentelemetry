FROM denoland/deno:latest

# Create working directory
WORKDIR /app

# Copy source
COPY . .

# Compile the main app
RUN deno cache server.ts

# Run the app
CMD ["deno", "run", "--allow-net", "--unstable-otel","server.ts"]