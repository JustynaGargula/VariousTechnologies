FROM ubuntu:24.04

RUN apt-get update
RUN apt-get install -y python3.10
RUN apt-get install -y openjdk-8-jre
RUN apt-get install -y default-jdk

RUN apt-get install -y zip
RUN apt-get install -y unzip
RUN apt-get install -y curl
SHELL ["/bin/bash", "-c"]
RUN mkdir gradleProject
WORKDIR /home/user/gradleProject
ADD build.gradle .
ADD HelloWorld.java .
RUN curl -s "https://get.sdkman.io" | bash && \
    source "/root/.sdkman/bin/sdkman-init.sh" && \
    sdk install kotlin && sdk install gradle && \
    ln -s $HOME/.sdkman/candidates/gradle/current/bin/gradle /usr/bin/gradle && \
    gradle build

CMD [ "gradle", "run" ]