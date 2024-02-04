FROM eclipse-temurin:17-jdk as builder

WORKDIR /usr/src/app

COPY . /usr/src/app

ENTRYPOINT ["./gradlew", "bootRun"]
#FROM eclipse-temurin:17-jdk

#WORKDIR /usr/src/app


#cOPY --from=builder /usr/src/app .

#ENTRYPOINT ["java", "-Djava.security.egd=file:/dev/./urandom", "-jar", "/eureka.jar"]
