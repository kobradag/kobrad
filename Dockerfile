

# Use the minimal Ubuntu 22.04 base image
FROM ubuntu:22.04
LABEL authors="hainakus"
# Set working directory
WORKDIR /app/

# Copy the built binary into the container's /app/ directory
COPY . /app/

# Move the kobrad binary to /usr/local/bin for global access
RUN mv /app/kobrad /usr/local/bin/kobrad

# Make sure the kobrad binary is executable
RUN chmod +x /usr/local/bin/kobrad

# Expose necessary ports for P2P and RPC communication
EXPOSE 44448 44447

# Set the entrypoint to run the 'kobrad' binary
ENTRYPOINT ["/usr/local/bin/Kobrad"]

# Default command with necessary flags
CMD ["--utxoindex"]
