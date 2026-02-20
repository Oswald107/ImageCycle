# ---- Base image with JDK 21 ----
FROM eclipse-temurin:21-jdk AS base

# Set work directory
WORKDIR /workspace

# Install useful dev tools
RUN apt-get update && apt-get install -y \
    curl unzip git vim bash \
    && rm -rf /var/lib/apt/lists/*

# Install Gradle (optional if you use the Gradle wrapper)
ENV GRADLE_HOME=/opt/gradle
ENV PATH=$GRADLE_HOME/bin:$PATH

# Copy Gradle wrapper files early to leverage caching
COPY gradlew gradle/ ./

# Ensure wrapper is executable
RUN chmod +x gradlew

# Copy the rest of the source
COPY . .

# Pre-download dependencies to speed up dev startup
RUN ./gradlew dependencies --no-daemon || true

# Default command (can be overridden by VS Code)
CMD ["./gradlew", "bootRun"]
