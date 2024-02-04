FROM eclipse-temurin:17-jdk as builder

WORKDIR /usr/src/app
COPY . /usr/src/app
ENTRYPOINT ["./gradlew", "bootRun"]
