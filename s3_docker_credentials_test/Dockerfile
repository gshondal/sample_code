#Dockerfile
# The following sample code is for illustrative purposes only, and must not be used in production environments.

FROM openjdk:7

#Environmental variables
ENV BASEDIR=/opt
ENV HOMEDIR=$BASEDIR/java/dev
ENV ARTIFACTID=sdk-sandbox
ENV VERSION=1.0

#Download maven
WORKDIR $BASEDIR
RUN wget https://mirrors.ocf.berkeley.edu/apache/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.zip
RUN unzip apache-maven-3.6.3-bin.zip
ENV PATH=$PATH:$BASEDIR/apache-maven-3.6.3/bin

#Creating Home directory
RUN mkdir -p $HOMEDIR
#Setup java structure
WORKDIR $HOMEDIR
RUN mvn archetype:generate -B -DarchetypeGroupId=pl.org.miki -DarchetypeArtifactId=java8-quickstart-archetype -DarchetypeVersion=1.0.0 -DgroupId=com.example -DartifactId=$ARTIFACTID -Dversion=$VERSION -Dpackage=com.example
WORKDIR $HOMEDIR/$ARTIFACTID

#Deploy deps
COPY pom.xml .
RUN mvn install
#Deploying code to image
COPY CopyObject.java ./src/main/java
RUN mvn install

#Test
ENTRYPOINT java -cp target/$ARTIFACTID-$VERSION-jar-with-dependencies.jar CopyObject KEY BucketA BucketB


