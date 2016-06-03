# CF Sample App Java

A sample Java [Spark](http://sparkjava.com) application to deploy to Cloud Foundry which works out of the box.

## Run locally

1. Install the [Java JDK](http://www.oracle.com/technetwork/java/javase/downloads/index.html)
1. Install [Gradle](https://gradle.org/gradle-download/)
1. Run `gradle build`
1. Run `java -jar build/libs/cf-sample-java-1.0.jar`
1. Visit [http://localhost:4567](http://localhost:4567)

## Run in the cloud

1. Install the [Java JDK](http://www.oracle.com/technetwork/java/javase/downloads/index.html)
1. Install [Gradle](https://gradle.org/gradle-download/)
1. Install the [cf CLI](https://github.com/cloudfoundry/cli#downloads)
1. Run `gradle build`
1. Run `cf push my-java-app -p build/libs/cf-sample-app-java-1.0.jar --random-route`
1. Visit the given URL
